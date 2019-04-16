package parser

import (
	"Demo01/crawler/engine"
	"Demo01/crawler/model"
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"io"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var mirriageRe  = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)

func PareseProfile(contents []byte) engine.ParseResult {
	profile :=model.Profile{}

	age,err := strconv.Atoi(extractString(contents,ageRe))

	if err!=nil{
		profile.Age=age
	}

	profile.Marriage = extractString(contents,mirriageRe)


	// todo  profile.name ...
	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}

	return result
}

func extractString(contents []byte,re *regexp.Regexp)  string{
	match := re.FindSubmatch(contents)
	if len(match) >=2{
		return string(match[1])
	}else {
		return ""
	}
}

//func exString(contents []byte)string{
//	doc,err := goquery.NewDocumentFromReader(bufio.ScanRunes(contents,true))
//}