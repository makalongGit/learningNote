package LoopQueue

import (
	"fmt"
	"testing"
)

func TestLoopQueue(t *testing.T) {
	queue := Constructor(10)
	for i := 0; i < 15; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
		if i%3 == 0 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}
