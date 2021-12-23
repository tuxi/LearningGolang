package main

import "fmt"

/*
# 一.defer使用

* Go语言中defer可以完成延迟功能,当前函数执行完成后执行defer功能
* defer最常用的就是关闭连接(数据库连接,文件等)可以打开连接后代码紧跟defer进行关闭,后面在执行其他功能
  * 在很多语言中要求必须按照顺序执行,也就是必须把关闭代码写在最后,但是经常会忘记关闭导致内存溢出,而Golang中defer很好的解决了这个问题.无论defer写到哪里都是最后执行
*/

func main() {
	demo1()
	demo2()

	fmt.Println("测试有返回值，但没有定义返回值变量的函数")
	i := demo3(0)
	fmt.Println(i) // 输出0

	fmt.Println("声明接收返回值变量，执行defer时修改了返回值内容")
	i1 := demo4()
	fmt.Println(i1) // 输出2
}

// 单个defer
func demo1() {
	fmt.Println("打开连接")
	defer fmt.Println("关闭连接")
	fmt.Println("执行连接")
	/*
		以上输出
		打开连接
		执行连接
		关闭连接
	*/
}

/*
多个defer
多重defer，采用栈结构，先产生后执行
在很多代码结构中，都可能会出现多个对象，而程序希望这些对象倒序关闭，多个defer正好解决这个问题
*/
func demo2() {
	fmt.Println("打开连接A")
	defer func() {
		fmt.Println("关闭连接A")
	}()

	fmt.Println("打开连接B")
	defer func() {
		fmt.Println("关闭连接B")
	}()

	fmt.Println("打开连接C")
	defer func() {
		fmt.Println("关闭连接C")
	}()

	/*
		输出
		 开连接A
		打开连接B
		打开连接C
		关闭连接C
		关闭连接B
		关闭连接A
	*/
	// 实现先打开的后关闭
}

/*
defer 和 return 组合
defer 与 return 同时存在，要把return 理解成两条执行结合（不是原子指令），一个指令是给返回值赋值，另一个指令返回跳出函数
defer 和 return 时整体执行顺序
	- 先给返回值赋值
	- 执行defer
	- 返回跳出函数
*/

// 没有定义返回值接收的变量，执行defer时返回值已经赋值
func demo3(i int) int {
	defer func() {
		i = i + 2
	}()
	return i // i 未被修改
}

// 声明接收返回值变量，执行defer时修改了返回值内容
// 由于return 后面没有内容，就无法黑返回值赋值，所以执行defer时返回值才有内容
func demo4() (i int) {
	defer func() {
		i = i + 2
	}()
	return // 2
}
