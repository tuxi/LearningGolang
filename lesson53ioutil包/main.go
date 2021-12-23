package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
# 一. ioutil包

* ioutil包下提供了对文件读写的工具函数,通过这些函数快速实现文件的读写操作
* ioutil包下提供的函数比较少,但是都是很方便使用的函数
```
func NopCloser(r io.Reader) io.ReadCloser
func ReadAll(r io.Reader) ([]byte, error)
func ReadFile(filename string) ([]byte, error)
func WriteFile(filename string, data []byte, perm os.FileMode) error
func ReadDir(dirname string) ([]os.FileInfo, error)
func TempDir(dir, prefix string) (name string, err error)
func TempFile(dir, prefix string) (f *os.File, err error)
```
*/

func main() {

	/*
		读取文件
	*/
	// 打开完文件后，可以使用ReadAll() 把文件中所有内容都读取到
	f, err := os.Open("/Users/xiaoyuan/Desktop/testDir11/go.txt")
	defer f.Close()
	if err != nil {
		fmt.Println("打开文件失败", err)
	} else {
		// 使用ReadAll 读取文件为字节切片
		b, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println("读取文件失败")
		} else {
			fmt.Println("读取文件成功m，内容:", string(b))
		}
	}

	// 也可以直接读取文件中的内容
	b, err := ioutil.ReadFile("/Users/xiaoyuan/Desktop/testDir11/go.txt")
	if err != nil {
		fmt.Println("读取文件失败", err)
	} else {
		fmt.Println("读取文件成功")
		fmt.Println("文件内容：", string(b))
	}

	/*
		写入文件
	*/
	// 写文件也很简单，直接使用writeFile函数即可，但是源码中已经规定此文件只能是是刻度状态，切不是尾加数据
	err = ioutil.WriteFile("/Users/xiaoyuan/Desktop/testDir11/abc.txt", []byte("123455666"), 0666)
	if err != nil {
		fmt.Println("写入文件失败", err)
	} else {
		fmt.Println("数据写入成功")
	}

	// ReadDir快速获取某个文件夹中所有文件信息的函数
	fs, err := ioutil.ReadDir("/Users/xiaoyuan/Desktop/testDir11")
	if err != nil {
		fmt.Println("读取文件夹失败", err)
	} else {
		// 遍历文件目录
		for _, filInfo := range fs {
			fmt.Println(filInfo.Name())
		}
	}

}
