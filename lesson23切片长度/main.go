package main

import "fmt"

func main() {

	// go语言可以使用make函数创建slice 、 map、 channel、 interface
	// 使用make函数定义无内容，但是不是nil的切片，意味着切片已经申请了内存空间
	// `make(类型,初始长度[,初始容量])`
	//  初始容量可以省略,默认和长度相等

	// 长度为0切片, 没有第三个参数表示容量和长度相等
	slice := make([]string, 0)
	// 长度为0，容量为3个的切片
	slice1 := make([]string, 0, 3)
	fmt.Println(len(slice), cap(slice))
	fmt.Println(len(slice1), cap(slice1))

	// 长度表示切片中元素的实际个数，容量表示切片占用的空间大小，且切片容量成倍数增加，当增加到1024后按照一定的百分比增加
	// 长度表示切片中元素的实际个数,容量表示切片占用空间大小,且切片容量成倍增加.当增加到1024后按照一定百分比增加.
	// len(slice) 查看切片的长度
	// cap(slice) 查看切片的容量

	// append函数
	// 可以向切片中添加一个或多个值，添加后必须使用append函数的返回值接收
	fmt.Println("append函数")
	s := make([]string, 0)
	fmt.Println(len(s), cap(s))
	s = append(s, "yuanyuan", "yangyang")
	fmt.Println(len(s), cap(s)) // 2 2
	fmt.Println(s)
	s = append(s, "hello")
	fmt.Println(len(s), cap(s)) // 3 4
	fmt.Println(s)

	// 如果一次添加多个值，且添加后的长度大于扩容一次的长度，容量和长度相等
	// 等到下次添加内容时如果不超出扩容大小,在现在的基础上进行翻倍
	s1 := make([]string, 0)
	fmt.Println(len(s1), cap(s1)) // 0, 0
	s1 = append(s1, "yang", "yuan")
	fmt.Println(len(s1), cap(s1)) // 2, 2
	s1 = append(s1, "4", "5", "6", "7", "8", "9")
	fmt.Println(len(s1), cap(s1)) //输出:8, 8
	s1 = append(s1, "10")
	fmt.Println(len(s1), cap(s1)) //输出:9 16

	// 直接把另一个切片的内容添加到切片中，需要注意的是语法有三个点
	s2 := make([]string, 0)
	s3 := []string{"yang", "yuan"}
	// 将切片s3中的元素全部添加到s2中，注意结尾有...
	s2 = append(s2, s3...)
	fmt.Println(s2)
}
