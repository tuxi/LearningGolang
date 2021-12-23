package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
互斥锁
- Go语言中多个协程操作一个变量时会出现冲突问题
- `go run -race main.go` 可以查看竞争
- 可以使用`sync.Mutex`对内容加锁
- 互斥锁的使用场景
	- 多个goroutine访问同一个函数(代码段)
	- 这个函数操作一个全局变量
	- 为了保证共享变量的安全性，值合法性
*/

/*
使用互斥锁模拟售票窗口
*/

var (
	// 票数
	num = 100
	wg  sync.WaitGroup
	// 互斥锁
	mu sync.Mutex
)

func sellTricker(i int) {
	defer wg.Done()
	for {
		// 加锁，多个goroutine互斥
		mu.Lock()
		if num >= 1 {
			fmt.Println("第", i, "个窗口卖了", num)
			num -= 1
		}
		// 解锁
		mu.Unlock()
		if num <= 0 {
			break
		}
		// 添加休眠，防止结果可能出现在一个goroutine中
		time.Sleep(time.Duration(rand.Int63n(1000) * 1e6))
	}
}

// 买票的例子
func demo1() {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 4个柜台一起卖票
	// 计算起始值和票数相同
	wg.Add(4)
	go sellTricker(1)
	go sellTricker(2)
	go sellTricker(3)
	go sellTricker(4)
	wg.Wait()
	fmt.Println("所有票卖完")

	/*
		4个柜台最终按顺序卖完所有的票
	*/
}

func main() {

	demo1()
}
