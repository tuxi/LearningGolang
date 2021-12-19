package main

import "fmt"

func main() {
	a := 1
	fmt.Println(&a)
	// 对变量赋值的时候，是修改内存空间的内容，不会修改指针指向
	a = 10
	// 打印指针地址
	fmt.Println(&a)
	// 打印指针指向的内存内容
	fmt.Println(a)

	// 把a地址中的值拿出来放到b的内存间中
	b := a
	// 修改b的内存中值，不会影响a的内存空间
	b = 3
	fmt.Println(&b, &a)
	fmt.Println(a, b)

	c := &a
	// 打印c的类型为*int类型的指针
	fmt.Printf("%T\n", c)

	// 声明一个指针类型a，指针指向的是一块内存空间的地址
	// 声明指针是不会开辟内存空间的，只是准备要指向内存某个空间，而声明变量会开辟内存空间，准备存放内存，所以指针变量类型就是把一个变量的地址赋值给指针变量

	// 声明一个空指针，其未指向任何地址
	var d *int
	fmt.Println(d == nil)
	f := 123
	// 给指针变量赋值f的内存地址
	d = &f

	// 通过*d 超过7⃣指针变量指向的内容空间
	fmt.Println(*d, f)
	f = 100

	fmt.Println(*d, f)
}
