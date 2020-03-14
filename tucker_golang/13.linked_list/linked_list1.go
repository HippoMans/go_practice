package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	root = &Node{nil, 0}

	for i := 1; i <= 10; i++ {
		AddNode(root, i*10)
	}
	PrintNode(root)
}

func AddNode(root *Node, val int) {
	var Tail *Node
	Tail = root

	for Tail.next != nil {
		Tail = Tail.next
	}
	var Last_node *Node
	Last_node = &Node{nil, val}
	Tail.next = Last_node
}

func PrintNode(root *Node) {
	var Tail *Node
	Tail = root
	for i := 0; Tail.next != nil; i++ {
		fmt.Printf("노드의 순서 번호 : %d, 값 : %d\n", i, Tail.val)
		Tail = Tail.next
	}
}
