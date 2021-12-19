package main

import (
	"fmt"
	"time"
)

func main() {

	// 声明时间
	// 仅仅声明未进行初始化是无意义的，0001-01-01 00:00:00 +0000 UTC
	var t time.Time
	fmt.Println(t)

	// 获取当前时间
	now := time.Now()
	fmt.Println(now)

	// 也可以通过时间戳创建时间类型变量(距离1970年1月1日的纳秒差)
	/*
		1秒(s)=1000毫秒(ms)
		1秒(s)=1000000微秒(μs)
		1秒(s)=1000000000纳秒(ns)
	*/
	// 根据时间戳创建时间
	t1 := time.Unix(0, t.UnixNano())
	fmt.Println(t.String())
	fmt.Println(t1)

	// 根据自己的要求创建时间
	// time.Local 取到本地时间位置对象，东八区
	t2 := time.Date(2022, 12, 20, 7, 8, 9, 0, time.Local)
	fmt.Println(t2) // 输出 2022-12-20 07:08:09 +0800 CST

	// 在time 包下提供了大量的函数或方法获取时间的某一项
	t = time.Now()
	fmt.Println(t)
	fmt.Println(t.Year())       //年
	fmt.Println(int(t.Month())) //月
	fmt.Println(t.Day())        //日
	fmt.Println(t.Date())       //三个参数,分别是:年,月,日
	fmt.Println(t.Hour())       //小时
	fmt.Println(t.Minute())     //分钟
	fmt.Println(t.Second())     //秒
	fmt.Println(t.Clock())      //三个参数,分别是:小时,分钟,秒
	fmt.Println(t.Nanosecond()) //纳秒
	fmt.Println(t.Unix())       //秒差
	fmt.Println(t.UnixNano())   //纳秒差

	// 时间和字符串转换

	// 时间转为string
	// 参数必须是这个时间,为Go语言出现时间
	s := t.Format("2006-01-02 15:04:05")
	fmt.Println(s)

	// string转换为Time
	s1 := "2022-02-04 22:02:04"
	t3, err := time.Parse("2022-02-04 22:02:04", s1)
	fmt.Println(t3, err)

}
