package main

import "fmt"

type Node struct {
	prev *Node
	val  int
	next *Node
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{l.root, val, nil}
		l.root.prev = nil
		l.tail = l.root
		return
	}
	l.tail.next = &Node{l.tail, val, nil}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = node.next
		l.root.prev = nil
		node.next.prev = l.root
		node.next = nil
		return
	}

	prev := node.prev

	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else {
		node.prev = nil
		prev.next = prev.next.next
		prev.next.prev = prev
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

func (l *LinkedList) PrintReverseNode() {
	node := l.tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}
	fmt.Printf("%d\n", node.val)
}

func main() {
	var list *LinkedList
	list = &LinkedList{}
	list.AddNode(10)
	list.AddNode(20)
	list.AddNode(30)
	fmt.Print("정방향 : ")
	list.PrintNode()
	fmt.Print("역방향 : ")
	list.PrintReverseNode()
	list.AddNode(5000)
	list.AddNode(6000)
	list.AddNode(7000)
	fmt.Print("정방향 : ")
	list.PrintNode()
	fmt.Print("역방향 : ")
	list.PrintReverseNode()
	list.RemoveNode(list.root.next.next)
	list.RemoveNode(list.tail.prev.prev)
	list.PrintNode()
	list.PrintReverseNode()
}
