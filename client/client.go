package client

import (
	"google.golang.org/grpc"
	"sync"
)

type GrpcClient struct {
	sync.RWMutex
	clientConns map[string]*grpc.ClientConn
}

func (gc *GrpcClient) ForkConn(serviceName string) *grpc.ClientConn,error {
	gc.RLock()
	if conn,ok:=gc.clientConns[serviceName];ok{
		gc.RUnlock()
		return conn,nil
	}
	gc.RUnlock()

    gc.Lock()
    defer gc.Unlock()

    if conn,err:=grpc.Dial(serviceName);err!=nil{
    	return nil err
    }

    gc.clientConns[serviceName]=conn
    return conn,nil
}

func (gc *GrpcClient)CloseConn(serviceName string) error {
	gc.Lock()
	defer gc.Unlock()
	if conn,ok:=gc.clientConns[serviceName];ok{
		delete(gc.clientConns,serviceName)
		return conn.Close()
	}
	return nil
}
