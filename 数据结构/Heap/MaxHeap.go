package Heap

import (
	"learningNote/数据结构/Utils/Interfaces"
)

/**
最大堆
使用切片存数据内容
*/

type MaxHeap struct {
	data []interface{}
}

func New() *MaxHeap {
	return &MaxHeap{data: make([]interface{}, 0)}
}

//获取该位置的左孩子的索引位置
func (heap MaxHeap) leftChild(index int) int {
	return index*2 + 1
}

//获取该节点的右孩子索引位置
func (heap MaxHeap) rightChild(index int) int {
	return index*2 + 2
}

//获取该节点的父亲节点索引位置
func (heap MaxHeap) parent(index int) int {
	if index == 0 {
		panic("0 没有父节点")
	}
	return (index - 1) / 2
}

//添加元素
func (heap *MaxHeap) Add(e interface{}) {
	heap.data = append(heap.data, e)
	heap.ShiftUp(len(heap.data) - 1)
}

//上浮该索引的元素
func (heap *MaxHeap) ShiftUp(index int) {
	for index > 0 && Interfaces.Compare(heap.data[index], heap.data[heap.parent(index)]) > 0 {
		heap.data[index], heap.data[heap.parent(index)] = heap.data[heap.parent(index)], heap.data[index]
		index = heap.parent(index)
	}
}

//取出堆顶元素
func (heap *MaxHeap) ExtractMax() interface{} {
	if len(heap.data) == 0 {
		panic("数据为空")
	}
	max := heap.data[0]
	heap.data[0] = heap.data[len(heap.data)-1]
	heap.data = heap.data[:len(heap.data)]
	heap.SiftDown(0)
	return max
}

//下沉改索引的元素
func (heap *MaxHeap) SiftDown(index int) {
	for heap.leftChild(index) < len(heap.data) { //大于就越界，说明已经到了堆低
		j := heap.leftChild(index)
		//取出左右孩子节点最大的那个
		if j+1 < len(heap.data) && Interfaces.Compare(heap.data[j+1], heap.data[j]) > 0 {
			j = heap.rightChild(index)
		}

		if Interfaces.Compare(heap.data[index], heap.data[j]) >= 0 {
			break
		}

		heap.data[index], heap.data[j] = heap.data[j], heap.data[index]
		index = j
	}
}
