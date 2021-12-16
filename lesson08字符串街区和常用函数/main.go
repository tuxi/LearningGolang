package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "yangxiaoyuan"
	fmt.Println(s)

	// 截取字符串

	s1 := s[5:]
	fmt.Println(s1)
	fmt.Printf("s1的类型：%T\n", s1)

	s2 := fmt.Sprintf("%s", s1)
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)

	// 字符串长度和获取字符
	arr := []rune(s)
	fmt.Println(len(arr))
	fmt.Println(arr[9])
	fmt.Printf("%T\n", arr[9])

	// 遍历字符
	for _, n := range s {
		fmt.Printf("%c\n", n)
	}

	// 获取某个字符在字符串中第一次出现的索引
	fmt.Println(strings.Index(s, "y"))
	// 获取某个字符在字符串中最后一次出现的索引
	fmt.Println(strings.LastIndex(s, "u"))
	// 是否以指定内容开头
	fmt.Println(strings.HasPrefix(s, "yang"))
	// 是否以某个字符结尾
	fmt.Println(strings.HasSuffix(s, "yuan"))
	// 是否包含某个字符
	fmt.Println(strings.Contains(s, "jjj"))
	// 全部转为小写
	fmt.Println(strings.ToLower(s))
	// 全部转为大写
	fmt.Println(strings.ToUpper(s))
	// 替换, 如果n小于0表示全部替换，
	//如果n大于old个数也表示全部替换
	fmt.Println(strings.Replace(s, "yu", "aaaaa", -1))
	// 把字符串重复n遍
	fmt.Println(strings.Repeat(s, 3))
	// 去掉字符串前后的字符
	fmt.Println(strings.Trim(s, " "))
	// 去掉前后空格
	fmt.Println(strings.TrimSpace(s))
	// 根据指定字符把字符串拆分成切片
	fmt.Println(strings.Split(s, "y"))
	// 使用指定分隔符把切片内容合并成字符串
	arr1 := []string{"xiao", "yuan"}
	fmt.Println(arr1)
	fmt.Println(strings.Join(arr1, ""))
}
