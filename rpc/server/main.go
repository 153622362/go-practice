package main

import (
	"log"
	"net"
	"net/rpc"
		"../../rpc"
	"net/rpc/jsonrpc"
)

//jsonrpc服务器
func main() {
	//{"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
		}
		go jsonrpc.ServeConn(conn)
	}
}
