package 面试题

import (
	"fmt"
	"learningNote/数据结构/ArrayStack"
	"testing"
)

func TestMyStack(t *testing.T) {
	param := []int{3, 4, 5, 1, 2, 1}
	myStack := MyStack{
		stackData: ArrayStack.Constructor(len(param)),
		stackMin:  ArrayStack.Constructor(len(param)),
	}

	for i := 0; i < len(param); i++ {
		myStack.Push(param[i])

		fmt.Println(myStack.stackMin)
	}
}
