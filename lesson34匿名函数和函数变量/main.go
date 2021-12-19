package main

import "fmt"

func main() {

	/*
		一.匿名函数

		匿名函数就是没有名称的函数
		正常函数可以通过名称多次调用,而匿名函数由于没有函数名,所以大部分情况都是在当前位置声明并立即调用(函数变量除外)
		匿名函数声明完需要调用,在函数结束大括号后面紧跟小括号
		```go
		func (){

		}()//括号表示调用
		```
	*/

	// 无参数的匿名函数
	func() {
		fmt.Println("这是匿名函数")
	}()

	// 有参数的匿名函数
	func(name string) {
		fmt.Println("有参数", name)
	}("yuan")

	// 有参数有返回值的匿名函数
	age := func(name string) int {
		fmt.Println("name: ", name)
		return 18
	}("远")
	fmt.Println("age: ", age)

	/*
		# 一.函数变量

		* 在Go语言中函数也是一种类型,函数有多少种形式,函数变量就有多少种写法
		```go
			var a func()           //无参数无返回值
			var b func(int)        //有一个int类型参数
			var c func(int) string //有一个int类型参数和string类型返回值
			fmt.Println(a, b, c)   //输出:<nil> <nil> <nil>
		```
		* 定义完函数变量后,可以使用匿名函数进行赋值.也可以使用已经定义好的函数进行赋值
		* 函数变量定义以后与普通函数调用语法相同,变量名就是普通函数声明的函数名
	*/

	// 定义函数变量
	var a func()
	a = func() {
		fmt.Println("执行的函数")
	}
	// 执行函数
	a()

	// 或者使用短变量
	b := func() {
		fmt.Println("执行b函数")
	}
	b()
	fmt.Printf("%p %T \n", b, b)

	/*
		函数变量类型是除了slice、map、channel、interface外第五种引用类型
	*/

	demo(func(c string) {
		fmt.Println(c)
	})

	fn := demo1()
	fmt.Println(fn())
}

// 函数作为参数或者返回值
func demo(a func(c string)) {
	fmt.Println("执行了demo函数")
	a("传递的参数")
}

// 函数作为返回值
func demo1() func() int {
	return func() int {
		return 20
	}
}
