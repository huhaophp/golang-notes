package structure

// 数据结构-队列篇

// Queue 存放整型数据
type Queue struct {
	items []int
}

// NewQueue 实例化队列
func NewQueue() *Queue {
	return &Queue{items: []int{}}
}

// Push 向队列放入元素
func (q *Queue) Push(n int) {
	q.items = append(q.items, n)
}

// Pop 从队列中取出最先进入队列的值
func (q *Queue) Pop() (n int) {
	n = q.items[0]
	q.items = q.items[1:]
	return
}

// Len 队列长度
func (q *Queue) Len() int {
	return len(q.items)
}

// IsEmpty 是否为为空
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
