package main

import (
	"fmt"
	"time"
)

func chanDemo()  {
	//var c chan int // c==nil
	var channels [10]chan<- int
	for i :=0;i<10;i++{
		channels[i] = createWork(i)
	}

	for i:=0;i<10;i++{
		channels[i]<- 'a'+i
	}

	for i:=0;i<10;i++{
		channels[i]<- 'A'+i
	}
	time.Sleep(time.Millisecond)
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

func channelClose()  {
	c := make(chan int,3)
	work(-1,c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	close(c)
	time.Sleep(time.Millisecond)
}

func bufferedChannel(){
	c := make(chan int,3)
	work(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func main()  {
	//chanDemo()

	//bufferedChannel()

	channelClose()
}
