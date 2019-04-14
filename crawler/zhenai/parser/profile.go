package parser

import (
	"Demo01/crawler/engine"
	"Demo01/crawler/model"
	"regexp"
	"strconv"
)

const ageRe = `<td><span class="label">年龄：</span>([\d]+)岁</td>`
const mirriageRe  = ``

func PareseProfile(contents []byte) engine.ParseResult {
	profile :=model.Profile{}

	re := regexp.MustCompile(ageRe)
	match := re.FindSubmatch(contents)


	if match != nil{
		age,err := strconv.Atoi(string(match[1]))
		if err != nil{
			// user age is age
			profile.Age=age
		}
	}
	re = regexp.MustCompile(mirriageRe)
	match = re.FindSubmatch(contents)
	if match !=nil{
		profile.Marriage = string(match[1])
	}

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