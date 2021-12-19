package main

import "fmt"

func main() {

	// continue 关键字控制结束本次循环
	for i := 0; i < 5; i++ {
		fmt.Println("开始")
		if i == 2 || i == 3 {
			continue
		}
		fmt.Println("结束")
	}

	//  在双重for循环中continue默认影响最内侧循环,与最外层循环无关
	for k := 0; k < 2; k++ {
		for i := 0; i < 3; i++ {
			if i == 1 {
				continue
			}
			fmt.Println(k, i, "结束")
		}
	}

	// Go 语言执行行标签写法，可以通过定义标签让continue控制影响哪个for循环
myWork:
	for k := 0; k < 2; k++ {
		for i := 0; i < 3; i++ {
			if i == 1 {
				continue myWork
			}
			fmt.Println(k, i, "结束")
		}
	}

	// break
	// break 可以终端for循环，无论for循环还有几次执行，立即停止
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}

	// 在双重for循环中，break默认也影响到最近的for循环
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				break
			}
			fmt.Println(i, j)
		}
	}

	// break 也可以通过定义标签，控制break 对哪个for循环生效
myFor:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if j == 1 {
				break myFor
			}
			fmt.Println(i, j)
		}
	}

}
