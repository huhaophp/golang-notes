package structure

// 数据结构-栈篇

// Stack 存放整型数据
type Stack struct {
	items []int
}

// NewStack 实例化栈
func NewStack() *Stack {
	return &Stack{items: []int{}}
}

// Push 向栈中放入元素
func (s *Stack) Push(n int) {
	s.items = append(s.items, n)
}

// Pop 取出栈中最后进入栈的值
func (s *Stack) Pop() (n int) {
	n = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return
}

// Len 栈长度
func (s *Stack) Len() int {
	return len(s.items)
}

// IsEmpty 栈是否为为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
