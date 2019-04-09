package main

import "fmt"

func adder()func(int)int  {
	sum :=0
	return func(i int) int {
		sum = i+sum
		return sum
	}
}

//++++ 这里是正统函数式编程
type iAdder func(int) (int,iAdder)

func adder2(base int)iAdder  {
	return func(v int) ( int,  iAdder) {
		return base +v,adder2(base+v)
	}
}
//+++


func main(){
	a := adder()
	for i:=0;i<10;i++{
		fmt.Println(a(i))
	}

}