package main

import "fmt"

func main() {

	/*
		数组
		1.具有固定长度相同类型元素序列
		2.内存是连续的
		3.声明数组后就会在内存中开辟一份连续的内存空间，每个值为数组的元素，且元素值为类型对应的默认值
		4.数组中每个元素按顺序都有自己对应的脚标，脚标第一个为0后依次+1
		5.数组在实际开发中充当临时容器，因为声明一个数组变量比声明多个相同类型的变量在操作时方便
	*/

	// 数组的创建和赋值
	// 完整写法
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	// 短变量写法
	arr1 := [3]int{3, 4, 5}
	fmt.Println(arr1)

	// 声明4个元素，初始3个元素，其余元素均为默认值0
	arr2 := [4]int{6, 7, 8}
	fmt.Println(arr2)

	// 赋值时不写长度，数组长度根据元素个数确定
	arr3 := [...]int{10, 11, 12, 13, 11}
	fmt.Println(arr3)

	// 通过脚标对数组中的y元素取值
	arr4 := [3]int{1, 2, 3}
	arr4[2] = 10
	arr4[1] = 11
	fmt.Println(arr, arr[1], arr[0])

	// 通过len 获取数组的长度
	fmt.Println(len(arr4))
	// arr4[5] = 19 // 越界 invalid array index 5 (out of bounds for 3-element array)

	// 数组是值类型的
	// 在go语言中数组是值类型的，把一个数组的变量赋值给其他数组变量是复制副本，重新开辟一块内存空间
	arr5 := [3]int{1, 2, 3}
	arr6 := arr5
	fmt.Println(arr5, arr6)
	fmt.Printf("%p, %p \n", &arr5, &arr6) //地址不同  xc0000b6048, 0xc0000b6060
	fmt.Println(arr5 == arr6)             // 值相同
	arr5[0] = 11
	fmt.Println(arr5 == arr6) // 值不相同
}
