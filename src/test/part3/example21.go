package main
import "fmt"
func main() {
    for i := 0; i < 2; i++ {
        for j := 0; j < 5; j++ {
            fmt.Println(i, j)
            switch j {
            case 2:
                fmt.Println("case2:",i, j)
                continue
            }
        }
    }
}