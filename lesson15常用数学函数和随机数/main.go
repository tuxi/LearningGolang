package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
 math包中提供了基本数学常数和数学函数
 * math包提供的数学常数
```go
// Mathematical constants.
const (
	E   = 2.71828182845904523536028747135266249775724709369995957496696763
	Pi  = 3.14159265358979323846264338327950288419716939937510582097494459
	Phi = 1.61803398874989484820458683436563811772030917980576286213544862
	Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974
	SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931
	SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779
	SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038
	Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009
	Log2E  = 1 / Ln2
	Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790
	Log10E = 1 / Ln10
)

// Floating-point limit values.
// Max is the largest finite value representable by the type.
// SmallestNonzero is the smallest positive, non-zero value representable by the type.
const (
	MaxFloat32= 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
	SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)

	MaxFloat64= 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
	SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324 // 1 / 2**(1023 - 1 + 52)
)

// Integer limit values.
const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)
```

* 列举出常用的数学函数
```go
	var i, j float64 = 12.3, 9.6
	//向下取整,
	fmt.Println(math.Floor(i)) //输出:12
	//向上取整
	fmt.Println(math.Ceil(i)) //输出:13
	//绝对值
	fmt.Println(math.Abs(i)) //输出:12.3
	//返回值分别整数位和小数位,小数位可能出现误差
	num, decimal := math.Modf(i)
	fmt.Println(num, decimal)
	//返回两个变量中大的值
	fmt.Println(math.Max(i, j)) //输出:12.3
	//返回两个变量中小的值
	fmt.Println(math.Min(i, j)) //输出:9.6
	//x的y次方
	fmt.Println(math.Pow(3, 2)) //输出:输出9
	//四舍五入
	fmt.Println(math.Round(i)) //输出:12
```

*/

func main() {
	// 常用的数学函数
	var i, j float64 = 12.3, 9.6
	// 向下取整
	fmt.Println(math.Floor(i)) // 输出 12
	// 向上取整
	fmt.Println(math.Ceil(i)) // 输出13
	// 绝对值
	fmt.Println(math.Abs(-3)) // 输出 3
	// 返回值分别整数位和小数位，小数位可能出现误差
	num, decimal := math.Modf(i)
	fmt.Println(num, decimal)
	// 返回两个变量的最大值
	fmt.Println(math.Max(i, j)) // 出书12.3
	// 返回两个变量的最小值
	fmt.Println(math.Min(i, j)) // 出书9.6
	// x的n次方
	fmt.Println(math.Pow(3, 2)) // 输出9
	// 四舍五入
	fmt.Println(math.Round(i)) // 输出12

	// 随机数
	// math/rand实现了伪随机数生成器
	// 在go语言中随机数需要设置种子，如果不设置，每次运行获取的随机数都是相同的
	// 默认种子是1，以及相同种子，每次产生的随机数都是相同的
	// 可以使用当前时间的纳秒差计算随机数，在一定程度上保持了种子的唯一性
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int63n(10))
}
