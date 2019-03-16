package consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/zeeshine/grpc/discovery"
	"net"
	"time"
	"fmt"
)

type ConsulRegisterOptions struct {
	ConsulAddress                  string
	ServiceName                    string
	ServicePort                    int
	Tags                           []string
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
}

type consulRegister struct {
	opts *ConsulRegisterOptions
}

func NewConsulRegister(opts *ConsulRegisterOptions) discovery.Register {
	return &consulRegister{
		opts: opts,
	}
}

func (r *consulRegister) Register() error {
	opts := r.opts
	config := api.DefaultConfig()
	config.Address = opts.ConsulAddress
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	agent := client.Agent()
	ip := localIp()
	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", opts.ServiceName, ip, opts.ServicePort), // 服务ID
		Name:    opts.ServiceName,                                                // 服务名称
		Tags:    opts.Tags,
		Port:    opts.ServicePort,
		Address: ip,
		Check: &api.AgentServiceCheck{
			Interval:                       opts.Interval.String(),
			GRPC:                           fmt.Sprintf("%v:%v/%v", ip, opts.ServicePort, opts.ServiceName), // grpc 健康检测的地址
			DeregisterCriticalServiceAfter: opts.DeregisterCriticalServiceAfter.String(),                    // 注销时间，相当于过期时间
		},
	}

	if err := agent.ServiceRegister(reg); err != nil {
		return err
	}

	return nil
}

func localIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}
