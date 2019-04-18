package main

import (
	"Demo01/rpc"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main()  {
	conn,err :=net.Dial("tcp",":1234")
	if err != nil{
		panic(err)
	}

	client := jsonrpc.NewClient(conn)
	var result float64
	client.Call("DemoService.Div",rpcdemo.Args{10,3},&result)

	fmt.Println(result,err)
}
