package Architecture

import "fmt"

type Node struct {
	Prev *Node
	val  int
	Next *Node
}

type Linkedlist struct {
	Root *Node
	Tail *Node
}

func (l *Linkedlist) AddNode(val int) {
	if l.Root == nil {
		l.Root = &Node{l.Root, val, nil}
		l.Tail = l.Root
		return
	}
	l.Tail.Next = &Node{nil, val, nil}
	prev_node := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = prev_node
}

func (l *Linkedlist) Back() int {
	if l.Tail != nil {
		return l.Tail.val
	}
	return 0
}

func (l *Linkedlist) Front() int {
	if l.Root != nil {
		return l.Root.val
	}
	return 0
}

func (l *Linkedlist) PopBack() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
}

func (l *Linkedlist) PopFront() {
	if l.Root == nil {
		return
	}
	l.RemoveNode(l.Root)
}

func (l *Linkedlist) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = node.Next
		if l.Root != nil {
			l.Root.Prev = nil
		}
		node.Next = nil
		return
	}

	prev_node := node.Prev
	if node == l.Tail {
		prev_node.Next = nil
		l.Tail.Prev = nil
		l.Tail = prev_node
	} else {
		node.Prev = nil
		prev_node.Next = prev_node.Next.Next
		prev_node.Next = prev_node
	}
	node.Next = nil
}
func (l *Linkedlist) PrintNode() {
	node := l.Root
	fmt.Print("정방향 : ")
	for node.Next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.val)
}

func (l *Linkedlist) PrintReverseNode() {
	node := l.Tail
	fmt.Print("역방향 : ")
	for node.Prev != nil {
		fmt.Printf("%d -> ", node.val)
	}
	fmt.Printf("%d\n", node.val)
}

func (l *Linkedlist) IsEmpty() bool {
	return l.Root == nil
}
