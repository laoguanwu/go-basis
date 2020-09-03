Go 语言
===

# ![go head](http://www.runoob.com/wp-content/uploads/2015/06/go128.png)

# https://golang.org/

###### Created by by wilson

---

# Go 语言介绍

Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。

Go是从2007年末由Robert Griesemer, Rob Pike, Ken Thompson主持开发，后来还加入了Ian Lance Taylor, Russ Cox等人，并最终于2009年11月开源，在2012年早些时候发布了Go 1稳定版本。现在Go的开发已经是完全开放的，并且拥有一个活跃的社区。

除了OOP外，近年出现了一些小众的编程哲学，Go语言对这些思想亦有所吸收。例如，Go语言接受了函数式编程的一些想法，支持匿名函数与闭包。Go语言接受了以Erlang语言为代表的面向消息编程思想，支持goroutine和通道，并推荐使用消息而不是共享内存来进行并发编程。总体来说，Go语言是一个非常现代化的语言，精小但非常强大。

---

# Go 语言特色和特性

- 简洁、快速、安全
- 并行、有趣、开源
- 内存管理、数组安全、编译迅速
- 自动垃圾回收
- 更丰富的内置类型
- 函数多返回值
- 错误处理
- 匿名函数和闭包
- 类型和接口
- 并发编程
- 反射
- 语言交互性


---


# Go 语言用途
Go 语言被设计成一门应用于搭载 Web 服务器，存储集群或类似用途的巨型中央服务器的系统编程语言。

对于高性能分布式系统领域而言，Go 语言无疑比大多数其它语言有着更高的开发效率。它提供了海量并行的支持，这对于游戏服务端的开发而言是再好不过了。

---


# Hello Word

hello.go 文件
```md
package main //定义了包名

import "fmt"  //引入了包

func main() {   // 函数名称
    fmt.Println("Hello, World!") //表达式
}
```

执行以上代码输出:
```
$ go run hello.go 
```

此外我们还可以使用 go build 命令来生成二进制文件：
```
$ go build hello.go 
$ ls
hello    hello.go
$ hello 
Hello, World!
```

---

# 数据类型
Go 支持的基本类型：

- bool true false
- 数字类型
 int8, int16, int32, int64, int 有符号
uint8, uint16, uint32, uint64, uint 无符号
float32, float64 32 位浮点数
complex64 实部和虚部都是 float32 类型的的复数。
complex128 实部和虚部都是 float64 类型的的复数
byte 是 uint8 的别名。
rune 是 int32 的别名。
- string
---

# 变量声明
标准格式
```
var 变量名 变量类型
```
批量格式
```
var (
    a int
    b string
    c []float32
    d func() bool
    e struct {
        x int
    }
)
```
简短格式
```
名字 := 表达式
```
---
# 变量的初始化
标准格式
```
var 变量名 类型 = 表达式
var age = 29
```
推导类型的格式
```
var hp = 100
```
短变量声明并初始化
```
hp := 100
```

注意：由于使用了:=，而不是赋值的=，因此推导声明写法的左值变量必须是没有定义过的变量。若定义过，将会发生编译错误。
```
var hp int // 声明 hp 变量
hp := 10 // 再次声明并赋值
```

---
# 多重赋值
变量交换在Go中的写法

```
var a int = 100
var b int = 200
b, a = a, b
fmt.Println(a, b)

```
多重赋值时，变量的左值和右值按从左到右的顺序赋值。

例如，使用 Go 语言进行排序时就需要使用交换，代码如下：
```
type IntSlice []int //将 []int 声明为 IntSlice 类型。
func (p IntSlice) Len() int           { return len(p) } 为这个类型编写一个 Len 方法，提供切片的长度。
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] } 根据提供的 i、j 元素索引，获取元素后进行比较，返回比较结果。
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] } 根据提供的 i、j 元素索引，交换两个元素的值。
```

---
# Go语言指针详解
Go语言为程序员提供了控制数据结构的指针的能力；但是，并不能进行指针运算。通过给予程序员基本内存布局，Go语言允许你控制特定集合的数据结构、分配的数量以及内存访问模式，这些对构建运行良好的系统是非常重要的：指针对于性能的影响是不言而喻的，而如果你想要做的是系统编程、操作系统或者网络应用，指针更是不可或缺的一部分。
认识指针地址和指针类型
一个指针变量可以指向任何一个值的内存地址它指向那个值的内存地址，在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关。当然，可以声明指针指向任何类型的值来表明它的原始性或结构性；你可以在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，这里的 * 号是一个类型更改器。使用一个指针

---

引用一个值被称为间接引用。
当一个指针被定义后没有分配到任何变量时，它的值为 nil。一个指针变量通常缩写为 ptr。
每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go 语言中使用&作符放在变量前面对变量进行“取地址”操作。


---
# 常量
常量定义
```
const pi = 3.14149
```
声明一组常量
```
const (
    e  = 2.7182818
    pi = 3.1415926
)
```
iota 常量生成器 周日将对应 0，周一为 1，如此等等。
```
type Weekday int
const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

---
# Go语言模拟枚举
```
type Weapon int
const (
     Arrow Weapon = iota    // 开始生成枚举值, 默认为0
     Shuriken
     SniperRifle
     Rifle
     Blower
)
// 输出所有枚举值
fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
// 使用枚举类型并赋初值
var weapon Weapon = Blower
fmt.Println(weapon)
```
---
# Go语言type关键字（类型别名）
区分类型别名与类型定义
类型别名的写法为：
```
type TypeAlias = Type
```

区分类型别名与类型定义 请看例子5

结果显示a的类型是 main.NewInt，表示 main 包下定义的 NewInt 类型。a2 类型是 int。IntAlias 类型只会在代码中存在，编译完成时，不会有 IntAlias 类型。

在结构体成员嵌入时使用别名 请看例子6

这个例子中，FakeBrand 是 Brand 的一个别名。在 Vehicle 中嵌入 FakeBrand 和 Brand 并不意味着嵌入两个 Brand。FakeBrand 的类型会以名字的方式保留在 Vehicle 的成员中。

如果改成a.show,会报错，在调用 Show() 方法时，因为两个类型都有 Show() 方法，会发生歧义，证明 FakeBrand 的本质确实是 Brand 类型。

---
# Go语言字符串和数字的相互转换
要将整数转换成字符串，一种选择是使用fmt.Sprintf，另一种做法是用函数strconv.itoa("integer to ASCII")：
```
x := 123
y := fmt.Sprintf("%d", x)
fmt.Println(y, strconv.Itoa(x)) // "123 123'
```
FormatInt 和 FormatUint 可以按不同的进位制格式化数字：
```
fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
```
fmt.Printf 里的谓词 %b、%d、%o 和 %x 往往比 Format 函数方便，若要包含数字以外的附加信息，它就尤其有用：
```
s := fmt.Sprintf("x=%b", x)                    //"x=1111011"
```


---

strconv 包内的 Atoi 函数或 ParseInt 函数用于解释表示整数的字符串，而 ParseUint 用于无符号整数：
```
x, err := strconv.Atoi("123")                      // x 是整型
y, err := strconv.ParseInt("123", 10, 64)    // 十进制，最长为 64 位
```
ParseInt 的第三个参数指定结果必须匹配何种大小的整型；例如，16 表示 int16，而 0 作为特殊值表示 int。任何情况下，结果 y 的类型总是 int64，可将他另外转换成较小的类型。

第一课结尾

---
# 嵌入类型  
Go语言允许用户扩展或者修改已有类型的行为。这个功能对代码复用很重要，在修改已有类型以符合新类型的时候也很重要。这个功能是通过嵌入类型（type embedding）完成的。嵌入类型是将已有的类型直接声明在新的结构类型里。被嵌入的类型被称为新的外部类型的内部类型。
- 示例7
由于内部类型的标识符提升到了外部类型，我们可以直接通过外部类型的值来访问内部类型的标识符。

- 让我们修改一下这个例子，加入一个接口，示例8

- 如果外部类型并不需要使用内部类型的实现，而想使用自己的一套实现，该怎么办？让我们看另一个示例程序是如何解决这个问题的，代码如下所示。
示例9

---
# 数组

数组写法 打印数组
```
var a [3]int
for k, v := range a {
    fmt.Printf("%d %d\n", k, v)
}
```
定义后面数量决定长度的数组
```
q := [...]int{1, 2, 3}
```
比较两个数组
```
a := [2]int{1, 2}
b := [...]int{1, 2}
c := [2]int{1, 3}
fmt.Println(a == b, a == c, b == c) // "true false false"
d := [3]int{1, 2}
fmt.Println(a == d) // 编译错误：无法比较 [2]int == [3]int
```

---
# 切片详解（slice）
在 Go语言中 Slice 代表变长的序列，序列中每个元素都有相同的类型。一个 slice 类型一般写作 []T，其中 T 代表 slice 中元素的类型；slice 的语法和数组很像，只是没有固定长度而已。
从数组生成切片，代码如下：
```
var a  = [3]int{1, 2, 3}
fmt.Println(a, a[1:2])
```
从数组或切片生成新的切片拥有如下特性：
- 取出的元素数量为：结束位置-开始位置。
- 取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取。
- 当缺省开始位置时，表示从连续区域开头到结束位置。
- 当缺省结束位置时，表示从开始位置到整个连续区域末尾。
- 两者同时缺省时，与切片本身等效。
- 两者同时为0时，等效于空切片，一般用于切片复位。

---
表示原有切片
```
a := []int{1, 2, 3}
fmt.Println(a[:])
```
 重置切片，清空拥有的元素
```
a := []int{1, 2, 3}
fmt.Println(a[0:0])
```
直接声明新的切片
```
var name []T
```
使用 make() 函数构造切片
```
make( []T, size, cap )
```
T：切片的元素类型。
size：就是为这个类型分配多少个元素。
cap：预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。

---
# 切片append()
```
var a []int
a = append(a, 1) // 追加1个元素
a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包
```
切片在扩容时，容量的扩展规律按容量的 2 倍数扩充，例如 1、2、4、8、16……，代码如下：
在开头添加元素
```
var a = []int{1,2,3}
a = append([]int{0}, a...) // 在开头添加1个元素
a = append([]int{-3,-2,-1}, a...) // 在开头添加1个切片
```
由于 append 函数返回新的切片，也就是它支持链式操作。我们可以将多个 append 操作组合起来，实现在切片中间插入元素：
```
var a []int
a = append(a[:i], append([]int{x}, a[i:]...)...) // 在第i个位置插入x
a = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入切片
```


---
# copy
```
slice1 := []int{1, 2, 3, 4, 5}
slice2 := []int{5, 4, 3}
copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
```
虽然通过循环复制元素更直接，不过内置的 copy 函数使用起来更加方便。copy 函数的第一个参数是要复制的目标 slice，第二个参数是源 slice，目标和源的位置顺序和 dst = src 赋值语句是一致的。两个 slice 可以共享同一个底层数组，甚至有重叠也没有问题。

copy 函数将返回成功复制的元素的个数，等于两个 slice 中较小的长度，所以我们不用担心覆盖会超出目标 slice 的范围。

示例10 通过代码演示对切片的引用和复制操作后对切片元素的影响。

part2 结束

---
# 切片删除
从开头位置删除
```
a = []int{1, 2, 3}
a = a[1:] // 删除开头1个元素
a = a[N:] // 删除开头N个元素
```
从中间位置删除
```
a = []int{1, 2, 3, ...}
a = append(a[:i], a[i+1:]...) // 删除中间1个元素
a = append(a[:i], a[i+N:]...) // 删除中间N个元素
a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素
```
从尾部开始删
```
a = []int{1, 2, 3}
a = a[:len(a)-1] // 删除尾部1个元素
a = a[:len(a)-N] // 删除尾部N个元素
```
【示例11】删除切片指定位置的元素。


---
# range 键字
Go语言有个特殊的关键字 range，它可以配合关键字 for 来迭代切片里的元素，如下所示：
```
slice := []int{10, 20, 30, 40}
// 迭代每一个元素，并显示其值
for index, value := range slice {
    fmt.Printf("Index: %d Value: %d\n", index, value)
}
```
range 创建了每个元素的副本，而不是直接返回对该元素的引用，如下所示。 【示例12】

使用空白标识符（下划线）来忽略索引值
```
slice := []int{10, 20, 30, 40}
// 迭代每个元素，并显示其值
for _, value := range slice {
    fmt.Printf("Value: %d\n", value)
}
```
【示例 13】使用传统的 for 循环对切片进行迭代

---
# 多维切片
创建多维切片
```
slice := [][]int{{10}, {100, 200}}
```
# ![go head](http://c.biancheng.net/uploads/allimg/190614/4-1Z61416004D92.gif)

【示例14】组合切片的切片append()
```
// 为第一个切片追加值为 20 的元素
slice[0] = append(slice[0], 20)
```

---
# map
map 是引用类型，可以使用如下声明：
```
var map1 map[keytype]valuetype
var map1 map[string]int
```
- keytype 为键类型。
- valuetype 是键对应的值类型。
map 容量
```
map2 := make(map[string]float, 100)
```
用切片作为 map 的值
```
mp1 := make(map[int][]int)
mp2 := make(map[int]*[]int)
```

---
# map操作
遍历
```
for k, v := range scene {
    fmt.Println(k, v)
}
```
删除
```
delete(map, 键)
```
Go语言中并没有为 map 提供任何清空所有元素的函数、方法，清空 map 的唯一办法就是重新 make 一个新的 map，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。
```
mp1 := make(map[int]int)
```

---
# sync.Map
 map 在并发情况下，只读是线程安全的，同时读写线程不安全。
需要并发读写时，一般的做法是加锁，但这样性能并不高。Go 语言在 1.9 版本中提供了一种效率较高的并发安全的 sync.Map。sync.Map 和 map 不同，不是以语言原生形态提供，而是在 sync 包下的特殊结构。

sync.Map有以下特性：
- 无须初始化，直接声明即可。
- sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用。Store 表示存储，Load 表示获取，Delete 表示删除。
- 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值。Range 参数中的回调函数的返回值功能是：需要继续迭代遍历时，返回 true；终止迭代遍历时，返回 false。【示例14】


---
# 列表list
初始化列表
```
变量名 := list.New()
var 变量名 list.List
```
在列表中插入元素
```
l := list.New()
l.PushBack("fist")  //在开头插入
l.PushFront(67)     //在结尾插入
```
看【示例list.go】

遍历list
```
for i := l.Front(); i != nil; i = i.Next() {
    fmt.Println(i.Value)
}
```

---
# nil：空值/零值
Go语言中 nil 是一个预先定义好的标识符，它是许多类型的零值表示，有过其他编程语言开发经验的开发者也许会把 nil 看作其他语言中的 null（NULL），其实这并不是完全正确的，因为Go语言中的 nil 和其他语言中的 null 有很多不同点。
nil 标识符是不能比较的 示例【15】
nil 没有默认类型【16】
不同类型 nil 的指针是一样的【17】
nil 是 map、slice、pointer、channel、func、interface 的零值【18】
不同类型的 nil 值占用的内存大小可能是不一样的【19】
具体的大小取决于编译器和架构，上面打印的结果是在 64 位架构和标准编译器下完成的，对应 32 位的架构的，打印的大小将减半。


---
# Go语言make和new关键字的区别及实现原理
在Go语言中，make 关键字的主要作用是初始化内置的数据结构，也就是我们在前面提到的数组、切片和 Channel，而当我们想要获取指向某个类型的指针时可以使用 new 关键字，只是知道如何使用 new 的人真的比较少，下面我们就来介绍一下 make 和 new 它们的区别以及实现原理。
虽然 make 和 new 都是能够用于初始化数据结构，但是它们两者能够初始化的结构类型却有着较大的不同，make 在Go语言中只能用于初始化语言中的基本类型：
```
slice := make([]int, 0, 100)
hash := make(map[int]bool, 10)
ch := make(chan int, 5)
```
slice 是一个包含 data、cap 和 len 的结构体；
hash 是一个指向 hmap 结构体的指针；
ch 是一个指向 hchan 结构体的指针。

---
而另一个用于初始化数据结构的关键字 new 的作用其实就非常简单了，它只是接收一个类型作为参数然后返回一个指向这个类型的指针：
```
i := new(int)
var v int
i := &v
```
上述代码片段中的两种不同初始化方法其实是等价的，它们都会创建一个指向 int 零值的指针。
make 实现原理
我们已经了解了 make 在创建数组和切片、哈希表和 Channel 的具体过程，所以在这里我们也只是会简单提及 make 相关的数据结构初始化原理。
# ![go head](http://c.biancheng.net/uploads/allimg/190903/4-1ZZ31I35HJ.gif)

---
在编译期间的类型检查阶段，Go语言其实就将代表 make 关键字的 OMAKE 节点根据参数类型的不同转换成了 OMAKESLICE、OMAKEMAP 和 OMAKECHAN 三种不同类型的节点，这些节点最终也会调用不同的运行时函数来初始化数据结构。
new 
内置函数 new 会在编译期间的 SSA 代码生成阶段经过 callnew 函数的处理，如果请求创建的类型大小是 0，那么就会返回一个表示空指针的 zerobase 变量，在遇到其他情况时会将关键字转换成 newobject：

总结
最后，简单总结一下Go语言中 make 和 new 关键字的实现原理，make 关键字的主要作用是创建切片、哈希表和 Channel 等内置的数据结构，而 new 的主要作用是为类型申请一片内存空间，并返回指向这片内存的指针。

---
# if else（分支结构）
```
var ten int = 11
if ten > 10 {
    fmt.Println(">10")
} else {
    fmt.Println("<=10")
}
```
特殊写法 if 还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断，代码如下：
```
if err := Connect(); err != nil {
    fmt.Println(err)
    return
}
```
Connect 是一个带有返回值的函数，err:=Connect() 是一个语句，执行 Connect 后，将错误保存到 err 变量中。

---
# for（循环结构）
```
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```
do while 写法
```
sum := 0
for {
    sum++
    if sum > 100 {
        break
    }
}
```
初始语句简化
```
step := 2
for ; step > 0; step-- {
    fmt.Println(step)
}
```

---

条件表达式简化写法
```
var i int
for ; ; i++ {
    if i > 10 {
        break
    }
}
```
无限循环 条件也可以不写
```
var i int
for ; ; i++ {
    if i > 10 {
        break
    }
}
```
只有条件循环
```
var i int
for i <= 10 {
    i++
}
```
---
# switch break continue
switch
```
var a = "hello"
switch a {
case "hello":
    fmt.Println(1)
default:
    fmt.Println(0)
}
```
break 语句可以结束 for、switch 和 select 的代码块，另外 break 语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的 for、switch 和 select 的代码块上。 【示例20】
continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 for 循环内使用，在 continue 语句后添加标签时，表示开始标签对应的循环【21】

part3 end

---


---
# 函数

函数声明
```
func functionname(parametername type) returntype {  
    // 函数体（具体实现的功能）
}
```
示例函数
```
func calculateBill(price int, no int) int {  
    var totalPrice = price * no // 商品总价 = 商品单价 * 数量
    return totalPrice // 返回总价
}
```
如果有连续若干个参数，它们的类型一致，那么我们无须一一罗列，只需在最后一个参数后添加该类型。
```
func calculateBill(price, no int) int {  
    var totalPrice = price * no
    return totalPrice
}
```

---
多返回值
```
func rectProps(length, width float64)(float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}
```
命名返回值
```
func rectProps(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}
```
空白符
_ 在 Go 中被用作空白符，可以用作表示任何类型的任何值。

---
# 包


---