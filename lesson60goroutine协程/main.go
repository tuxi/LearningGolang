package main

import (
	"fmt"
	"time"
)

/*
goroutine简介
- Golang中最迷人的一个优点就是从语言层面就支持并发
- 在Golang中goroutine(协程)类似于其他语言的线程
- 并发和并行
	- 并发(parallelism)指不同的代码片段同时在不同的物理处理器上支持
	- 并行(concurrency)指同时管理多个事情，物理处理器上可能运行某个内容一半后就去处理其他事情
	- 在一般看来并发的性能要好于并行，因为计算机的物理资源是固定的，较少的，而程序需要执行的内容时很多的，所以并发是`以较少的资源去做更多的事情`
- 几种主流并发模型
	- 多线程，每个线程只处理一个请求，只有请求结束后，对应的线程才能接收下一个请求，这种模式在高并发下，性能开销极大
	- 基于回调的异步IO，在程序运行过程中可能产生大量回调导致导致维护成本加大, 程序执行流程也不便于思维
	- 协程，不需要抢占调用，可以有效提升线程任务的并发性，弥补了多线程模式的缺点，Golang在语言层面就支持，而其他语言就很少支持
- Goroutine的语法
	- 表达式可以是一条语句
	- 表达式也可以是函数，函数返回值即使有，也无效，当函数执行完此goroutine自动结束
```go
	go 表达式
```
*/

func main() {

	/*
		对比多次调用函数和使用goroutine的效果
	*/

	/*
		添加go关键字后发现控制台什么也没有输出
		原因：把demo()设置到协程后没等到函数执行，主线程执行结束
	*/
	for i := 0; i < 3; i++ {
		go demo(i)
	}

	/*
		添加休眠等待goroutine执行结束
		- 这种方式很大的问题就是休眠时间，如果休眠时间设置过小，可能goroutine并没有执行完成，如果休眠时间设置过大，影响程序执行，找到本次执行的休眠时间，下次程序时这个休眠时间可能可能过大或者过小

		通过程序运行结构发现每次执行结果都不一定是一样，因为每个demo()都是并发执行
	*/

	time.Sleep(3e9)
	fmt.Println("程序执行结束")

}

func demo(index int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("第%d次执行，i的值为%d\n", index, i)
	}
}
