package engine

import (
	"Demo01/crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine)Run(seeds ...Request)  {
	var requests []Request
	for _,r := range seeds{
		requests = append(requests,r)
	}

	for len(requests)>0{
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetchering %s",r.Url)
		//body,err := fetcher.Fetch(r.Url)
		parseResult,err := worker(r)
		if err != nil{
			log.Printf("Fetcher: error"+"fetcher url %s: %v",
				r.Url,err)
			continue
		}

		//parseResult := r.ParserFunc(body)
		// parseResult.Requests... 将parseResult 展开加进去
		requests = append(requests,parseResult.Requests...)

		for _, item := range parseResult.Items{
			log.Printf("Got item %s",item)
		}

	}

}

func worker(r Request)(ParseResult,error)  {
	log.Printf("Fetching %s",r.Url)
	body,err := fetcher.Fetch(r.Url)
	if err != nil{
		log.Printf("Fetcher: error"+"fetching url %s: %v",r.Url,err)
		return ParseResult{},err
	}
	return r.ParserFunc(body),nil
}