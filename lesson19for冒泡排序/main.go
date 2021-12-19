package main

import "fmt"

func main() {

	/*
		冒泡排序
		排序就是把一组数组数据按照特定的顺序重新排列，可以是升序、降序
		冒泡排序利用双重for循环把最大(小)的值移动到一侧,每次可以判断出一个数据,如果有n个数组,执行n-1次循环就可以完成排序
		排序代码(升序还是降序主要是看if判断是大于还是小于)
	*/
	arr := [5]int{1, 3, 7, 2, 8}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}
