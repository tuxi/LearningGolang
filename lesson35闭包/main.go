package main

import "fmt"

/*
闭包
闭包不是Go语言特有的概念，在很多语言中都有闭包
闭包就是解决局部变量不能被外部访问的一种解决方案
是把函数当作返回值的一种应用
*/

func main() {

	// 总体思想为:在函数内部定义局部变量,把另一个函数当作返回值,局部变量对于返回值函数就相当于全局变量,所以多次调用返回值函数局部变量的值跟随变化

	fmt.Println("第一个测试")
	test := test1()
	fmt.Println(test()) // 1
	fmt.Println(test()) // 2
	fmt.Println(test()) // 3

	fmt.Println("第二个测试")
	f := test2()
	fmt.Printf("输出f的地址 %p", f) // 0x108bca0
	fmt.Println("f: ", f())    // 1
	fmt.Println("f: ", f())    // 2

	k := test2()
	fmt.Printf("输出k的地址 %p", k)
	fmt.Println("k: ", k())
	fmt.Println("k: ", k())
	fmt.Println("k ", k())
}

func test1() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func test2() func() int {
	i := 0
	return func() int {
		fmt.Printf("%p", &i)
		i += 1
		return i
	}
}
