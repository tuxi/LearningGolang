package main

import "fmt"

/*
接口
- 接口解释：接口是一组行为规范的定义
- 接口中只能有方法的声明，方法只能有名称、参数、返回值，不能有方法体
- 每个接口可以有多个方法声明，结构体把接口中所有方法都重写后，结构体就属于接口类型
- Go语言中接口和结构体之间的关系是传统面向对象is-like-a的关系
- 接口可以继承接口，且Go语言推荐把接口中方法拆分为多个接口
- 定义接口类型关键字是interface
```go
type 接口名 interface{
  方法名(参数列表) 返回值列表
}
```
*/

/*
	接口中声明方法，结构体重写接口中方法后，编译器认为结构体实现了接口
	重写的方法要必须和接口中方法名称、方法参数（参数名称可以不同）、返回值列表完全相同
*/

type People struct {
	name string
	age  int
}

type Live interface {
	run(run int)
	Eat // 接口继承接口
}

/*
接口可以继承接口
*/
type Eat interface {
	eat()
}

func (p *People) run(run int) {
	fmt.Println(p.name, "正在跑步, 跑了", run, "米")
}

func (p *People) eat() {
	fmt.Println(p.name, "正在吃饭")
}

func main() {
	p := People{"杨校园", 18}
	p.run(1000)
}
