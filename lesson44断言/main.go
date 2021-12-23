package main

import (
	"fmt"
)

/*
断言
- 只要实现了接口的全部方法认为这个类型属于接口类型，如果编写一个接口，这个接口中没有任何方法，这时认为所有类型都实现了这个接口，所哟Go语言中`interface{}`代表任意类型
- 如果`interface{}`作为方法参数就可以接收任意类型，但是在程序中有时需要知道这个参数到底是什么类型，这个时候就需要使用断言
- 断言使用时使用`interface{}变量.()`，()中判断是否属于的类型
```go
i interface{}
i.(Type)
```

断言的两大作用：
- 判断是否是指定的类型
- 把interface{}转换为特定的类型
*/

func main() {

	/*
		demo函数传入参数是456时，程序运行正常，输出：
			456
		demo函数传入参数是false时报错
			panic: interface conversion: interface {} is bool, not int
	*/
	demo(456)

	/*
		demo函数传入参数是456时，程序运行正常，输出：
			456 true
		demo函数传入参数是"abc"时，程序运行正常，输出：
			0 false
	*/
	demo1("abc")
}

// 断言可以有一个返回值，如果判断结果是指定类型返回值，如果不是指定类型报错
func demo(i interface{}) {
	result := i.(int)
	fmt.Println(result)
}

/*
断言也可以有两个返回值，这时无论是否指定类型都不会报错
- 第一个参数
	- 如果正确，返回值变量值
	- 如果错误，返回判断类型的默认值
- 第二个参数：
	- 返回值为bool类型，true表示正确，false表示错误
*/
func demo1(i interface{}) {
	result, ok := i.(int)
	fmt.Println(result, ok)
}
