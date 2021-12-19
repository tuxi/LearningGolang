package main

import (
	"fmt"
	"sort"
)

/*
* Go语言标准库中sort提供了排序API
* sort包提供了多种排序算法,这些算法是内部实现的,每次使用sort包排序时,会自动选择最优算法实现
  * 插入排序
  * 快速排序
  * 堆排
* sort包中最上层是一个名称为Interface的接口,只要满足sort.Interface类型都可以排序
```go
// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```

* Go语言标准库默认提供了对int、float64、string进行排序的API
* 很多函数的参数都是sort包下类型,需要进行转换.
*/

func main() {

	// 对int类型切片排序
	num := []int{1, 7, 5, 2, 6}
	// 升序排序
	sort.Ints(num)
	fmt.Println(num)
	// 降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)

	// 对float类型进行排序
	f := []float64{1.5, 7.2, 5.8, 2.3, 6.9}
	// 升序排序
	sort.Float64s(f)
	fmt.Println(f)
	// 降序排序
	sort.Sort(sort.Reverse(sort.Float64Slice(f)))
	fmt.Println(f)

	// 对string类型切片排序
	s := []string{"我", "我是中国人", "a", "d", "国家", "你", "我a"}
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s) // [a d 你 国家 我 我a 我是中国人]
	// sort.Strings(s)
	// fmt.Println(s) // [a d 你 国家 我 我a 我是中国人]

	// 查找内容的索引，如果不存在，返回内容应该在升序排序切片的哪个位置插入
	n := sort.SearchStrings(s, "远")
	fmt.Println(n)
	// 降序排序
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
