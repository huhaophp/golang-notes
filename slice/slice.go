package slice

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
}
