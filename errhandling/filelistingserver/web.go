package main

import (
	"Demo01/errhandling/filelistingserver/filelisting"
	"fmt"
	"net/http"
	"os"
)


type appHandler func(writer http.ResponseWriter,request *http.Request)error

func errWrapper(
	handler appHandler)func(writer http.ResponseWriter,request *http.Request){

		return func(writer http.ResponseWriter, request *http.Request) {
			err := handler(writer,request)
			if err!=nil{

				code := http.StatusOK

				switch {
				case os.IsNotExist(err):
					code = http.StatusNotFound
					http.Error(
						writer,
						http.StatusText(http.StatusNotFound),
						http.StatusNotFound)
				case os.IsPermission(err):
					code = http.StatusForbidden
				default:
					code = http.StatusInternalServerError
				}
				http.Error(writer,http.StatusText(code),code)
			}
		}

}

func main()  {
	http.HandleFunc("/list/",errWrapper(filelisting.HandlerFileList))


	err := http.ListenAndServe(":8888",nil)
	if err !=nil{
		fmt.Println(err)
	}
}
