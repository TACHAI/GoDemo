package main

import (
	"Demo01/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main()  {
	rpc.Register(rpcdemo.DemoService{})
	listener,err :=net.Listen("tcp",":1234")

	if err !=nil{
		panic(err)
	}

	for{
		conn,err := listener.Accept()
		if err !=nil{
			log.Panicf("accept error: %v",err)
			continue
		}


		go jsonrpc.ServeConn(conn)
	}
}
