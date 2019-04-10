package main

import (
	"Demo01/functional/fib"
	"bufio"
	"errors"
	"fmt"
	"os"
)

func tryDefer()  {
	for i:=0;i<100;i++{
		defer fmt.Println(i)
		if i==30{
			panic("printed too many")
		}
	}
}

func wreterFile(fileName string)  {
	//file,err :=os.Create(fileName)
	// 创建文件
	file,err :=os.OpenFile(fileName,os.O_EXCL|os.O_CREATE,0666)

	// 通用处理文件打开错误
	err = errors.New("this is a custom error")
	if err != nil{
		fmt.Println("file already exists")

		if pathError,ok :=err.(*os.PathError);!ok{
			panic(err)
		}else {
			fmt.Println(pathError.Op,pathError.Path,pathError.Err)
		}

		return
		//panic(err)
	}
	// defer 延迟 是栈 先进后出
	defer file.Close()

	writer := bufio.NewWriter(file)
	// defer 延迟
	defer writer.Flush()

	f := fib.Fibonacci()

	for i :=0;i<20;i++{
		fmt.Fprintln(writer,f())
	}
}

func main()  {
	//tryDefer()

	wreterFile("fib.txt")
}