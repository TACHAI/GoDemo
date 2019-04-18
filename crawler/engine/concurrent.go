package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}


type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	//ConfigureMasterWorkerChan(chan Request)
	//WorkerReady(chan Request)
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)
	out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()
	for i :=0;i<e.WorkerCount;i++ {
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}

	for _,r :=range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount :=0

	for{
		result := <- out
		for _, item := range result.Items{
			fmt.Printf("Got item: %v",item)
			itemCount++

			go func() {
				e.ItemChan <- item
			}()
		}

		for _,request := range result.Requests {
			e.Scheduler.Submit(request)
		}

	}

}

func createWorker(in chan Request,out chan ParseResult,ready ReadyNotifier)  {
	go func() {
		for {
			// tell scheduler i`m ready
			ready.WorkerReady(in)
			request := <- in
			result,err := worker(request)
			if err !=nil{
				continue
			}

			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool{
	if visitedUrls[url]{
		return true
	}

	visitedUrls[url] = true
	return false
}