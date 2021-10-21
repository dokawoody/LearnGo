package main

import "fmt"

// var a int = 10
// var b int = 20
// var c float64 = 1.234
// var d float64 = 5.678

var (
	a, b = 10, 20
	c, d = 1.234, 5.678
)

func main() {
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(c * d)
	fmt.Println(c / d)
}
