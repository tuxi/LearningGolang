package main

import "fmt"

/*
方法
方法和函数比较像，区别是函数属于包，而方法属于结构体，通过结构体变量调用
默认是函数，隶属于包，所以需要添加标识，告诉编译器这个方法属于哪个结构体
调用方法时就把调用者赋值给接收者(下面的变量名就是接受者)
```go
func (变量名 结构体类型) 方法名(参数列表) 返回值列表{
  //方法体
}
```

* Go语言中已经有函数了,又添加了对方法的支持主要是保证Go语言是面向对象的.Go语言官方对面向对象的解释

  * 翻译如下:虽然面向对象没有统一的定义，但是对于我们来说对象仅仅是一个有着方法的值或变量,而方法就是一个属于特定类型的函数

* 从上面的解释可以看出,官方给出可明确说明,方法类似于函数.方法归属于特定类型
*/

// 定义结构体
type People struct {
	Name   string  // 姓名
	Weight float64 // 体重
}

func (p People) run() {
	fmt.Println(p.Name, " 正在跑步")
}

/*
如果设定需求每次跑步后，体重-0.1斤
因为结构体是值类型的，修改方法中结构体变量的值，主函数p的值不会改变，因为传递的是副本
所以修改方法结构体类型改为结构体指针类型即可
*/
func (p *People) run1() {
	p.Weight -= 10
	fmt.Println(p.Name, " 正在跑步", "体重: ", p.Weight)
}

func main() {
	p := People{"yangxiaoyuan", 120}
	p.run()

	p.run1()
	p.run1()
	p.run1()

}
