package main

import (
	"fmt"
)

/*
多态
多态：同一件事由于条件不同产生的结果不同
- 由于Go语言中结构体不能相互转换，所以没有结构体（父子结构体）的多态，只有基于接口的多态，这也符合Go语言对面向对象的诠释
- 多态在代码层面最常见的一种方式是接口当作方法的参数
*/

/*
- 结构体实现了接口的全部方法，就认为结构体属于接口类型，这是可以把结构体变量赋值给接口变量
- 重写接口时接受者为`Type`和`*Type`的区别
	- `*Type`可以调用`*Type`和`Type`作为接受者的方法，所以只要接口中多个方法中至少出现一个`*Type`作为接受者进行重写的方法，就必须把结构体指针赋值给接口变量，否则编译报错
	- `Type`只能调用`Type`作为接受者的方法
*/

type Live interface {
	run()
	eat()
}

type People struct {
	name string
}

func (p *People) run() {
	fmt.Println(p.name, "正在跑步")
}

func (p People) eat() {
	fmt.Println(p.name, "在吃饭")
}

/*
既然接口可以接收实现接收所有方法的结构体变量，接口也就可以作为方法（函数）参数
*/
type Live1 interface {
	run1()
}
type People1 struct{}
type Animate1 struct{}

func (p *People1) run1() {
	fmt.Println("人在跑")
}
func (a *Animate1) run1() {
	fmt.Println("动物在跑")
}

func sport(live Live1) {
	live.run1()
}

func main() {
	// 重写接口
	var run Live = &People{"yuan"}
	run.run()
	run.eat()

	// 接口作为方法的参数
	p := &People1{}
	p.run1() // 输出：人在跑
	a := &Animate1{}
	a.run1() // 输出：动物在跑

}
