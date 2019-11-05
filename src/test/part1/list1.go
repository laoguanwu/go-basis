
package main
import "container/list"
import "fmt"
func main() {
    l := list.New()
    // 尾部添加
	l.PushBack("canon")
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("================")
    // 头部添加
	l.PushFront(67)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("================")
    // 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("================")
    // 在fist之后添加high
	l.InsertAfter("high", element)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("================")
    // 在fist之前添加noon
	l.InsertBefore("noon", element)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("================")
    // 使用
	l.Remove(element)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}