package 面试题

import (
	"fmt"
	"testing"
)

func TestRemoveLastKthNode(t *testing.T) {
	param := []int{1, 2, 3, 4, 5}
	head := &Node{}
	cur := head
	for i := 0; i < len(param); i++ {
		if i == 0 {
			cur.value = param[i]
		} else {
			cur.next = &Node{
				value: param[i],
				next:  nil,
			}
			cur = cur.next
		}
	}
	cur1 := head
	fmt.Print(cur1)
	for cur1.next != nil {
		fmt.Print(cur1.next)
		cur1 = cur1.next
	}

	RemoveLastKthNode(head, 2)

	cur1 = head
	fmt.Print(cur1)
	for cur1.next != nil {
		fmt.Print(cur1.next)
		cur1 = cur1.next
	}


}
