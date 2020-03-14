package Architecture

import "fmt"

type Node struct {
	Prev *Node
	val  int
	Next *Node
}
type LinkedList struct {
	Root *Node
	Tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.Root == nil {
		l.Root = &Node{l.Root, val, nil}
		l.Tail = l.Root
		return
	}
	l.Tail.Next = &Node{l.Tail, val, nil}
	prev_node := l.Tail
	l.Tail = prev_node.Next
	l.Tail.Prev = prev_node
}

func (l *LinkedList) Back() int {
	if l.Tail != nil {
		return l.Tail.val
	}
	return 0
}

func (l *LinkedList) PopBack() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
}

func (l *LinkedList) RemoveNode(node *Node) {
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
		prev_node.Next.Prev = prev_node
	}
	node.Next = nil
}

func (l *LinkedList) PrintNode() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.val)
}

func (l *LinkedList) PrintReverseNode() {
	node := l.Tail
	for l.Tail.Prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.Prev
	}
	fmt.Printf("%d\n", node.val)
}

func (l *LinkedList) IsEmpty() bool {
	return l.Root == nil
}
