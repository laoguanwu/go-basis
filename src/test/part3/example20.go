package main
import "fmt"
func main() {
OuterLoop: //外层循环的标签。
    for i := 0; i < 2; i++ {
        for j := 0; j < 5; j++ {
            switch j {
            case 2:
                fmt.Println(i, j)
                break
            case 3:
                fmt.Println(i, j)
                break OuterLoop
            }
        }
    }
}