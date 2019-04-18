package view

import (
	"Demo01/crawler/model"
	"fmt"
	template2 "html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T)  {
	template := template2.Must(
		template2.ParseFiles("xxxx.html"))

	//page := model.SearchResult{}
	//err := template.Execute(os.Stdout,page)
	//
	//if err!=nil{
	//	fmt.Print("332423")
	//}


	out,err := os.Create("xxx.html")
	page := model.SearchResult{}

	err = template.Execute(out,page)
	if err!=nil{
		fmt.Print("332423")
	}

}
