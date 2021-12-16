package main

import (
	"fmt"
)

func main() {

	// 1.先声明变量后赋值
	var name string
	name = "ivan"
	fmt.Println(name)

	// 2.声明同时赋值
	var name1 string = "ivan1"
	fmt.Println(name1)

	// 3.声明同时赋值（省略类型）
	var name2 = "ivan2"
	fmt.Println(name2)
	fmt.Printf("%T", name)

	// 4.短变量
	name3 := "ivan3"
	fmt.Println(name3)

	// 5.一次生命多个变量
	var a, b, c int
	a, b, c = 1, 2, 3
	fmt.Println(a, b, c)

	// 6.声明并赋值
	var a1, b1, c1 = 1, 2, false
	fmt.Println(a1, b1, c1)

	// 7.使用var 声明并赋值
	var (
		a2 = 1
		b2 = 2
		c2 = false
	)
	fmt.Println(a2, b2, c2)

	// 8.使用短变量声明并赋值多个变量时，要求必须至少有一个变量没有声明
	var (
		a3 = 1
		b3 = 2
		c3 = false
	)
	b3, c3, d := 5, true, "ddd"
	fmt.Println(a3, b3, c3, d)
}
