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
	/* 数组 */
	var nums [10]int
	// var nums = [...]int{1,2,3} //另一种初始化方式
	for sub := range nums {
		sub = 0
		fmt.Println(sub)
	}
	/* 函数 */
	c, d := test(18, 20)
	fmt.Printf("%d,%d\n", c, d)
	a := []int{0, 0}
	for i, sub_a := range a {
		fmt.Printf("%d,%d", i, sub_a)
	}
}
