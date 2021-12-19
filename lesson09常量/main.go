package main

import (
	"fmt"
)

func main() {

	/*
		- 1.Go 语言的常量一般使用const声明
		- 2.Go语言的常量智能使用布尔型、数字型（蒸熟、浮点、复数）和字符串型
		- 3.Go语言的常量可以不指定类型，由编译器自己推断
		- 4.常量在运行时不可以改变
	*/

	/*
		常量的定义
	*/
	// 单个常量定义
	const a string = "xiaoyuan"
	const b = 1
	const c = 1.5

	fmt.Printf("查看类型：%T %T \n", b, c)
	// 常量未明确声明类型，可以相加
	fmt.Println(b + c)

	i1 := 1
	i2 := 1.5
	// 变量未明确的两个类型，不可以相加
	fmt.Println(i1 + int(i2))

	// 声明常量的表达式中不能有变量
	const d = 1*2 + 3
	i := 2
	fmt.Println(i)

	// 定义常量组
	const (
		a1 = 5
		b1 = 6
		c1 = 7
	)
	fmt.Println(a1, b1, c1)

	/*
		常量生成器 iota
		- 1.iota只能被使用在const限定中，相当于一个常量的计数器
		- 2.iota 相当于一个枚举值，默认从0开始，在一个const中，会进行+1
		- 3.iota不会因为const中被赋值了固定值，就不会增加，const中每当有一个常量就+1
		- 4.每当进入一个新的const，iota都会被重新开始计算
	*/
	// 常量生成器 iota
	const (
		a2 = iota
		b2
		c2
	)
	fmt.Println(a2, b2, c2)

	const (
		a3 = iota + 8
		b3
		c3
	)
	fmt.Println(a3, b3, c3)

	const (
		f = iota + 8
		g
		h
		o = iota
		p
		q = 1
	)
	fmt.Println(f, g, h, o, p, q)
}
