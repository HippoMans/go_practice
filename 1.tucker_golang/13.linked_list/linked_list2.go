package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	//처음 root의 노드
	var root *Node
	root = &Node{nil, 0}

	//Node.next에 nil이 있는 node를 가리키는 것을 tail 노드
	var tail *Node
	tail = root

	var want_end int
	fmt.Print("몇 번째까지 추가를 원하나요?  : ")
	fmt.Scanf("%d", &want_end)

	for i := 1; i <= want_end; i++ {
		tail = AddNode(tail, i*10)
	}
	PrintAllNode(root)
	LastNode(root)
}

func AddNode(node *Node, val int) *Node {
	var tail *Node
	tail = &Node{nil, val}
	node.next = tail
	return tail
}

func PrintAllNode(root *Node) {
	var i int
	for i = 0; root.next != nil; i++ {
		fmt.Printf("노드의 위치 : %d, 노드값 : %d\n", i, root.val)
		root = root.next
	}
	fmt.Printf("노드의 위치 : %d, 노드값 : %d\n", i, root.val)
}

func LastNode(root *Node) {
	var tail *Node
	tail = root
	for tail.next != nil {
		tail = tail.next
	}
	if tail.next == nil {
		fmt.Printf("마지막 노드 값 : %d\n", tail.val)
	}
}
