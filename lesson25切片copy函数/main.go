package main

import "fmt"

/*
# 一.copy函数

* 通过copy函数可以把一个切片内容复制到另一个切片中
* Go语言标准库源码定义如下
  * 第一个参数是目标切片,接收第二个参数内容
  * 第二个参数是源切片,把内容拷贝到第一个参数中
```go
// The copy built-in function copies elements from a source slice into a
// destination slice. (As a special case, it also will copy bytes from a
// string to a slice of bytes.) The source and destination may overlap. Copy
// returns the number of elements copied, which will be the minimum of
// len(src) and len(dst).
func copy(dst, src []Type) int
```
* 拷贝时严格按照脚标进行拷贝.且不会对目标切片进行扩容
*/

func main() {

	// 将短切片拷贝到长切片中
	s1 := []int{1, 2}
	s2 := []int{3, 4, 5, 6}
	copy(s2, s1)
	fmt.Println(s1) // [1 2]
	fmt.Println(s2) // [1 2 5 6]

	// 把长切片拷贝到短切片中
	s3 := []int{1, 2}
	s4 := []int{3, 4, 5, 6}
	// 将长切片拷贝到短切片中，
	copy(s3, s4)
	fmt.Println(s3) // [3 4]
	fmt.Println(s4) // [3 4 5 6]

	// 把切片片段拷贝到切片中
	s5 := []int{1, 2}       // [4 5]
	s6 := []int{3, 4, 5, 6} // [3 4 5 6]
	copy(s5, s6[1:])
	fmt.Println(s5)
	fmt.Println(s6)

	// 使用copy完成删除
	s7 := []int{1, 2, 3, 4, 5, 6, 7}
	// 要删除的元素索引
	n := 2
	newSlice := make([]int, n)
	copy(newSlice, s7[0:n])
	newSlice = append(newSlice, s7[n+1:]...)
	fmt.Println(s7)       // 原切片长度不变 [1 2 3 4 5 6 7]
	fmt.Println(newSlice) // 新切片删除了n位置的元素 [1 2 4 5 6 7]
}
