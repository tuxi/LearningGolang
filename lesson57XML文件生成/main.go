package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

/*
XML文件生成
# 一.生成XML

* 生成XML只要在学习下encoding/xml包下的Marshal()函数,结合输入流就可以完成xml文件生成
* 在encoding/xml中有常量,常量中是xml文档头
```go
const (
	// Header is a generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)
```

*/

type People struct {
	XMLNAme xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}

func main() {

	peo := People{Id: 111, Name: "yuan", Address: "海淀"}
	b, err := xml.MarshalIndent(peo, "", "	")
	if err != nil {
		fmt.Println("生成xml失败", err)
	} else {
		// 给xml拼接头部信息
		b = append([]byte(xml.Header), b...)
		err = ioutil.WriteFile("/Users/xiaoyuan/Desktop/testDir11/people.xml", b, 0666)
		if err != nil {
			fmt.Println("xml文件写入失败", err)
		} else {
			fmt.Println("xml文件写入成功")
		}
	}

}
