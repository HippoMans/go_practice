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
		l.root = &Node
	}
}
