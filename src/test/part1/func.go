package main

import (  
    "fmt"
)

// func calculateBill(price int, no int) int {  
//     var totalPrice = price * no // 商品总价 = 商品单价 * 数量
//     return totalPrice // 返回总价
// }

func calculateBill(price, no int) int {  
    var totalPrice = price * no
    return totalPrice
}
func main() {  
    price, no := 90, 6 // 定义 price 和 no,默认类型为 int
    totalPrice := rectProps(price, no)
    fmt.Println("Total price is", totalPrice) // 打印到控制台上
}

func rectProps(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}