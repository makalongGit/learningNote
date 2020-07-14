package LinkedList

import (
	"fmt"
)

type Node struct {
	e    interface{}
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}

type LinkedList struct {
	head *Node
	size int
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) AddFirst(e interface{}) {
	l.head = &Node{e, l.head}
	l.size++
}
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("index < 0")
	}
	if index == 0 {
		l.AddFirst(e)
	} else {
		preNode := l.head
		for i := 0; i < index-1; i++ {
			preNode = preNode.next
		}
		node := &Node{e: e}
		node.next = preNode.next
		preNode.next = node
		l.size++
	}
}
func (l *LinkedList) AddLast(e interface{}) {
	l.AddFirst(e)
}
