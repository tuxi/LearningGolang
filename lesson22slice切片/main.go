package main

import "fmt"

/*
切片的英文是slice
切片：具有可变长度相同类型的元素序列
由于长度是可变的，可以解决数组长度在数据个数不确定的情况下浪费内存的问题
切片和数组声明时最主要的区别就是长度
切片是引用类型
*/

func main() {

	// 只声明切片，只声明切片时未分配内存空间
	var slice []string
	// 只声明数组，只声明数组时初始化时分配了内存空间
	var array [3]string
	fmt.Println(slice, array)
	fmt.Println(slice == nil) // nil
	fmt.Printf("%p\n", slice) // 0x0

	// 通过直接指定初始值初始化一个切片变量
	names := []string{"yang", "xiaoyuan", "yuan"}
	fmt.Println(names)

	// 定义完切片后，可以通过`切片对象[脚标]`，取出和修改脚标对应的元素内容，语法和数组相同

	// 切片是引用类型
	// 引用类型在变量之间赋值时传递的是内存地址，引用类型变量就是这个类型的指针，而切片就是引用类型
	// 值类型在变量之间传递的是值的副本
	names1 := names
	names1[0] = "杨"
	fmt.Println(names, names1)            // 输出的内容相同： [杨 xiaoyuan yuan] [杨 xiaoyuan yuan]
	fmt.Printf("%p %p \n", names1, names) // 内存地址相同

}
