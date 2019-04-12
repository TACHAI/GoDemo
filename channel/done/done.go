package main

import (
	"fmt"
	"sync"
	"time"
)

func chanDemo()  {
	//var c chan int // c==nil
	var wg sync.WaitGroup
	var workers [10]worker
	for i :=0;i<10;i++{
		workers[i] = createWork(i,&wg)
	}

	wg.Add(20)

	for i:=0;i<10;i++{
		workers[i].in<- 'a'+i
	}

	for i:=0;i<10;i++{
		workers[i].in<- 'A'+i

	}

	wg.Wait()
	//for _ ,worker :=range workers{
	//	<- worker.done
	//	<- worker.done
	//}
	//time.Sleep(time.Millisecond)
}

func createWork(id int,wg *sync.WaitGroup) worker {
	w := worker{
		in : make(chan int),
		wg: wg,
	}
	//c := make(chan int)
	//go func() {
	//	for{
	//		n :=<-c
	//		fmt.Printf("worker %d received %c\n",id,n)
	//
	//	}
	//}()

	go doWork(id,w.in,wg)

	return w
}

func doWork(id int,c chan int,wg *sync.WaitGroup){
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
			wg.Done()
		}
	}()
}

type worker struct {
	in chan int
	//done chan bool
	wg *sync.WaitGroup
}

func channelClose()  {
	c := make(chan int,3)
	//doWork(-1,c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	close(c)
	time.Sleep(time.Millisecond)
}

func bufferedChannel(){
	c := make(chan int,3)
	//work(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func main()  {
	chanDemo()

	//bufferedChannel()

	//channelClose()
}
