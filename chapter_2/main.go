package main

import (
	"fmt"
	"runtime"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) { // 指针作为传参，可以实现值的改变
	// 同时既可以使用指针调用，也可以使用对象调用
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	os := runtime.GOOS
	fmt.Println(os)
	test_int := int(10)
	fmt.Println(test_int)
	v := &Vertex{3, 4}
	s := Vertex{6, 8}
	v.Scale(5)
	s.Scale(5)
	fmt.Println(v, s)
}
