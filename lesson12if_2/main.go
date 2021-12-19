package main

import "fmt"

func main() {

	/*
		if ... else 语句
	*/

	i := 50
	if i > 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// if 多重判断

	score := 77
	if score >= 60 {
		if score >= 60 && score < 70 {
			fmt.Println("及格")
		}
		if score >= 70 && score < 80 {
			fmt.Println("中等")
		}
		if score >= 80 && score < 90 {
			fmt.Println("良好")
		}
		if score >= 90 && score <= 100 {
			fmt.Println("优秀")
		}
	} else {
		fmt.Println("不及格")
	}

	// if ... else if .... else结构
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 70 {
		fmt.Println("中等")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
