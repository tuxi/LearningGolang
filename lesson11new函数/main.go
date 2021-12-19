package main

import "fmt"

func main() {

	// new 函数可以直接创建一个类型的指针
	a := new(int)
	fmt.Println(a)
	*a = 10
	fmt.Println(*a)

	// 只声明的指针变量时空指针，是不能通过*a赋值的，需要先赋值指针指向
	var b *int
	// 必须先赋值指向的地址
	*b = 1000
	fmt.Println(b)
}
