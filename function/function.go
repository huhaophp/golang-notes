package function

import "fmt"

var (
	v = 10
)

func Demo() {
	// 具名函数
	fmt.Println(Add(1, 2))

	// 匿名删除
	fmt.Println(Sub(1, 2))

	// 多返回值
	a, b := Swap(1, 2)
	fmt.Println(a, b)

	// 可变参数
	s1 := Sum(1, []int{1, 2}...) // ... 解包
	s2 := Sum(1, 2, 3, 4)
	fmt.Println(s1, s2)

	// 空接口参数
	i1 := []interface{}{123, "abc"}
	Print(i1...)
	Print(i1)

	// 函数的返回值命名
	m1 := map[int]int{100: 100}
	value, ok := Find(m1, 100)
	fmt.Println(value, ok) // 100 true

	fmt.Println(Inc())
}

// Add 具名函数
func Add(a, b int) int {
	return a + b
}

// Sub 匿名函数
var Sub = func(a, b int) int {
	return a - b
}

// 多个参数和多个返回值
func Swap(a, b int) (int, int) {
	return a, b
}

// 可变数量的参数
// more 对应 []int 切片类型或者 1,2,3 多个参数
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

// Print 当可变参数是一个空接口类型时，调用者是否解包可变参数会导致不同的结果：
func Print(a ...interface{}) {
	fmt.Println(a...)
}

// 不仅函数的参数可以有名字，也可以给函数的返回值命名
func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

// 如果返回值命名了，可以通过名字来修改返回值，也可以通过defer语句在return语句之后修改返回值
func Inc() (v int) {
	defer func() {
		v++
	}()
	return 42
}
