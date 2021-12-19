package main

import "fmt"

/*
# 一. 可变参数函数

* Go语言支持可变参数函数
* 可变参数指调用参数时,参数的个数可以是任意个
* 可变参数必须在参数列表最后的位置,在参数名和类型之间添加三个点表示可变参数函数
```go
func 函数(参数,参数,名称 ... 类型 ){
	//函数体
}
```
* 输出语句就是可变参数函数,源码如下
```go
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```
* 声明函数时,在函数体把可变参数当作切片使用即可

*/

func main() {

	demo("yuan", "hello", "yayayya")
}

func demo(name string, hover ...string) {
	fmt.Println("name: ", name)
	for i, value := range hover {
		fmt.Println(i, value)
	}
}
