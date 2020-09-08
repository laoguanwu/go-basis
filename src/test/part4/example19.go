package main
import (
    "fmt"
    "runtime"
)
// 崩溃时需要传递的上下文信息
type panicContext struct {
    function string // 所在函数
}
// 保护方式允许一个函数
func ProtectRun(entry func()) {
    // 延迟处理的函数
    defer func() { //使用 defer 将闭包延迟执行，当 panic 触发崩溃时，ProtectRun() 函数将结束运行，此时 defer 后的闭包将会发生调用。
        // 发生宕机时，获取panic传递的上下文并打印
        err := recover()
        switch err.(type) {
        case runtime.Error: // 运行时错误
            fmt.Println("runtime error:", err)
        default: // 非运行时错误
            fmt.Println("error:", err)
        }
    }()
    entry()
}
func main() {
    fmt.Println("运行前")
    // 允许一段手动触发的错误
    ProtectRun(func() { //使用 panic 手动触发一个错误，并将一个结构体附带信息传递过去，此时，recover 就会获取到这个结构体信息，并打印出来。
        fmt.Println("手动宕机前")
        // 使用panic传递上下文
        panic(&panicContext{
            "手动触发panic",
        })
        fmt.Println("手动宕机后")
	})
	fmt.Println("分割线===========")
    // 故意造成空指针访问错误
    ProtectRun(func() { //模拟代码中空指针赋值造成的错误，此时会由 Runtime 层抛出错误，被 ProtectRun() 函数的 recover() 函数捕获到。
        fmt.Println("赋值宕机前")
        var a *int
        *a = 1
        fmt.Println("赋值宕机后")
    })
    fmt.Println("运行后")
}