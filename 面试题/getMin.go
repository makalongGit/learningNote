package 面试题

import (
	"learningNote/数据结构/Stack"
	"learningNote/数据结构/Utils/Interfaces"
)

/**
实现一个特殊的栈，在实现了栈功能的基础上，在实现返回栈中的最小值
*/

type MyStack struct {
	stackData Stack.Stack
	stackMin  Stack.Stack
}

func (stack *MyStack) Push(newNum interface{}) {
	if stack.stackMin.IsEmpty() {
		stack.stackMin.Push(newNum)
	} else if Interfaces.Compare(stack.getMin(), newNum) >= 0 {
		stack.stackMin.Push(newNum)
	}
	stack.stackData.Push(newNum)
}

func (stack *MyStack) Pop() interface{} {
	if stack.stackData.IsEmpty() {
		panic("stack is empty")
	}
	value := stack.stackData.Pop()
	if Interfaces.Compare(stack.getMin(), value) == 0 {
		stack.stackMin.Pop()
	}
	return value
}

func (stack MyStack) getMin() interface{} {
	if stack.stackMin.IsEmpty() {
		panic("stack is Empty")
	}
	return stack.stackMin.Peek()
}
