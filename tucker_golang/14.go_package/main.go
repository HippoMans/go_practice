package main

import (
	"go_package/Architecture"
)

func main() {
	list := &Architecture.LinkedList{}
	list.AddNode(10)
	list.AddNode(20)
	list.AddNode(30)
	list.PrintNode()
	list.PrintReverseNode()
	list.AddNode(10000)
	list.AddNode(20000)
	list.AddNode(30000)
	list.PrintNode()
	list.PrintReverseNode()
	list.RemoveNode(list.Root.Next)
	list.RemoveNode(list.Tail.Prev)
	list.PrintNode()
	list.PrintReverseNode()
}
