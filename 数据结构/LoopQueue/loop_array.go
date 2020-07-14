package LoopQueue

import (
	"bytes"
	"fmt"
)

type loopQueue struct {
	data  []interface{}
	tail  int
	front int
	size  int
}

func Constructor(capacity int) *loopQueue {
	return &loopQueue{
		data: make([]interface{}, capacity+1),
	}
}

func (l *loopQueue) GetCapacity() int {
	//删除一个空位，与上对应
	return len(l.data) - 1
}

func (l *loopQueue) GetSize() int {
	return l.size
}

func (l *loopQueue) IsEmpty() bool {
	return l.tail == l.front
}

//扩容
func (l *loopQueue) resize(NewCapacity int) {
	newData := make([]interface{}, NewCapacity+1)
	for i := 0; i < l.size; i++ {
		newData[i] = l.data[(i+l.front)%len(l.data)]
	}
	l.data = newData
	l.front = 0
	l.tail = l.size
}

//入队
func (l *loopQueue) Enqueue(e interface{}) {
	if (l.tail+1)%len(l.data) == l.front {
		l.resize(l.GetCapacity() * 2)
	}
	l.data[l.tail] = e
	l.tail = (l.tail + 1) % len(l.data)
	l.size++
}

func (l *loopQueue) Dequeue() interface{} {
	if l.IsEmpty() {
		panic("cannot dequeue from an empty queue")
	}
	e := l.data[l.front]
	l.front = (l.front + 1) % len(l.data)
	l.size--
	if l.size == l.GetCapacity()/4 && l.GetCapacity()/2 != 0 {
		l.resize(l.GetCapacity() / 2)
	}
	return e
}

func (l *loopQueue) getFront() interface{} {
	if l.IsEmpty() {
		panic("cannot dequeue from an empty queue")
	}
	return l.data[l.front]
}

func (l *loopQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Queue: size=%d, capacity=%d\n", l.size, l.GetCapacity()))
	buffer.WriteString("front [")
	for i := l.front; i != l.tail; i = (i + 1) % len(l.data) {
		buffer.WriteString(fmt.Sprint(l.data[i]))
		if (i+1)%len(l.data) != l.tail {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")
	return buffer.String()
}
