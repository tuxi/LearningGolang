package main

import (
	"errors"
	"fmt"
)

/*
# 一. 错误

* 在程序执行过程中出现的不正常情况称为错误
* Go语言中使用builtin包下error接口作为错误类型,官方源码定义如下
  * 只包含了一个方法,方法返回值是string,表示错误信息
```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
```

* Go语言中错误都作为方法/函数的返回值,因为Go语言认为使用其他语言类似try...catch这种方式会影响到程序结构
* 在Go语言标准库的errors包中提供了error接口的实现结构体errorString,并重写了error接口的Error()方法.额外还提供了快速创建错误的函数
```go
package errors

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

```
* 如果错误信息由很多变量(小块)组成,可以借助`fmt.Errorf("verb",...)`完成错误信息格式化,因为底层还是errors.New()
```go
// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
}
```
*/

func main() {
	result, error := demo(6, 0)
	fmt.Println(result, error)

	result, error = demo1(6, 0)
	fmt.Println(result, error)

	// Go语言中处理错误的方式
	result, _ = demo1(6, 2)
	fmt.Println(result)

	// 使用if处理错误，原则上每个错误都要处理
	if error != nil {
		fmt.Println("发生错误", error)
	} else {
		fmt.Println("程序执行成功", result)
	}
}

// 使用Go语言标准库创建错误，并返回
func demo(i, k int) (d int, e error) {
	if k == 0 {
		e = errors.New("初始值不能为0")
		d = 0
		return
	}
	d = i / k
	return
}

// 如果错误内容由多个组成，可以使用以下方式
func demo1(i, k int) (d int, err error) {
	if k == 0 {
		err = fmt.Errorf("%s%d和%d", "除数不能是0，两个参数分别为：", i, k)
		d = 0
		return
	}
	d = i / k
	return
}
