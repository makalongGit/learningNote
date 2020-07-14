package ArrayStack

import (
	"bytes"
	"fmt"
	"learningNote/数据结构/Array"
)

type ArrayStack struct {
	array *Array.Array
}

func Constructor(capacity int) *ArrayStack {
	return &ArrayStack{
		array: Array.Constructor(capacity),
	}
}

func (array *ArrayStack) GetSize() int {
	return array.array.GetSize()
}

func (array *ArrayStack) IsEmpty() bool {
	return array.array.IsEmpty()
}

// 压入栈顶元素
func (array *ArrayStack) Push(element interface{}) {
	array.array.AddLast(element)
}

// 弹出栈顶元素
func (array *ArrayStack) Pop() interface{} {
	return array.array.RemoveLast()
}

// 查看栈顶元素
func (array *ArrayStack) Peek() interface{} {
	return array.array.GetLast()
}

func (array *ArrayStack) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	for i := 0; i < array.array.GetSize(); i++ {
		buffer.WriteString(fmt.Sprint(array.array.Get(i)))
		if i != array.array.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")

	return buffer.String()
}
