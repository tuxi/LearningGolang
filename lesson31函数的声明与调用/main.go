package main

import "fmt"

/*
# 一. 函数

* 函数:一段代码块
* 所有的流程控制代码和业务代码都只能写在函数内部
* 为什么使用函数:
  * 把一个功能提出作为一个函数便于后期维护,结构清晰
* 函数声明的语法
  * 函数声明后不会执行,必须调用后才会执行
```
func 函数名(参数列表) 返回值{
  //函数体
}
```
* 函数调用的语法
```go
返回值:=函数名(参数)
```
*/

func main() {
	demo1()
	show("yuan", 18)
	fmt.Println(add(10, 20))
	fmt.Println(add2(10, 20))
}

// 无参数的函数
func demo1() {
	fmt.Println("执行了demo1")
}

// 有参数的函数
func show(name string, age int) {
	fmt.Println("姓名：", name, "年龄：", age)
}

// 有返回值的函数，函数结尾必须有return
func add(a int, b int) int {
	return a + b
}

// 也可以在返回值类型前面写变量名称， 结尾必须有renturn
func add2(a int, b int) (sum int) {
	sum = a + b
	return
}
