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

type Poster interface {
	Post(url string,form map[string]string)string
}

const url  = "http://www.imooc.com"
func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func post(poster Poster)  {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name":"ccmouse",
			"course":"golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}
func session(s RetrieverPoster)  string{
	//s.Get()
	s.Post(url,map[string]string{
		"contents":"ssd",
	})
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
	reteiever := mock.Retriever{"this is imooc.com"}
	fmt.Printf("%T %v\n",r,r)
	inspect(r)
	r = &real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	fmt.Printf("%T %v\n",r,r)
	inspect(r)


	fmt.Printf(session(reteiever))
	//fmt.Println(download(r))
}