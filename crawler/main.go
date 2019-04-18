package main

import (
	"Demo01/crawler/engine"
	"Demo01/crawler/persist"
	"Demo01/crawler/scheduler"
	"Demo01/crawler/zhenai/parser"
	"fmt"
	"regexp"
)

func main()  {

	engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan: persist.ItemSaver(),
	}.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})


}


func printCityList(contents []byte){
	re := regexp.MustCompile(`(<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

	matches := re.FindAll(contents,-1)
	//[][][]byte
	for _,m := range matches{


		fmt.Printf("%s\n",m)

		for _, subMatch :=range m{
			fmt.Printf("%s\n",subMatch)
		}

	}

}
