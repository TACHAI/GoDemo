package main

import (
	"fmt"
	"regexp"
)

const name  = "my email is 23234234@gmail.com"

func main()  {
	re :=regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+.[a-zA-Z0-9]+`)
	//re :=regexp.MustCompile(".+@.+\\..+")

	match :=re.FindAllString(name,-1)
	//submatch := re.FindStringSubmatch(name)
	for i:=0;i<len(match);i++{
		fmt.Println(match[i])
	}


	for m:=range match{
		fmt.Println(m)

	}
}