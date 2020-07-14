package 面试题

import "learningNote/数据结构/Stack"

/**
由两个栈组成的队列
*/

type StackQueue struct {
	stackPush Stack.Stack //压入栈
	stackPop  Stack.Stack //弹出栈
}

func (sq *StackQueue) pushToPop() {
	if sq.stackPop.IsEmpty() {
		for !sq.stackPush.IsEmpty() {
			sq.stackPop.Push(sq.stackPush.Pop())
		}
	}
}

func (sq *StackQueue) add(newNum interface{}) {
	sq.stackPush.Push(newNum)
	sq.pushToPop()
}

func (sq *StackQueue) poll() interface{} {
	if sq.stackPush.IsEmpty() && sq.stackPop.IsEmpty() {
		panic("queue is empty")
	}
	sq.pushToPop()
	return sq.stackPop.Pop()
}

func (sq StackQueue) peek() interface{} {
	if sq.stackPush.IsEmpty() && sq.stackPop.IsEmpty() {
		panic("queue is empty")
	}
	sq.pushToPop()
	return sq.stackPop.Peek()
}
