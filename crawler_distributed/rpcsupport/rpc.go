package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//服务提供-服务端
func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}


//服务提供-客户端
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return  nil,  err
	}

	return jsonrpc.NewClient(conn), nil
}