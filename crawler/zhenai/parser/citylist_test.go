package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	//contents,err :=fetcher.Fetch("http://www.zhenai.com/zhenghun")

	contents,err :=ioutil.ReadFile("citylist_test_data.html")
	if err != nil{
		panic(err)
	}


	//ParserCityList(contents)

	fmt.Printf("%s\n",contents)
	// verify result
}
