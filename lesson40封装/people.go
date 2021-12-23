package main

import "fmt"

/*
封装
* 封装主要体现在两个方面:封装数据、封装业务
* Go语言中通过首字母大小控制访问权限.属性首字母小写对外提供访问方法是封装数据最常见的实现方式
* 可以通过方法封装业务
  * 提出方法是封装
  * 控制结构体属性访问,对外提供访问方法也是封装
* 在面向对象中封装的好处:
  * 安全性.结构体属性访问受到限制,必须按照特定访问渠道
  * 可复用性,封装的方法实现可复用性
  * 可读写,多段增加代码可读性
*/

// Go语言同包任意位置可以访问全局内容,封装控制可以控制包外访问结构体中数据
type People struct {
	name string
	age  int
}

func (p *People) SetName(name string) {
	p.name = name
}

func (p *People) GetName() string {
	return p.name
}

func (p *People) SetAge(age int) {
	p.age = age
}

func (p *People) GetAge() int {
	return p.age
}

func main() {
	var p People
	p.SetAge(18)
	p.SetName("yuan")
	fmt.Println(p)
}
