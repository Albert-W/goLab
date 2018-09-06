package queue

// A FIFO queue.
type Queue []int  //实际上是slice
//type Queue []interface{}  //任务类型。

// Pushes the element into the queue.
// 		e.g. q.Push(123)
//需要使用指针才能生效
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pops element from head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
	//return head.(int) // 强制转换为int;
}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
