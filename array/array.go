package array

import "fmt"

// 数组:

// 数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。
// 数组的长度是数组类型的组成部分。
// 因为数组的长度是数组类型的一个部分，
// 不同长度或不同类型的数据组成的数组都是不同的类型，
// 因此在Go语言中很少直接使用数组（不同长度的数组因为类型不同无法直接赋值.

func Demo() {

	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6
	fmt.Println(a, b, c, d)

	// 遍历数组使用 for range
	for k, v := range d {
		fmt.Println(k, v)
	}

	// 遍历数组使用 for
	for i := 0; i < len(d); i++ {
		fmt.Println(i, d[i])
	}

	// d = append(b, 1) // 向固定长度数组追加长度会报错: first argument to append must be slice; have [3]int

	// 为了避免复制数组带来的开销，可以传递一个指向数组的指针，但是数组指针并不是数组
	a1 := [3]int{1, 2, 3}
	a2 := &a1
	fmt.Println(a1, a2, a2[0])
}
