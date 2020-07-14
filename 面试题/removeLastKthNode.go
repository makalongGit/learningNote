package 面试题

import "fmt"

/**
在单链表和双链表中删除倒数第K个节点
题目:
分别实现两个函数，一个可以删除单链表中倒数第K个节点，另一个可以删除双链表中倒数第K个节点
*/

type Node struct {
	value int
	next  *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%d  ", n.value)
}

func RemoveLastKthNode(head *Node, lastKth int) *Node {
	if head == nil || lastKth < 0 {
		return head
	}
	index := lastKth
	cur := head
	for cur.next != nil {
		index--
		cur = cur.next
	}
	if index == 0 {
		head = head.next
	}
	if index < 0 {
		cur = head
		for index != 0 {
			cur = cur.next
			index ++
		}
		cur.next = cur.next.next
	}
	return head
}
