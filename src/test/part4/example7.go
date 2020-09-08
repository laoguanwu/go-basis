package main
import (
    "flag"
    "fmt"
)
var skillParam = flag.String("skill", "", "skill to perform") //定义命令行参数 skill，从命令行输入 --skill 可以将=后的字符串传入 skillParam 指针变量。
func main() {
    flag.Parse() //解析命令行参数，解析完成后，skillParam 指针变量将指向命令行传入的值。
    var skill = map[string]func(){ //定义一个从字符串映射到 func() 的 map，然后填充这个 map。
        "fire": func() { 
            fmt.Println("chicken fire") // 初始化 map 的键值对，值为匿名函数。
        },
        "run": func() {
            fmt.Println("soldier run")
        },
        "fly": func() {
            fmt.Println("angel fly")
        },
    }
    if f, ok := skill[*skillParam]; ok { //skillParam 是一个 *string 类型的指针变量，使用 *skillParam 获取到命令行传过来的值，并在 map 中查找对应命令行参数指定的字符串的函数。
		fmt.Println(ok)
		f()
    } else {
		fmt.Println(ok)
        fmt.Println("skill not found") //如果在 map 定义中存在这个参数就调用，否则打印“技能没有找到”。
    }
}