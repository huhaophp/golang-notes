package main

import (
	"fmt"
	t "note/structure"
)

func main() {
	queue := t.NewQueue()
	queue.Push(1)
	fmt.Println(queue.Len())
}
