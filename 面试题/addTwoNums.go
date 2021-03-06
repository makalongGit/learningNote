package 面试题

/*
给定两个非空链表来代表两个非负整数，位数按照逆序方式存储，它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

type Node struct {
	Val  int
	Next *Node
}

func addTwoNums(l1 *Node, l2 *Node) *Node {
	head := &Node{}
	h := head
	carry := 0
	for l1 != nil || l2 != nil || carry > 0{
		sum := carry
		h.Next = new(Node)
		h = h.Next
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		carry = sum / 10
		v := sum % 10
		h.Val = v
	}
	return head.Next
}
