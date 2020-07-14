package 面试题

import (
	"learningNote/数据结构/ArrayStack"
	"learningNote/数据结构/Stack"
	"learningNote/数据结构/Utils/Interfaces"
)

/**
仅用递归函数和栈操作逆序一个栈
*/

//将栈底元素返回并移除
func getAndRemoveLastElement(stack Stack.Stack) interface{} {
	result := stack.Pop()
	if stack.IsEmpty() {
		return result
	}
	last := getAndRemoveLastElement(stack)
	stack.Push(result)
	return last

}

//逆序栈元素
func reverse(stack Stack.Stack) {
	if stack.IsEmpty() {
		return
	}
	i := getAndRemoveLastElement(stack)
	reverse(stack)
	stack.Push(i)
}

/**
使用一个栈实现另一个栈的排序
*/
func SortStackByStack(stack Stack.Stack){
	help := ArrayStack.Constructor(stack.GetSize())
	for !stack.IsEmpty() {
		cur := stack.Pop()
		for !help.IsEmpty() && Interfaces.Compare(help.Peek(), cur) < 0 {
			stack.Push(help.Pop())
		}
		help.Push(cur)
	}

	for !help.IsEmpty() {
		stack.Push(help.Pop())
	}
}
