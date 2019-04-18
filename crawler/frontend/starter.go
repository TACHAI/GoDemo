package main

import (
	"Demo01/crawler/frontend/controller"
	"net/http"
)

func main()  {
	http.Handle("/", http.FileServer(
		http.Dir("crawler/frontend/view")))
	http.Handle("/search", controller.SearchResultHandler{})
	err :=http.ListenAndServe(":8888",nil)
	if err!=nil{
		panic(err)
	}
}
