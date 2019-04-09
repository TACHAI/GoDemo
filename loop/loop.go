package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func printFile(fileName string)  {
	file, err :=os.Open(fileName)
	
	if err!=nil{
		panic(err)
	}
	
	
	printFileContents(file)
	//scanner := bufio.NewScanner(file)
	//
	//for scanner.Scan(){
	//	fmt.Println(scanner.Text())
	//}
	
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main()  {
	s:=`"sdfsdfds"
	123
	erdsg

	p`

	printFileContents(strings.NewReader(s))
}