package main

import "fmt"

func main() {
	i := 65
	i2 := 97
	i3 := 24352
	i4 := 0x7388

	fmt.Printf("%c\n", i)
	fmt.Printf("%c\n", i2)
	fmt.Printf("%c\n", i3)
	fmt.Printf("%c\n", i4)
	fmt.Println(i4)
}

/*
字符型概述
- 字符型存放单子字母或者单个文字
- Go 语言不支持字符类型，在Go 语言中所有字符值都转换为对应的编码表中的int32值
- Go 语言默认使用UFT-8 编码
*/
