package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int  {
	out := make(chan int)
	go func() {
		i :=0
		for {
			time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func work(id int,c chan int){
	go func() {
		//for{
		//	n ,ok:=<-c
		//	if !ok{
		//		break
		//	}
		//	fmt.Printf("worker %d received %c\n",id,n)
		//
		//}

		for n:=range c{
			fmt.Printf("worker %d received %c\n",id,n)

		}
	}()
}

func createWork(id int) chan<- int {

	c := make(chan int)
	go func() {
		for{
			n :=<-c
			fmt.Printf("worker %d received %c\n",id,n)

		}
	}()

	return c
}


func main()  {
	// c1 and c2 = nil
	var c1,c2 = generator(),generator()
	var worker = createWork(0)

	n := 0
	hasValue := false
	for {
		var activeWorker chan<- int

		if hasValue{
			activeWorker = worker
		}

		select {
		case n = <- c1:
			//w <- n
			hasValue=true
			fmt.Println("Received from c1:", n)
		case n = <- c2:
			hasValue=true
			//w <- n
			fmt.Println("Received form c2:", n)
		case activeWorker <- n:
			hasValue=false
		default:
			fmt.Println("No value received")
		}

		//fmt.Println("Received from c1:", n)

	}
}
