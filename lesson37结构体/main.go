package main

import "fmt"

/*
结构体：就是将一个变量或多个变量组合到一起，形成新的类型，这就是结构体struct
Go语言中的结构体和c++的结构体有些类似，而Java和Objective-c中的类本质上就是结构体
结构体是值类型的
结构体定义的语法
  通过语法可以看出,Go语言发明者明确认为结构体就是一种自定义类型
```
    type 结构体名称 struct{
      名称 类型//成员或属性
    }
```
*/

/*
	定义结构体
	结构体可以在函数内部也可以在函数外部（和普通变量一样），定义的位置影响到结构体的访问范围
	结构体如果定义在外部，结构体的首字母大小写将影响到结构体是否可以跨包访问
	如果结构体首字母大写，那么结构体属性成员的首字母大小写将影响到
*/

type People struct {
	Name string
	Age  int
}

func main() {

	// 声明结构体，
	// 由于结构体是值类型，所以声明就会开辟内存空间
	// 所有成员为类型对应的初始值
	var p People
	fmt.Println(p)         // { 0}
	fmt.Printf("%p\n", &p) // 0xc0000a4018

	// 初始化结构体
	p1 := People{"yuan", 18}
	fmt.Println(p1) // {yuan 18}

	// 初始化时也可以明确指定成员名称
	p2 := People{Age: 18, Name: "yangxiaoyuan"}
	fmt.Println(p2)

	p3 := People{Name: "sdkasdk"}
	fmt.Println(p3)

	// 初始化还可以通过.语法初始化
	var p4 People
	p4.Age = 18
	p4.Name = "yangxiaoyuan"
	fmt.Println(p4)

	// ==号判断两个结构体的内容是否相同
	fmt.Println(p4 == p2)            // true
	fmt.Printf("%p %p \n", &p4, &p2) // 0xc00000c0c0 0xc00000c078
}
