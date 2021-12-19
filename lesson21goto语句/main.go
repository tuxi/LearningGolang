package main

import "fmt"

func main() {

	/*
		goto是go语言的一个关键字
		goto让编译器执行时跳转到特定的位置
	*/

	// goto语句
	fmt.Println("执行程序")
	i := 6
	if i == 6 {
		goto Loop
	}

Loop:
	fmt.Println("loop")

	// goto也可以跳出循环，执行指定标签位置代码
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			goto abc
		}
	}

	fmt.Println("for循环执行结束")
abc:
	fmt.Println("abc")
	fmt.Println("程序结束")
}
