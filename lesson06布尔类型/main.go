package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 创建布尔类型
	var a bool = true
	var b bool = false
	var c = true
	d := false
	fmt.Println(a, b, c, d)

	// 查看占用byte
	a1 := false
	var b1 byte = 1
	var c1 int8 = 0
	var d1 int = 1
	fmt.Println(unsafe.Sizeof(a1), unsafe.Sizeof(b1), unsafe.Sizeof(c1), unsafe.Sizeof(d1))

	// 相互转换
	var a2 int8 = 1
	var b2 byte = 1
	var c2 bool = true

	// a2 = int8(c2)
	// b2 = byte(c2)
	// c2 = bool(a2)
	// c2 = bool(b2)
	// b2 = byte(a2)
	// b2 = int8(b2)

	a3 := 5 < 3

	fmt.Println(a2, b2, c2, a3)
}
