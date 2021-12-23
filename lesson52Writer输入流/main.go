package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

/*
# 一. 输入流

* 输入流就是把程序中数据写出到外部资源
* Go语言标准库中输出流是Writer接口
```go
// Writer is the interface that wraps the basic Write method.
//
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

* 注意:输入流时不要使用`os.Open()`因为这种方式获取的文件是只读的
*/

func main() {

	// 注意:输入流时不要使用`os.Open()`因为这种方式获取的文件是只读的
	fp := "/Users/xiaoyuan/Desktop/testDir11/go.txt"
	f, err := os.OpenFile(fp, os.O_APPEND, 0660)
	defer f.Close()
	if err != nil {
		fmt.Println("文件不存在， 创建文件")
		f, _ = os.Create(fp)
	}
	/*
		内容中识别特殊字符
		\r\n 换行
		\t 缩进
	*/

	// 使用文件对象重写的Writer接口，参数[]byte
	f.Write([]byte("使用Write接口写入数据\n"))

	// 使用stringWriter接口的方法，参数是字符串，使用更方便
	f.WriteString("写了\t一段\r\r内容123")
	fmt.Println("程序执行结束")

}
