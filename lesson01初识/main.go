// 包名
package main

// 引入包
import (
	"fmt"
	"os"
	//	"os"
)

func main() {
	const name, dept = "GeeksForGeeks", "CS"
	fmt.Fprintln(os.Stdout, name, "is a", dept, "portal")
	fmt.Fprint(os.Stdout, name, "is a", dept, "portalss")

}

/*
go 静态语言
支持类型推导
*/
