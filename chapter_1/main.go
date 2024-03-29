package main

import "fmt"

func test(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

type Books struct {
	title  string
	author string
	number int
}

func main() {
	/* 基本变量 */
	// num1 := 10
	// num1 = 20
	// var num2 = 20
	// var num3 int = 30

	/* 数组 */
	// var nums [10]int
	// // var nums = [...]int{1,2,3} //另一种初始化方式
	// for sub := range nums {
	// 	sub = 0
	// 	fmt.Println(sub)
	// }

	/* 函数 */
	// c, d := test(18, 20)
	// fmt.Printf("%d,%d\n", c, d)
	// a := []int{0, 0}
	// for i, sub_a := range a {
	// 	fmt.Printf("%d,%d", i, sub_a)
	// }

	/*结构体*/
	// book := Books{number: 10} 短变量声明，上下二选一
	// var book = Books{number: 10}
	// book_ptr := &book
	// fmt.Printf("the title:%s the author:%s and the number:%d", book.title, book.author, book.number)
	// fmt.Printf("\n%d", book_ptr.number)
	// fmt.Printf("%#v", book_ptr)

	/* 切片 */
	s := []int{1, 3, 5, 7}
	sub_a := s[0:len(s)]
	for _, c := range sub_a {
		fmt.Printf("%d ", c)
	}
	sub_a = append(sub_a, 2, 4, 6, 8)
	fmt.Printf("%+v", sub_a)

}
