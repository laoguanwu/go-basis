package main
import (
    "fmt"
)

func main() {
    func(data int) {
		fmt.Println("hello", data)
	}(100)

	// f := func(data int) {
	// 	fmt.Println("hello", data)
	// }
	// // 使用f()调用
	// f(100)
}