package main

/*
# 一.双向链表概述

* 双向链表结构如下
  ![双向链表](images/2_5_15_lianbiao.png)

* 双向链表结构中元素在内存中不是紧邻空间,而是每个元素中存放上一个元素和后一个元素的地址

  * 第一个元素称为头(head)元素,前连接(前置指针域)为nil
  * 最后一个元素称为尾(foot)元素,后连接(后置指针域)为nil

* 双向链表的优点:

  * 在执行新增元素或删除元素时效率高,获取任意一个元素,可以方便的在这个元素前后插入元素
  * 充分利用内存空间,实现内存灵活管理
  * 可实现正序和逆序遍历
  * 头元素和尾元素新增或删除时效率较高

* 双向链表的缺点

  * 链表增加了元素的指针域,空间开销比较大
  * 遍历时跳跃性查找内容,大量数据遍历性能低
# 二. 双向链表容器List
* 在Go语言标准库的container/list 包提供了双向链表List
* List结构体定义如下
  * root表示根元素
  * len表示链表中有多少个元素
```Go
// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}
```

* 其中Element结构体定义如下
  * next表示下一个元素,使用Next()可以获取到
  * prev表示上一个元素,使用Prev()可以获取到
  * list表示元素属于哪个链表
  * Value表示元素的值,interface{}在Go语言中表示任意类型
```Go
  // Element is an element of a linked list.
  type Element struct {
  	// Next and previous pointers in the doubly-linked list of elements.
  	// To simplify the implementation, internally a list l is implemented
  	// as a ring, such that &l.root is both the next element of the last
  	// list element (l.Back()) and the previous element of the first list
  	// element (l.Front()).
  	next, prev *Element

  	// The list to which this element belongs.
  	list *List

  	// The value stored with this element.
  	Value interface{}
  }
```

*/

func main() {

}
