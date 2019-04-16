package parser

import (
	"Demo01/crawler/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func PareseCity(
	contents []byte) engine.ParseResult {
	 re := regexp.Compile(cityListRe)

}