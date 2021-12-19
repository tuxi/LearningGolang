package main

import "fmt"

func main() {

	// 通过数组产生切片
	// 定义数组后，取出数组中的某个片段，这个片段就是片段slice类型
	names := [3]string{"yang", "xiaoyuan", "sss"}
	s := names[0:2]
	fmt.Printf("%T \n", s) // 输出 []]string 类型
	fmt.Println(s)         // 输出 [yang xiaoyuan]

	// 切片是指针，指向数组元素的地址，修改切片的内容，数组的内容会跟随改变
	s[0] = "Go语言"
	fmt.Println(s)     // [Go语言 xiaoyuan]
	fmt.Println(names) // [Go语言 xiaoyuan sss]

	/*
		当切片内容增加时
		1.如果增加后切片的长度没有超出数组，修改切片也是在修改数组
		2.如果增加后切片的长度超出数组，会重新开辟一块空间放切片的内容
		3，通过下面代码也正面了切片中内容存在一块连续的空间（和数组一样）
	*/

	names1 := [3]string{"杨", "校园", "sda"}
	s1 := names1[0:2]
	fmt.Printf("%p %p \n", s1, &names1[0]) // 0xc00007a1e0 0xc00007a1e0
	s1[0] = "区块链"
	s1 = append(s1, "go语言")
	fmt.Println(s1)                        // [区块链 校园 go语言]
	fmt.Println(names1)                    // [区块链 校园 go语言]
	fmt.Printf("%p %p \n", s1, &names1[0]) // 地址相同 0xc0000981e0 0xc0000981e0

	// 当超出数组长度的情况
	s1 = append(s1, "超出了数组长度")
	fmt.Println(s1)
	fmt.Println(names1)
	// 地址不同，由于超出了原数组的长度，切片重新开辟了内容空间
	fmt.Printf("%p %p \n", s1, &names1[0]) // 0xc00006e180 0xc00007a1e0

	// 切片删除
	// Go 语言标准库中没有提供删除的函数
	// 切片也可以取其中的一段形成子切片，利用这个特性可以实现删除效果
	num := []int{1, 2, 3, 4, 5, 6}
	// 要删除脚标为n的元素
	n := 2
	// 先取出n前面的所以元素
	num1 := num[0:n]
	// 在把n后面的所有元素和前面的元素添加到一起，这样就忽略了n这个元素，实现了删除n
	num1 = append(num1, num[n+1:]...)
	fmt.Println(num1)

}
