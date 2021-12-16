package main

import (
	"fmt"
	"math"
)

func main() {
	// 整型和浮点型的转换
	var f float64 = 1
	var i int = 2
	f = float64(i)
	fmt.Println(f, i)

	// 最大值
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	// 浮点型之间相互转换
	var a float32 = 1.5
	var b float64 = 1.6
	fmt.Println(a + float32(b))
	fmt.Println(float64(a) + b)

	// 数学运算结果
	var c, d int = 3, 2
	var e, g float64 = 3, 2
	fmt.Println(c / d)
	fmt.Println(e / g)
}
