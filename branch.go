package main

import (
	"fmt"
)


func swap(a,b *int){
	*a,*b=*b,*a;
}

func arr()  {
	var arr [5]int
	 arr2 := [5]int{1,2,3,4,5,}
	 arr3 := [...]int{1,2,3,4,5,}

	 fmt.Println(arr,arr2,arr3)
}


func maps(){
	m := map[string]string {
		"name": "daling",
		"courser": "goler",
		"site": "imooc",
		"quality": "notepad",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println("Traversing map")

	for k:= range m{

	}
}


func main() {
	//const filename = "abc.txt"
	//
	//if contents, err := ioutil.ReadFile(filename);err != nil{
	//	fmt.Println(err)
	//}else {
	//	fmt.Println("%s\n",contents)
	//}
	//if err != nil {
	//	fmt.Println(contents)
	//}else {
	//	fmt.Println("%s\n",contents)
	//
	//}

	a,b:=3,4
	swap(&a,&b)
	fmt.Println(a,b)
}
