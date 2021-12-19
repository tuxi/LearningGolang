package main

import (
	"fmt"
)

func main() {
	// if单独使用时只影响到自己对应的代码块
	score := 65
	if score >= 60 {
		fmt.Println("及格")
	}

	// 可以在if表达式中声明变量
	if score := 65; score >= 60 {
		fmt.Println("及格")
	}

	// 多个if相互使用没有影响
	if score >= 60 {
		fmt.Println("及格")
	}
	if score < 60 {
		fmt.Println("不及格")
	}
}
