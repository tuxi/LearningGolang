package main

import (
	"fmt"
	"runtime"
)

/*
select
- Golang中select和switch结构特别像，但是select中case的条件只能是I/O
- select 的语法(condition是条件)
```go
select{
  case condition:
  case condition:
  default:
}
```

- select执行过程
	* 每个case 必须是一个IO操作
	* 哪个case可以执行就执行哪个
	* 多个case都可以执行，随机执行一个
	* 所有case都不能执行时，执行default
	* 所有case都不能执行，且没有default，将会阻塞
*/

func main() {
	// demo1()
	demo2()
}

func demo1() {
	runtime.GOMAXPROCS(1)
	ch1 := make(chan int, 1)
	ch2 := make(chan string, 1)
	ch1 <- 1
	ch2 <- "Hello world"
	select {
	case value := <-ch1:
		fmt.Println(value)
	case value := <-ch2:
		fmt.Println(value)
	}
	fmt.Println("demo1执行结束")
	/*
		以上可能输出1，也可能输出Hello world，每次输出的结果都不一定相同
	*/
}

/*
select和for循环结合使用，
下面例子演示了一直在接收消息的例子
*/
func demo2() {

	// 创建一个无缓存的通道
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}

	// 如果一直接收消息，应该是死循环for{} ，下面例子中明确知道消息个数
	for i := 0; i < 5; i++ {
		select {
		case value := <-ch:
			fmt.Println("接收到消息:", value)
		default:
			fmt.Println("没有消息，进入了default")
		}
	}
	fmt.Println("demo2执行结束")

	/*
		break可以对select生效，如果for中潜逃select，break选择最近结构
	*/
}
