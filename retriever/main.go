package main

import (
	"Demo01/retriever/mock"
	"Demo01/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever)  {
	switch v :=r.(type) {
	case mock.Retriever:
		fmt.Printf("Contents",v.Contents)
	case *real.Retriever:
		fmt.Printf("UserAgent",v.UserAgent)
	}
}

func main()  {
	var r Retriever
	r = mock.Retriever{"this is imooc.com"}
	fmt.Printf("%T %v\n",r,r)
	inspect(r)
	r = &real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	fmt.Printf("%T %v\n",r,r)
	inspect(r)

	//fmt.Println(download(r))
}