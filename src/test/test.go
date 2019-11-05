package main

import (  
	"fmt"
)

func main() {  
	var a  = [3]int{1, 2, 3}
	fmt.Println(a[1:])
	
	fmt.Println(a[:2])
	
	fmt.Println(a[:])
	b := a[:]
	
	fmt.Println(b)
	fmt.Println(b[0:0])
	fmt.Println(b)

	

	//fmt.Println(a[])
}

func ttt() (int, string){
	return 1,"dddd"
}