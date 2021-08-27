package slice

import "fmt"

// 切片就是一种简化版的动态数组。
// 因为动态数组的长度是不固定，切片的长度自然也就不能是类型的组成部分了。
// 数组虽然有适用它们的地方，但是数组的类型和操作都不够灵活，因此在Go代码中数组使用的并不多。

var (
	a []int               // nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
	b = []int{}           // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
	c = []int{1, 2, 3}    // 有3个元素的切片, len和cap都为3
	d = c[:2]             // 有2个元素的切片, len为2, cap为3
	e = c[0:2:cap(c)]     // 有2个元素的切片, len为2, cap为3
	f = c[:0]             // 有0个元素的切片, len为0, cap为3
	g = make([]int, 3)    // 有3个元素的切片, len和cap都为3
	h = make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
	i = make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
)

func Demo() {

	// 遍历切片的方式和遍历数组的方式类似

	// 1.for
	s1 := []string{"a", "b", "c", "d"}
	for i := 0; i < len(s1); i++ {
		fmt.Println(i, s1[i])
	}
	// 2.for range
	for i, v := range s1 {
		fmt.Println(i, v)
	}
	// 3.for range
	for i := range s1 {
		fmt.Println(i, s1[i])
	}

	// 添加元素到切片
	var s2 []int // 申明一个为 nil 的片，其 len:0  cap:0
	// s2[0] = 1    // panic: runtime error: index out of range [0] with length 0

	// append 可以在切片的尾部追加N个元素
	s2 = append(s2, 1)                 // 追加一个元素
	s2 = append(s2, 1, 2, 3)           // 手写解包方式
	s2 = append(s2, []int{1, 2, 3}...) // ...对切片元素解包变成 1,2,3
	fmt.Printf("%v\n", s2)

	// 切片的开头添加元素
	// 开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。
	// 因此，从切片的开头添加元素的性能一般要比从尾部追加元素的性能差很多。
	var s3 = []int{1, 2, 3}
	s3 = append([]int{0}, s3...) // 在开头添加1个元素
	fmt.Println(s3)
	s3 = append([]int{-3, -2, -1}, s3...) // 在开头添加1个切片
	fmt.Println(s3)

	// 复制切片到目标切片
	dst := []int{1, 2, 3}
	src := []int{1: 4, 2: 5, 3: 6}
	copy(dst, src)   // 复制 src 到 dst
	fmt.Println(dst) // [0, 4, 5]
	dst1 := []int{1, 2, 3}
	src1 := []int{4, 5}
	copy(dst1, src1)  // 复制 src 到 dst
	fmt.Println(dst1) // [4, 5, 3]

	// 删除切片元素
	// 删除元素的位置有三种情况：从开头位置删除，从中间位置删除，从尾部删除。其中删除切片尾部的元素最快
	d1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d2 := d1[1:]           // 删除第一个元素
	fmt.Printf("%v\n", d2) // [2 3 4 5 6 7 8 9 10] 索引从新构建从0开始
	d3 := d1[0 : len(d1)-1]
	fmt.Printf("%v\n", d3) // 删除最后一个元素

	d4 := d3[:0]
	fmt.Println(d4 == nil)
	fmt.Println(d4, len(d4), cap(d4))

	// 当原切片长度小于1024时，新切片的容量会直接翻倍。
	t := make([]int, 0, 2)
	fmt.Println(len(t), cap(t)) // 0 2

	t = append(t, 1, 2, 3)
	fmt.Println(len(t), cap(t)) // 3 4

	t = append(t, 2, 3)
	fmt.Println(len(t), cap(t)) // 5 8

	// 而当原切片的容量大于等于1024时，会反复地增加25%，直到新容量超过所需要的容量
	for i := 0; i < 1024; i++ {
		t = append(t, i)
	}

	fmt.Println(len(t), cap(t)) // 1029 1536
}

// 切片高效操作的要点是要降低内存分配的次数，尽量保证append操作不会超出cap的容量，降低触发内存分配的次数和每次分配内存大小。
// TrimSpaces 删除空格利用空切片但是cap不为空
func TrimSpaces(s []byte) []byte {
	b := s[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}
