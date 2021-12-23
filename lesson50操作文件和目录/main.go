package main

import (
	"fmt"
	"os"
)

/*
os包内容介绍
* 使用os包中内容进行操作系统文件或目录
* File结构体表示操作系统文件(夹)
```go
// File represents an open file descriptor.
type File struct {
	*file // os specific
}
```
```go
// file is the real representation of *File.
// The extra level of indirection ensures that no clients of os
// can overwrite this data, which could cause the finalizer
// to close the wrong file descriptor.
type file struct {
	pfd     poll.FD
	name    string
	dirinfo *dirInfo // nil unless directory being read
}
```
* 操作系统的文件都是有权限控制的,包含可读,可写等,在os包中FileMode表示文件权限,本质是uint32,可取值都以常量形式提供
```go
// A FileMode represents a file's mode and permission bits.
// The bits have the same definition on all systems, so that
// information about files can be moved from one system
// to another portably. Not all bits apply to all systems.
// The only required bit is ModeDir for directories.
type FileMode uint32
```
```go
// The defined file mode bits are the most significant bits of the FileMode.
// The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
// The values of these bits should be considered part of the public API and
// may be used in wire protocols or disk representations: they must not be
// changed, although new bits might be added.
const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
	ModeAppend                                     // a: append-only
	ModeExclusive                                  // l: exclusive use
	ModeTemporary                                  // T: temporary file; Plan 9 only
	ModeSymlink                                    // L: symbolic link
	ModeDevice                                     // D: device file
	ModeNamedPipe                                  // p: named pipe (FIFO)
	ModeSocket                                     // S: Unix domain socket
	ModeSetuid                                     // u: setuid
	ModeSetgid                                     // g: setgid
	ModeCharDevice                          // c: Unix character device, when ModeDevice is set
	ModeSticky                                     // t: sticky

	// Mask for the type bits. For regular files, none will be set.
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice

	ModePerm FileMode = 0777 // Unix permission bits
)
```
* FIleInfo是一个interface表示文件的信息
```go
// A FileInfo describes a file and is returned by Stat and Lstat.
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
}
```
# 二. 资源路径

* 在获取系统资源时资源路径分为相对路径和绝对路径
* 相对路径:在Go语言中相对路径用于是GOPATH,也就是项目的根目录
* 绝对路径:磁盘根目录开始表示资源详细路径的描述
*/

func main() {
	/*
		Go语言标准库中提供了两种创建文件的方式
	*/

	// 1.创建文件夹
	// 如果文件夹已经存在，则报错file exists
	error := os.Mkdir("/Users/xiaoyuan/Desktop/testDir11", os.ModeDir)
	if error != nil {
		fmt.Println("文件夹创建失败：", error)
	} else {
		fmt.Println("文件夹创建成功")
	}

	// 2.创建文件夹
	// 如果文件夹已经存在，不报错，保留原文件夹
	// 如果父目录不存在帮助创建
	error = os.MkdirAll("/Users/xiaoyuan/Desktop/testDir11", os.ModeDir)
	if error != nil {
		fmt.Println("文件夹创建失败:", error)
	} else {
		fmt.Println("文件夹创建成功")
	}

	// 3.创建空文件
	// 创建文件时要求文件目录必须已经存在
	// 如果文件已经存在则会创建一个空文件覆盖之前的文件
	file, err := os.Create("/Users/xiaoyuan/Desktop/testDir11/test.txt")
	if err != nil {
		fmt.Println("文件创建失败:", err)
	} else {
		fmt.Println("文件创建成功", file.Name())
	}

	// 重命名文件夹
	err = os.Rename("/Users/xiaoyuan/Desktop/testDir11", "/Users/xiaoyuan/Desktop/testDir22")
	if err != nil {
		fmt.Println("重命名文件夹失败", err)
	} else {
		fmt.Println("重命名文件夹成功")
	}

	// 重命名文件
	err = os.Rename("/Users/xiaoyuan/Desktop/testDir22/test.txt", "/Users/xiaoyuan/Desktop/testDir22/test1.txt")
	if err != nil {
		fmt.Println("重命名文件失败", err)
	} else {
		fmt.Println("重命名文件成功")
	}

	// 获取文件（夹）信息
	file, err = os.Open("/Users/xiaoyuan/Desktop/testDir22/test1.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败", err)
	} else {
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Println("获取文件信息失败", err)
		} else {
			fmt.Println(fileInfo.Name())    // 文件名称
			fmt.Println(fileInfo.IsDir())   // 是否是文件夹，返回bool，true表示文件夹，false表示文件
			fmt.Println(fileInfo.Mode())    // 文件权限
			fmt.Println(fileInfo.ModTime()) // 修改时间
			fmt.Println(fileInfo.Size())    // 文件大小
		}
	}

	// 删除文件或文件夹
	// 删除的内容只能是一个文件或空文件夹，且必须存在的路径

	// Remove删除一个不存在的文件夹
	err = os.Remove("/Users/xiaoyuan/Desktop/testDir22/aaa")
	if err != nil {
		fmt.Println("Remove文件删除失败", err)
	} else {
		fmt.Println("Remove删除文件成功")
	}

	//
	// 只要文件夹存在，删除文件夹
	// 无论文件夹内是否有内容都会删除
	// 无论文件夹是否有内容都会删除
	// 如果删除目标是文件，则删除文件
	err = os.RemoveAll("/Users/xiaoyuan/Desktop/testDir22")
	if err != nil {
		fmt.Println("RemoveAll删除失败", err)
	} else {
		fmt.Println("RemoveAll删除成功")
	}
}
