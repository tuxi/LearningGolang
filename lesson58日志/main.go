package main

/*
# 一.日志简介

* 使用开发工具时,控制台打印的信息就是日志信息
* 项目最终发布后是没有开发工具的,而需要记录日志应该把信息输出到文件中,这个功能也是日志的功能
* 在Go语言标准的log包提供了对日志的支持
* 有三种级别日志输出
  * Print() 输出日志信息
  * Panic()  打印日志信息,并触发panic,日志信息为Panic信息
  * Fatal()  打印日志信息后调用os.Exit(1)
* 所有日志信息打印时都带有时间,且颜色为红色
* 每种级别日志打印都提供了三个函数
  * Println()
  * Print()
  * Printf()
* 日志文件扩展名为log
# 二.普通日志信息打印
* 官方源码如下
```go
func Println(v ...interface{}) {
	std.Output(2, fmt.Sprintln(v...))
}
```
* 直接使用log包调用Println()即可
```go
log.Println("打印日志信息")
```
# 三.Panic日志信息打印
* 通过源码可以看出在日志信息打印后调用了panic()函数,且日志信息为panic信息
```go
// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	std.Output(2, s)
	panic(s)
}
```
* 执行后输出日志信息,同时也会触发panic
```go
log.Panicln("打印日志信息")
```
# 四.致命日志信息
* 打印日志后,终止程序
```go
// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}
```
* 执行日志打印后,程序被终止
```go
log.Fatal("打印日志信息")
```
*/

import (
	"fmt"
	"log"
	"os"
)

func main() {

	defer func() {
		recover()
	}()
	fileLog()
	panicLog()
	fatalLog()
}

func fileLog() {
	/*
		将日志写入到文件中
	*/
	f, err := os.OpenFile("/Users/xiaoyuan/Desktop/testDir11/golog.log", os.O_APPEND|os.O_CREATE, 0777)
	defer f.Close()
	if err != nil {
		fmt.Println("打开文件失败", err)
	} else {
		fmt.Println("打开文件成功")
		log.New(f, "[info]\t", log.Ltime)
		log.Println("输出日志信息")
	}
}

func panicLog() {

	/*
		panic日志信息打印
		通过源码可以看出日志信息打印后调用了panic()函数，且日志信息为panic信息
	*/
	// 执行输出日志信息，同时也会触发panic
	log.Panicln("打印日志信息")
}

func fatalLog() {
	// 执行日志打印后，程序被中断
	log.Fatal("打印日志信息")
}
