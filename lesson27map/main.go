package main

import "fmt"

/*
map以散列表的方式存储键值对集合
map中每个元素都是键值对
key是操作map的唯一标志，可以通过key对map进行增删改查
key是唯一的，添加重复key会覆盖原有元素
map是值类型，只声明时为空指针（nil）
map读写数据时，并不是并发安全的，可以结合RWMutex保证并发安全
*/

func main() {

	// 只声明map
	var m map[string]string
	fmt.Println(m)               // map[]
	fmt.Println(m == nil)        // true
	fmt.Printf("%p  %T\n", m, m) // 0x0  map[string]string

	/*
		实例化map的几种方式
	*/

	// 1.使用make函数实例化一个没有初始值的map
	m1 := make(map[string]string)
	fmt.Println(m1)        // map[]
	fmt.Println(m1 == nil) // false
	fmt.Printf("%p\n", m1) // 0xc00007a1b0

	// 声明map时，直接给map赋值
	m2 := map[string]string{"yang": "188101010", "yuan": "w87879"}
	fmt.Println(m2)
	fmt.Println(len(m2))
	fmt.Println(m2["yuan"])

	// 使用key，如果key存在则覆盖原油元素，如果key不存在则新增元素
	m3 := map[string]int{"money": 10}
	m3["money"] = 11
	m3["money"] = 13
	fmt.Println(m3)
	fmt.Println(m3["money"])

	// delete删除元素，如果key存在则删除元素，如果key不存在也不会报错
	m4 := map[string]float64{"yuan": 1.86, "阿拉": 1.98}
	delete(m4, "yuan")
	fmt.Println(m4)

	// 获取map中指定key的值
	// 使用:map变量[key]获取key对应的值
	// 如果key不存在则返回map[key]value中value的默认值，例如value时string类型，则返回空字符串“”
	// 返回值可以是一个也可以是两个： 1个表示key对应的值，两个表示key对应的值和这个key是否存在
	m5 := map[string]string{"yang": "海淀区", "yuan": "房山区"}
	yang := m5["yang"]
	fmt.Println(yang) // 海淀区
	value, ok := m5["yuan"]
	fmt.Println(value, ok) // 房山区

	// 遍历map
	for key, value := range m5 {
		fmt.Println(key, value)
	}
}
