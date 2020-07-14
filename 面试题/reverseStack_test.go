package 面试题

import (
	"fmt"
	"learningNote/数据结构/ArrayStack"
	"testing"
)

func TestReverseStack(t *testing.T) {
	param := []int{3, 4, 5, 1, 2, 1}
	stack := ArrayStack.Constructor(len(param))
	for i := 0; i < len(param); i++ {
		stack.Push(param[i])
	}
	fmt.Println(stack)
	reverse(stack)
	fmt.Println(stack)
}

func TestSortStackByStack(t *testing.T){
	param := []int{3, 4, 5, 1, 2, 1}
	stack := ArrayStack.Constructor(len(param))
	for i := 0; i < len(param); i++ {
		stack.Push(param[i])
	}

	SortStackByStack(stack)
	fmt.Println(stack)
}
