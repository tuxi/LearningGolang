package main

import "fmt"

/*
结构体指针
由于结构体是值类型的，在方法传递时如果希望传递结构体地址，使用结构体指针即可
可以结合new(T)函数创建结构体指针
*/

type People struct {
	Name string
	Age  int
}

func main() {

	// 1.使用New函数创建结构体指针
	p := new(People)
	// 由于结构体本质是值类型，所以创建时就开辟了内容空间
	fmt.Println(p == nil) // false
	fmt.Println(p)        // &{0}
	// 由于结构体中的属性不是指针类型，所以可以直接调用
	p.Name = "yangxiaoyuan"
	p.Age = 18
	fmt.Println(p) // &{yangxiaoyuan 18}

	// 2.声明指针变量初始化结构体指针
	var p1 *People
	// 给结构体指针赋值
	p1 = &People{"yuan", 19}
	fmt.Println(p1)
	// 还可以通过短变量的方式去写
	p2 := &People{Name: "yaya"}
	fmt.Println(p2)

	/*
				判断结构体指针是否相等
		结构体指针比较的是地址
		*结构体指针，取出的是结构体指针指向内存的值
	*/

	// 比较两个结构体是否相等，内容相同就相同
	p3 := People{"yuan", 18}
	p4 := People{"yuan", 18}
	fmt.Printf("%p %p\n", &p3, &p4)
	fmt.Println(p3 == p4) // true

	// 比较两个结构体指针是否相等
	p6 := &People{"yuan", 18}
	// 结构体变量不能和指针比较，可以使用*指针取出地址中的值，进行比较
	fmt.Println(p3 == *p6) // true

	p7 := new(People)
	p8 := &People{"yuan", 18}
	fmt.Println(p7 == p8) // false
}
