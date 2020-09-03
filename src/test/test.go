package main

import (  
	"fmt"
)

func main() {  
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "hello2":
		fmt.Println(2)
		break
	default:
		fmt.Println(0)
}

	

	//fmt.Println(a[])
}

func ttt() (int, string){
	return 1,"dddd"
}