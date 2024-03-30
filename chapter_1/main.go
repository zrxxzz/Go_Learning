package main

import (
	"fmt"
	"strconv"
)

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

/* 接口 */
type Animal interface {
	eat()
	roll()
}

type Dog struct{}
type Cat struct{}

func (d Dog) eat() {
	fmt.Println("Dog begin eating")
}

func (d Dog) roll() {
	fmt.Println("Dog start barking")
}

func (c Cat) eat() {
	fmt.Println("Cat begin eating")
}
func (c Cat) roll() {
	fmt.Println("Meow")
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
	// s := []int{1, 3, 5, 7}
	// sub_a := s[0:len(s)]
	// for _, c := range sub_a {
	// 	fmt.Printf("%d ", c)
	// }
	// sub_a := make([]int, 10) // make 构造定义对象
	// sub_a = append(sub_a, 2, 4, 6, 8)
	// fmt.Printf("%+v", sub_a)

	/* Map */
	// map_test := make((map[int]int))
	// map_test[2] = 20
	// map_test[1] = 10
	// fmt.Printf("\n %#v\n", map_test)
	// for key, value := range map_test {
	// 	fmt.Printf("%d : %d\n", key, value)
	// }
	// delete(map_test, 1)
	// map_test[3] = 30
	// fmt.Printf("\n %#v", map_test)
	// for key, value := range map_test {
	// 	fmt.Printf("%d : %d\n", key, value)
	// }

	/* 接口测试 */
	// var animal Animal
	// animal = new(Dog)
	// animal.eat()
	// animal = new(Cat)
	// animal.roll()

	/* 类型转换 */
	var a string
	a = "123"
	num, _ := strconv.Atoi(a)
	fmt.Print(num)
}
