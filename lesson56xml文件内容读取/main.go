package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

/*
XML文件内容读取
# 一.Go语言标准库提供的API

* 在encoding/xml包下提供了对XML序列化和反序列化的API
* 使用Unmarshal可以直接把XML字节切片数据转换为结构体
* 转换时按照特定的转换规则进行转换,且数据类型可以自动转换
```
* 如果结构体字段的类型为字符串或者[]byte，且标签为",innerxml"，
  Unmarshal函数直接将对应原始XML文本写入该字段，其余规则仍适用。
* 如果结构体字段类型为xml.Name且名为XMLName，Unmarshal会将元素名写入该字段
* 如果字段XMLName的标签的格式为"name"或"namespace-URL name"，
  XML元素必须有给定的名字（以及可选的名字空间），否则Unmarshal会返回错误。
* 如果XML元素的属性的名字匹配某个标签",attr"为字段的字段名，或者匹配某个标签为"name,attr"
  的字段的标签名，Unmarshal会将该属性的值写入该字段。
* 如果XML元素包含字符数据，该数据会存入结构体中第一个具有标签",chardata"的字段中，
  该字段可以是字符串类型或者[]byte类型。如果没有这样的字段，字符数据会丢弃。
* 如果XML元素包含注释，该数据会存入结构体中第一个具有标签",comment"的字段中，
  该字段可以是字符串类型或者[]byte类型。如果没有这样的字段，字符数据会丢弃。
* 如果XML元素包含一个子元素，其名称匹配格式为"a"或"a>b>c"的标签的前缀，反序列化会深入
  XML结构中寻找具有指定名称的元素，并将最后端的元素映射到该标签所在的结构体字段。
  以">"开始的标签等价于以字段名开始并紧跟着">" 的标签。
* 如果XML元素包含一个子元素，其名称匹配某个结构体类型字段的XMLName字段的标签名，
  且该结构体字段本身没有显式指定标签名，Unmarshal会将该元素映射到该字段。
* 如果XML元素的包含一个子元素，其名称匹配够格结构体字段的字段名，且该字段没有任何模式选项
  （",attr"、",chardata"等），Unmarshal会将该元素映射到该字段。
* 如果XML元素包含的某个子元素不匹配以上任一条，而存在某个字段其标签为",any"，
  Unmarshal会将该元素映射到该字段。
* 匿名字段被处理为其字段好像位于外层结构体中一样。
* 标签为"-"的结构体字段永不会被反序列化填写。
```
*/

/*
根据`student.xml`中的结构创建结构体，用于装载XML数据
	- 结构体中属性首字母必须大写，否则无法装配
*/
type People struct {
	XNLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}

type Peoples struct {
	XMLName xml.Name `xml:"peoples"`
	Version string   `xml:"version,attr"`
	Peos    []People `xml:"people"`
}

func main() {

	/*
		单层XML读取
	*/
	// 读取student.xml中的内容
	s := new(People)
	b, err := ioutil.ReadFile("people1.xml")
	if err != nil {
		fmt.Println("读取people.xml内容失败", err)
	} else {
		fmt.Println(string(b))
		// 装载
		err = xml.Unmarshal(b, s)
		if err != nil {
			fmt.Println("装载xml数据失败", err)
		} else {
			fmt.Println("装载xml数据成功")
			fmt.Println(s) // &{{ } 888 杨孝远 北京市房山}
		}
	}

	/*
		多层嵌套XML文件读取
	*/
	peo := new(Peoples)
	b, err = ioutil.ReadFile("people.xml")
	if err != nil {
		fmt.Println("读取people.xml内容失败", err)
	} else {
		fmt.Println(string(b))
		err = xml.Unmarshal(b, peo)
		if err != nil {
			fmt.Println("装载xml失败", err)
		} else {
			fmt.Println("装载xml成功")
			fmt.Println(peo)
			/*
							装载xml成功
				&{{ peoples} 0.9 [{{ } 888 yangxiaoyuan 北京市房山} {{ } 999 杨校园 北京市海淀区}]}
			*/
		}
	}

}
