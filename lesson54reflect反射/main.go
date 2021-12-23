package main

import (
	"fmt"
	"reflect"
)

/*
反射介绍
- 在Go语言标准库中reflect包提供了运行时反射，程序运行过程中动态操作结构体
- 当变量存储结构体属性名称时，想要对结构体这个属性赋值或查看时，就可以使用反射
- 反射还可以用作判断变量类型
- 整个reflect包中最重要的两个类型
	- reflect.type 类型
	- reflect.Value 值
- 获取到Type和Value的函数
	- reflect.TypeOf(interface{}) 返回Type
	- reflect.ValueOf(interface{}) 返回值Value
*/

type People struct {
	Id   int
	Name string
}

/*
	 结构体支持标记（tag）,标记通常都是通过反射技术获取到，结构体标记语法
	 type 结构体名称 struct {
		 属性名 类型 `key:"value"` // 注意key和value之间不要有空格，否则获取不到value
	 }
*/
type People1 struct {
	Name    string `xml:"name"`
	Address string `xml:"address"`
}

func main() {

	// 判断变量类型
	a := 1.6
	fmt.Println(reflect.TypeOf(a)) // 输出 float64

	// 获取结构体属性的值
	p := People{1, "yuan"}
	// 获取p的值
	v := reflect.ValueOf(p)
	// 获取属性个数，如果v不是结构体类型panic
	fmt.Println(v.NumField()) // 输出 2

	// 获取第0个属性id，并转换为Int64类型
	fmt.Println(v.Field(0).Int()) // 输出 1
	// 获取第一个属性name，并转换为string类型
	fmt.Println(v.Field(1).String()) // 输出 yuan

	// 根据名字获取类型，并把类型名称转换为string类型
	idvalue := v.FieldByName("Id")
	fmt.Println(idvalue.Kind().String()) // 输出int

	// 设置结构体属性的值时，要传递结构体指针，否则无法获取设置的结构体对象
	// 反射结构体属性时，要求属性名首字母必须大写，否则无法设置
	p1 := People{2, "杨校园"}
	// 反射时获取p的地址
	// Elem() 获取指针指向地址的封装
	// 地址的值必须调用Element() 才可以继续操作
	v = reflect.ValueOf(&p1).Elem()
	fmt.Println(v.FieldByName("Id").CanSet()) // 输出 true
	v.FieldByName("Name").SetString("小鸭")
	fmt.Println(p1) //  {2 小鸭}

	// 获取结构体标记（tag）
	t := reflect.TypeOf(People1{})
	name, _ := t.FieldByName("Name")
	// 获取完整标记
	fmt.Println(name.Tag)
	// 获取标记中xml对应的内容
	fmt.Println(name.Tag.Get("xml"))
}
