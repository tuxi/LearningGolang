package main

import "fmt"

/*
继承
* 按照传统面向对象思想,继承就是把同一类事物提出共同点为父类,让子类可以复用父类的可访问性内容.
* 继承有多种实现方式
  * 通过关键字继承,强耦合实现方式
  * 组合式继承,松耦合继承方式
* 使用过Java或C#的应该知道尽量少用继承而是使用组合代替继承,可以使用高内聚,低耦合.Java之父之前在一次采访的时候也说过,如果给他一次机会重新做Java,他最希望修改的地方就是继承
* Go语言中的继承是通过组合实现
*/

// 在Go语言中支持匿名属性（结构体中的属性名字），但是每个最多只能存在匿名属性，编译器认为类型就是属性名，我们在使用时就把类型当作属性进行使用
type People struct {
	string
	int
}

/*
* 传统面向对象中类与类之间的关系
  * 继承:is-a,强耦合性,一般认为类与类之间具有强关系
  * 实现:like-a,接口和实现类之间的关系
  * 依赖:use-a,具有偶然性的、临时性的、非常弱的，但是B类的变化会影响到A,一般作为方法参数
  * 关联:has-a一种强依赖关系，比如我和我的朋友；这种关系比依赖更强、不存在依赖关系的偶然性、关系也不是临时性的，一般是长期性的，而且双方的关系一般是平等的、关联可以是单向、双向的
  * 聚合:has-a,整体与部分、拥有的关系
  * 组合:contains-a,他体现的是一种contains-a的关系，这种关系比聚合更强，也称为强聚合；他同样体现整体与部分间的关系，但此时整体与部分是不可分的，整体的生命周期结束也就意味着部分的生命周期结束
  * 组合>聚合>关联>依赖
* Go语言中标准的组合关系
*/
type Teacher struct {
	peo       People
	classroom string
}

/*
使用匿名属性完成Go语言中的继承
- Go语言中继承很好实现，把另一个结构体类型当作另一个结构体属性，可以直接调用另一个结构体中的内容
- 因为Go语言中结构体不能互相转换，所以不能把字结构体变量赋值给父结构体变量
*/
type Student struct {
	People    // 使用匿名属性 让Student 继承自 People
	classroom string
}

func main() {
	p := People{"yuan", 18}
	fmt.Println(p.string, p.int)

	teacher := Teacher{People{"yuan", 18}, "三年二班"}
	fmt.Println(teacher.peo.string, teacher.peo.string, teacher.classroom)

	student := Student{People: People{"yuan", 18}, classroom: "三年二班"}
	fmt.Println(student.classroom, student.People.string, student.People.int)

}
