package main

import "fmt"

/*
# 一. 多返回值函数

* 在Go语言中每个函数声明时都可以定义成多返回值函数
* Go语言中所有的错误都是通过返回值返回的
* 声明多返回值函数的语法
```go
func 函数名(参数列表) (返回值,返回值){
  //函数体
}
```
* 调用函数的语法
```go
变量,变量:=函数名(参数)
```
* 调用函数时如果不想接收可以使用下划线占位
```go
变量,_:=函数名(参数)
```

* 理论上函数返回值个数可以无限多个,但是一般不去定义特别多个返回值(用结构体代替多返回值)

*/

func main() {

	/*
		函数的返回值可以不接收,表示执行函数
		函数的返回值如果接收,用于接收返回值的变量个数与返回值个数相同
		不想接收的使用占位符(_)占位
	*/

	name, age := demo1()
	fmt.Println(name, age)

}

func demo1() (string, int) {
	return "yuan", 18
}

func demo2() (name string, age int) {
	name = "yuan"
	age = 19
	return
}
