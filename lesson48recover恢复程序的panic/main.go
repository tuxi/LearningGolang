package main

import "fmt"

/*
# 一.recover

* recover()表示恢复程序的panic(),让程序正常运行
* recover()是和panic(v)一样都是builtin中函数,可以接收panic的信息,恢复程序的正常运行
```go
// The recover built-in function allows a program to manage behavior of a
// panicking goroutine. Executing a call to recover inside a deferred
// function (but not any function called by it) stops the panicking sequence
// by restoring normal execution and retrieves the error value passed to the
// call of panic. If recover is called outside the deferred function it will
// not stop a panicking sequence. In this case, or when the goroutine is not
// panicking, or if the argument supplied to panic was nil, recover returns
// nil. Thus the return value from recover reports whether the goroutine is
// panicking.
func recover() interface{}
```
*/

func main() {
	// demo1()

	fmt.Println("程序开始")
	demo2()
	fmt.Println("程序结束")

	/*
		执行顺序
		程序开始
		demo2上半部分
		demo3上半部分
		demo4上半部分
		demo2下半部分
		程序结束
	*/
}

// recover()一般用在defer内部,如果没有panic信息返回nil,如果有panic,recover会把panic状态取消
func demo1() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("出现了panic，使用recover获取错误信息：", error)
		}
	}()
	fmt.Println("1111111111")
	panic("出现了panic")
	fmt.Println("2222222222")

	/*
		执行日志：
		1111111111
		出现了panic，使用recover获取错误信息： 出现了panic
	*/
}

/*
函数调用过程中panic和recover()
recover() 只能恢复当前函数级或者当前函数调用函数中的panic(), 恢复后调用当前级别函数结束，但是调用次函数的函数可以继续执行

panic会一致向上传递，如果没有recover() 则表示终止程序，但是碰见了recover(),recover()所在级别函数表示没有panic，panic就不会想上传递
*/
func demo2() {
	fmt.Println("demo2上半部分")
	demo3()
	fmt.Println("demo2下半部分")
}

func demo3() {
	defer func() {
		recover() // 此处进行恢复
	}()
	fmt.Println("demo3上半部分")
	demo4()
	fmt.Println("demo3下半部分")
}

func demo4() {
	fmt.Println("demo4上半部分")
	panic("在demo4出现了panic")
	fmt.Println("demo4下半部分")
}
