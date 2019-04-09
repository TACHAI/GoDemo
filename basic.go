package main

import (
	"fmt"
	"math"
	"strconv"
)


var aa = 2


func variableZeroValue()  {
	var a int
	var s string
	fmt.Println("%d %q\n",a,s)
}

func variableInitalValue()  {
	var a, b  = 3, 4
	var s  = "abc"
	fmt.Println(a,b,s)
}

func varialbeTypeDeduction()  {
	var a,b,c, s = 3,4,true,"def"
	fmt.Println(a,b,c,s)
}

func varialbeShorter(){
	a,b,c, s := 3,4,true,"def"

	b = 5
	g:="sdfsdf"
	fmt.Println(a,b,c,s,g)
}




func consts(){
	const filename = "abc.txt"
	const a, b =3,4
	var c int
	c = int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}


func enums(){
	//const(
	//	cpp =0
	//	java =1
	//	python =2
	//	golang =3
	//)
	const(
		cpp =iota
		java
		python
		golang
	)

	fmt.Println(cpp,java,python,golang)

}

func sum(index int) int {
	sum:=0
	for i:=0;i<index;i++{
		sum+=i
	}
	return sum
}

func converToBin(n int) string{
	if n==0{
		return "0"
	}
	s:=""
	for ;n>0;n/=2{
		m:=n%2
		s=strconv.Itoa(m)+s
	}
	return s
}


func main()  {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitalValue()
	varialbeTypeDeduction()

	consts()
	enums()


	fmt.Println(sum(100))
	//fmt.Println(
	//	converToBin(5),
	//	converToBin(13),
	//
	//	)
	//converToBin(100)
	fmt.Println(converToBin(0))
}