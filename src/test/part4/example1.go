package main
import "fmt"
import "math"
func main() {
	fmt.Println(hypot(3,4))

	c,d := namedRetValues()
	fmt.Println(c,d)
}

func hypot(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
} 

func namedRetValues() (a, b int) {
	a = 1
    b = 2
    return
}