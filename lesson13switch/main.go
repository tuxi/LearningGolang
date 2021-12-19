package main

import "fmt"

func main() {

	// 条件判断语句switch

	// 1.当变量是固定值的可以使用switch结构
	num := 16
	switch num {
	case 2:
		fmt.Println("二进制")
	case 8:
		fmt.Println("八进制")
	case 10:
		fmt.Println("十进制")
	case 16:
		fmt.Println("十六进制")
	default:
		fmt.Println("内容不正确")
	}

	// switch 也支持在条件判断中1定义变量
	switch num := 8; num {
	case 2:
		fmt.Println("2进制")
	case 8:
		fmt.Println("8进制")
	case 10:
		fmt.Println("10进制")
	case 16:
		fmt.Println("16进制")
	default:
		fmt.Println("内容不正确")
	}

	// 当条件判断不是固定值时
	score := 70
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 70:
		fmt.Println("中等")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// case 支持多个条件的值，每个值用,分割
	month := 5
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31天")
	case 2:
		fmt.Println("28天或者29天")
	default:
		fmt.Println("30天")
	}

	// switch 穿透fallthrough,中断break
	// 穿透fallthrough，可以让当前case 执行后，还可以执行下一个case（不管其条件是否满足，都会被执行）
	switch num := 1; num {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("不是2，3，4")
	}

	// 终端break
	switch num := 1; num {
	case 1:
		fmt.Println("1")
		break
		fmt.Println("直接结束switch，后面的代码不会被执行")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("不是2，3，4")
	}
}
