package main

import "fmt"

func zz(){
	var a int =2
	var pa *int  =&a
	*pa = 3
	fmt.Println(a)
	b := 3
	var d  = &b
	*d = 5
	fmt.Println(b)
}


func swap(a,b *int){
	*a,*b=*b,*a
}

func main(){
	zz()
	a:=3
	b:=7

	swap(&a,&b)
	fmt.Println(a,b)
}



