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
		l.Root.Next = nil
		l.Tail = l.Root
		return
	}

	l.Tail.Next = &Node{l.Tail, val, nil}
	prev_node := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = prev_node
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = node.Next
		l.Root.Prev = nil
		node.Next.Prev = l.Root
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
		prev_node.Next = prev_node.Next.Next //prev_node.Next = node => prev_node.Next = node.next
		prev_node.Next.Prev = prev_node      //warning prev_node.Next is prev_node.Next.Next
		//prev_node.Next.Next is not right
	}
	node.Next = nil
}

func (l *LinkedList) PrintNode() {
	var node *Node
	node = l.Root
	fmt.Print("정방향 : ")
	for node.Next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.val)
}

func (l *LinkedList) PrintReverseNode() {
	var node *Node
	node = l.Tail
	fmt.Print("역방향 : ")
	for node.Prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.Prev
	}
	fmt.Printf("%d\n", node.val)
}
