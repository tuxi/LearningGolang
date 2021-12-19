package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 5
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// 使用for 循环遍历数组
	arr := [3]string{"yang", "xiaoyuan", "远"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 使用for 和 range 遍历数组
	for i, n := range arr {
		fmt.Println(i, n)
	}

	for _, n := range arr {
		fmt.Println(n)
	}
}
