package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{nil, val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{nil, val}
	l.tail = l.tail.next
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = node.next
		return
	}
	var prev *Node
	prev = l.root

	for prev.next != node {
		prev = prev.next
	}

	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		prev.next = node.next
	}
	node.next = nil
}

func (l *LinkedList) PrintNode() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func main() {
	list := &LinkedList{}
	list.AddNode(10)
	list.AddNode(20)
	list.AddNode(30)
	list.PrintNode()
	list.AddNode(100)
	list.AddNode(200)
	list.AddNode(300)
	list.RemoveNode(list.root.next.next)
	list.PrintNode()
	list.AddNode(10000)
	list.AddNode(20000)
	list.AddNode(30000)
	list.RemoveNode(list.tail)
	list.PrintNode()
}
