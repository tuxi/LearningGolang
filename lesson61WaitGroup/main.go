package main

import (
	"fmt"
	"sync"
)

/*
WaitGroup 介绍
- Golang中sync包中提供了基本同步单元，如互斥锁等，除了Once和WaitGroup类型，大部分都只适用于低水平程序线程，高水平同步线程使用channel通信会好一些
WaitGroup直译为等待组，其实就是计数器，只要计数器中有内容将一直阻塞
- 在Golang中的WaitGroup存在于sync包中，在sync包中类型都是不应该被拷贝的，源码定义如下：
```go
// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of
// goroutines to wait for. Then each of the goroutines
// runs and calls Done when finished. At the same time,
// Wait can be used to block until all goroutines have finished.
//
// A WaitGroup must not be copied after first use.
type WaitGroup struct {
	noCopy noCopy

	// 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
	// 64-bit atomic operations require 64-bit alignment, but 32-bit
	// compilers do not ensure it. So we allocate 12 bytes and then use
	// the aligned 8 bytes in them as state.
	state1 [12]byte
	sema   uint32
}
```
- Go语言标准库中WaitGroup只有三个方法
	- Add(delta int)表示向内部计数器添加增量(delta)，其中参数delta可以是负数
	- Done()表示减少WaitGroup计数器的值，应当在程序最后执行梦想但与Add(-1)
	- Wait()表示阻塞直到WaitGroup的计数器为0
```go
type WaitGroup
  func (wg *WaitGroup) Add(delta int)
  func (wg *WaitGroup) Done()
  func (wg *WaitGroup) Wait()
```
*/

var wg sync.WaitGroup

func main() {

	/*
		使用WaitGroup可以有效解决goroutine未执行完成主协程执行完成，导致程序结束，最终goroutine未执行问题
	*/

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go demo(i, func() {
			fmt.Println("当前任务结束: ", i)
			wg.Done()
		})
	}
	fmt.Println("开始阻塞")
	wg.Wait()
	fmt.Println("任务结束完成，解除阻塞")
}

func demo(index int, fn func()) {
	for i := 0; i < 5; i++ {
		fmt.Printf("第%d次执行,i的值为:%d\n", index, i)
	}
	fn()
}
