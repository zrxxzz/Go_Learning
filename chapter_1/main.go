package main

import "fmt"

func test(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}
func main() {
	// num1 := 10
	// num1 = 20
	// var num2 = 20
	// var num3 int = 30
	c, d := test(18, 20)
	fmt.Printf("%d,%d\n", c, d)
	a := []int{0, 0}
	for i, sub_a := range a {
		fmt.Printf("%d,%d", i, sub_a)
	}
}
