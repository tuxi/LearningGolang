package main

import (
	"container/list"
	"fmt"
)

func main() {

	/*
		 list 双向链表
		使用container/list包下面的New()创建一个新的空list
	*/

	myList := list.New()
	fmt.Println(myList)        // &{{0xc00010e030 0xc00010e030 <nil> <nil>} 0}
	fmt.Println(myList.Len())  // 0
	fmt.Printf("%p\n", myList) // 0xc00010e030

	/*
	 对list进行操作
	*/
	// 添加元素到链表结尾
	myList.PushBack("a") // [a]
	// 添加元素到链表头部
	myList.PushFront("b") // [b a]
	// 向第一个元素后面添加一个元素
	myList.InsertAfter("c", myList.Front()) // [b c a]
	// 向最后一个元素前面添加一个元素
	myList.InsertBefore("e", myList.Back()) // [b c e a]

	// 取出链表中的元素
	fmt.Println(myList.Back().Value)  // a
	fmt.Println(myList.Front().Value) // b

	// 根据位置n查找链表中的元素
	n := 2
	var current *list.Element
	if n >= 0 && n < myList.Len() {
		if n == 0 {
			current = myList.Front()
		} else if n == myList.Len()-1 {
			current = myList.Back()
		} else {
			current = myList.Front()
			for i := 0; i < n; i++ {
				current = current.Next()
			}
		}
	} else {
		fmt.Println("n的数值不对哦")
	}

	fmt.Println(current.Value) //  e

	// 移动元素的顺序
	// 将第一个元素移动到最后
	myList.MoveToBack(myList.Front())
	// 将最后一个元素移动到头部
	myList.MoveToFront(myList.Back())
	// 将第一个参数元素，移动到第二个参数元素的前面
	myList.MoveBefore(myList.Front(), myList.Back())
	// 将第一个参数元素，移动到第二个元素的后面
	myList.MoveAfter(myList.Back(), myList.Front())

	// 删除最后一个元素
	myList.Remove(myList.Back())

	// 遍历所有链表
	for curr := myList.Front(); curr != nil; curr = curr.Next() {
		fmt.Println(curr.Value)
	}

}
