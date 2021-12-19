package main

import "fmt"

/*

* 讨论值传递和引用传递时,其实就是看值类型变量和引用类型变量作为函数参数时,修改形参是否会影响到实参
* 在Go语言中五个引用类型变量,其他都是值类型
  * slice
  * map
  * channel
  * interface
  * func()
* 引用类型作为参数时,称为浅拷贝,形参改变,实参数跟随变化.因为传递的是地址,形参和实参都指向同一块地址
* 值类型作为参数时,称为深拷贝,形参改变,实参不变,因为传递的是值的副本,形参会新开辟一块空间,与实参指向不同
* 如果希望值类型数据在修改形参时实参跟随变化,可以把参数设置为指针类型
*/

func main() {

	// 值类型作为形参时，是深拷贝，不会因为实参改变而改变
	name := "yuanx"
	age := 10
	demo(name, age)
	fmt.Println(name, age)

	// 引用类型作为形参时，会跟随实参的改变而改变
	s := []int{10, 20, 30}
	demo1(s)
	fmt.Println(s)

	// 如果希望值类型作为形参改变，可以把值类型指针作为参数
	index := 1
	str := "asdasd"
	demo2(&index, str)
	fmt.Println(index, str)
}

// 值船渡
func demo(name string, age int) {
	name = "yuan"
	age = 18
}

// 引用传递
func demo1(arg []int) {
	arg[len(arg)-1] = 100
}

// 指针作为参数传递
func demo2(content *int, s string) {
	*content = 10000
	s = "苏苏说"
}
