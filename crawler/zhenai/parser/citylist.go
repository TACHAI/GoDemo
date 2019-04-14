package parser

import (
	"Demo01/crawler/engine"
	"regexp"
)

const cityListRe  = `(<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte)engine.ParseResult  {
	re := regexp.MustCompile(cityListRe)

	matches := re.FindAll(contents,-1)
	//[][][]byte

	result := engine.ParseResult{}
	for _,m := range matches{

		result.Items = append(result.Items,string(m[2]))

		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:engine.NilParser,
		})
	}
	return result;
}
