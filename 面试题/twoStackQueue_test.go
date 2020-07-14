package 面试题

import (
	"fmt"
	"learningNote/数据结构/ArrayStack"
	"testing"
)

func TestStackQueue(t *testing.T){
	param := []int{3, 4, 5, 1, 2, 1}
	queue := StackQueue{
		ArrayStack.Constructor(len(param)),
		ArrayStack.Constructor(len(param)),
	}
	for i:=0;i<len(param);i++{
		queue.add(param[i])
	}
	for i := 0; i < len(param); i++ {
		fmt.Println(queue.poll())
	}
}