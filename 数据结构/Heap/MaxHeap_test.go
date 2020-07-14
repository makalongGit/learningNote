package Heap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMaxHeap_Add(t *testing.T) {
	n := 10
	maxHeap := New()
	for i := 0; i < n; i++ {
		maxHeap.Add(rand.Intn(1000))
	}
	for i := 0; i < n; i++ {
		fmt.Println(maxHeap.data)
		fmt.Println(maxHeap.ExtractMax())
	}
}
