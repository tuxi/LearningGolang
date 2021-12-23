package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
# 二.RWMutex读写锁

* RWMutex 源码如下
```go
// There is a modified copy of this file in runtime/rwmutex.go.
// If you make any changes here, see if you should make them there.

// A RWMutex is a reader/writer mutual exclusion lock.
// The lock can be held by an arbitrary number of readers or a single writer.
// The zero value for a RWMutex is an unlocked mutex.
//
// A RWMutex must not be copied after first use.
//
// If a goroutine holds a RWMutex for reading and another goroutine might
// call Lock, no goroutine should expect to be able to acquire a read lock
// until the initial read lock is released. In particular, this prohibits
// recursive read locking. This is to ensure that the lock eventually becomes
// available; a blocked Lock call excludes new readers from acquiring the
// lock.
type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // semaphore for writers to wait for completing readers
	readerSem   uint32 // semaphore for readers to wait for completing writers
	readerCount int32  // number of pending readers
	readerWait  int32  // number of departing readers
}
```
* Go语言标准库中API如下
```go
type RWMutex
  func (rw *RWMutex) Lock()//禁止其他协程读写
  func (rw *RWMutex) Unlock()
  func (rw *RWMutex) RLock()//禁止其他协程写入,只能读取
  func (rw *RWMutex) RUnlock()
  func (rw *RWMutex) RLocker() Locker
```
* Go语言中的map不是线程安全的,多个goroutine同时操作会出现错误.
* RWMutex可以添加多个读锁或一个写锁.读写锁不能同时存在.
  * map在并发下读写就需要结合读写锁完成
  * 互斥锁表示锁的代码同一时间只能有一个人goroutine运行,而读写锁表示在锁范围内数据的读写操作
```go
*/

func main() {

	var rwm sync.RWMutex
	m := make(map[string]string)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(j int) {
			// 没有锁在map时可能出现问题
			rwm.Lock()
			m["key"+strconv.Itoa(j)] = "value" + strconv.Itoa(j)
			fmt.Println(m)
			rwm.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("程序结束")
}
