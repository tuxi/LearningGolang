package main

import (
	"fmt"
	"time"
)

/*
DeadLock 死锁
- 在主goroutine中向无环村channel添加内容或在主goroutine中向channel添加内容的个数已经大于channel缓存个数就会产生死锁
	死锁报错log
	```
	fatal error: all goroutines are asleep - deadlock!
	```
- 死锁：在程序中多个进程(golang中goroutine)由于相互竞争资源而产生的阻塞（等待）状态，而这种状态一致保持下去，此时称这个线程是死锁状态
- 主协程中ch<-1，无缓存channel无论添加还是取出数据都会阻塞goroutine，当前程序无其他代码，主goroutine会一直被阻塞下去，此时goroutine就是死锁状态
*/

func main() {
	// demo1()
	// demo2()
	demo4()
}

/*
无缓存通道
*/
func demo1() {
	// 创建无缓存通道的channel
	ch := make(chan int)
	// 写入数据
	ch <- 1
	// 无读取，会造成死锁
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println("demo1执行结束")
}

func demo2() {
	ch := make(chan int)
	go func() {
		ch <- 1
		fmt.Println("执行goroutine") // 无输出
	}()
	time.Sleep(5e9)
	fmt.Println("demo2执行结束") // demo2执行结束

	/*
				以上输出·emo2执行结束·，而未输出‘执行goroutine’，说明goroutine中的代码一直在`ch <- 1`时处于等待中，而主通道中的代码在sleep结束后正常执行，而为造成死锁
				结果：
				- 而下面代码就不会产生死锁
		  		- 通过代码示例可以看出,在使用无缓存channel时,特别要注意的是在主协程中有操作channel代码
	*/
}

/*
	有缓存通道
*/
func demo3() {

	// 创建一个有缓存通道的channel，缓存大小3，里面消息个数小于等于3时，都不会阻塞goroutine
	ch := make(chan int, 3)
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1 // 此行出现死锁，超出缓存大小数量
	/*
		以上创建了3个缓存通道，但是写入了4次，会造成死锁
		fatal error: all goroutines are asleep - deadlock!
	*/
	fmt.Println("demo3执行结束")
}

/*
在goroutine中有缓存channel的缓存大小是不能改变的，但是只要不超过缓存数量大小，都不会出现阻塞状态
*/
func demo4() {
	// 创建一个有缓存通道的channel，缓存大小3，里面消息个数小于等于3时，都不会阻塞goroutine
	ch := make(chan int, 3)
	ch <- 1
	fmt.Println(<-ch) // 输出1
	ch <- 2
	fmt.Println(<-ch) // 输出2
	ch <- 3
	ch <- 4
	// 输出channel中消息个数
	fmt.Println(len(ch)) // 输出 2
	// 输出缓存大小总量
	fmt.Println(cap(ch)) // 输出3
	fmt.Println("demo4执行结束")
}
