package _map

import (
	"fmt"
	"sort"
	"strings"
)

// Go语言中提供的映射关系容器为map，其内部使用散列表（hash）实现
// map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用

func Demo() {

	// 定义一个map变量 未分配内存，map类型的变量默认初始值为nil
	var m map[string]string

	fmt.Println(m == nil) // true

	// m["name"] = "name"    // panic: assignment to entry in nil map
	// fmt.Printf("v%", m)

	// 使用 make() 函数来分配内存
	m = make(map[string]string, 10) // map 的容量，不是必须但是 建议在初始化map的时候就为其指定一个合适的容量
	fmt.Printf("%v \n", m)          // map[] key,value 为空字符串
	fmt.Println(m[""])              // 空字符串

	// 使用 make 函数初始化并且分配内存，slice和function不能作为Key的类型，value 任何类型都可以
	m1 := make(map[int]int)
	m1[1] = 1
	m1[2] = 2
	fmt.Println(m1)

	// 直接初始化一个 map
	m2 := map[string]string{
		"key": "value",
	}
	fmt.Println(m2)

	// 判断 map 某个键是否存在
	v, ok := m2["key"]
	if ok {
		fmt.Printf("存在，值为 %v", v)
	} else {
		fmt.Println("不存在")
	}

	// map 遍历使用 for range
	m3 := make(map[string]string, 2)
	m3["username"] = "eric"
	m3["address"] = "chengdu"
	for key, value := range m3 {
		fmt.Println(key, value)
	}

	// 使用 delete() 函数删除键值对
	m4 := make(map[string]string, 2)
	m4["a"] = "a"
	m4["b"] = "b"
	delete(m4, "a")
	fmt.Println(m4)

	// 按照指定顺序遍历 map
	m5 := make(map[int]int)
	m5[0] = 1
	m5[1] = 2
	m5[2] = 3
	m5[3] = 4
	m5[4] = 5

	nums := make([]int, len(m5))
	for k, v := range m5 {
		nums[k] = v
	}
	sort.Ints(nums) // 使用标准 sort 排序库，将 nums 从小到大排序
	for i := 0; i < len(nums); i++ {
		fmt.Println(m5[i]) // 按照从小到大顺序打印 map 的 value 值
	}

	// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
	str := "how do you do"
	fmt.Println(countWord(str))
}

// countWord 统计每个单词出现的次数
func countWord(str string) map[string]int {
	m := make(map[string]int)
	s := strings.Split(strings.Trim(str, " "), " ")
	for _, v := range s {
		m[v] += 1
	}
	return m
}
