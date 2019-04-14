package main

import (
	"Demo01/crawler/engine"
	"Demo01/crawler/zhenai/parser"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main()  {

	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

	resp,err := http.Get("http://www.zhenai.com/zhenghun")
	if err!=nil{
		panic(err)
	}

	defer  resp.Body.Close()

	all,err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK{

		//utf8Reader := transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
		//all,err := ioutil.ReadAll(utf8Reader)

		if err !=nil{
			panic(err)
		}
		fmt.Printf("%s\n",all)


	}
	printCityList(all)



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
