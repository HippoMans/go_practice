package Architecture

type Queue struct {
	list *Linkedlist
}

func NewQueue() *Queue {
	return &Queue{&Linkedlist{}}
}

func (q *Queue) Push(val int) {
	q.list.AddNode(val)
}

func (q *Queue) Pop() int {
	queue_value := q.list.Front()
	q.list.PopFront()
	return queue_value
}

func (q *Queue) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue) PrintQueue() {
	q.list.PrintNode()
}
