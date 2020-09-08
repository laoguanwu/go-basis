package main
import "fmt"
// 用于测试值传递效果的结构体
type Data struct { //将 Data 声明为结构体类型，结构体是拥有多个字段的复杂结构。
    complax []int // 测试切片在参数传递中的效果 complax 为整型切片类型，切片是一种动态类型，内部以指针存在。
    instance InnerData // 实例分配的innerData instance 成员以 InnerData 类型作为 Data 的成员。
    ptr *InnerData // 将ptr声明为InnerData的指针类型 将 ptr 声明为 InnerData 的指针类型。
}
// 代表各种结构体字段
type InnerData struct { //声明一个内嵌的结构 InnerData。
    a int
}

// 值传递测试函数
func passByValue(inFunc Data) Data {
    // 输出参数的成员情况
    fmt.Printf("inFunc value: %+v\n,%+v", inFunc,inFunc.ptr) // 使用格式化的%+v动词输出 inFunc 变量的详细结构，以便观察 Data 结构在传递前后内部数值的变化情况。
    // 打印inFunc的指针
	fmt.Printf("inFunc ptr: %p\n", &inFunc) // 打印传入参数 inFunc 的指针地址，在计算机中，拥有相同地址且类型相同的变量，表示的是同一块内存区域。

    return inFunc //将传入的变量作为返回值返回，返回的过程将发生值复制。
}
func main() {
    // 准备传入函数的结构
    in := Data{
        complax: []int{1, 2, 3},
        instance: InnerData{
            5,
        },
        ptr: &InnerData{6},
    }
    // 输入结构的成员情况
    fmt.Printf("in value: %+v\n", in)
    // 输入结构的指针地址
    fmt.Printf("in ptr: %p\n", &in)
    // 传入结构体，返回同类型的结构体
	out := passByValue(in)
	
	
    // 输出结构的成员情况
    fmt.Printf("out value: %+v\n", out)
    // 输出结构的指针地址
	fmt.Printf("out ptr: %p\n", &out)
}