package main
import (
    "fmt"
)
func main() {
    var cat int = 1 //声明整型 cat 变量。
    var str string = "banana" //声明字符串 str 变量。
    fmt.Printf("%p %p", &cat, &str) //使用 fmt.Printf 的动词%p输出 cat 和 str 变量取地址后的指针值，指针值带有0x的十六进制前缀。
}