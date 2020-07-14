package Array

import (
	"bytes"
	"fmt"
	"learningNote/数据结构/Utils/Interfaces"
)

type Array struct {
	data []interface{}
	size int //容量
}

//构造函数
func Constructor(capacity int) *Array {
	return &Array{data: make([]interface{}, capacity)}
}

//获取数组容量
func (array *Array) GetCapacity() int {
	return len(array.data)
}

//获取数组中的元素个数
func (array *Array) GetSize() int {
	return array.size
}

//数组是否为空
func (array *Array) IsEmpty() bool {
	return array.size == 0
}

//在第index个位置添加一个新元素
func (array *Array) Add(index int, value interface{}) {
	if index < 0 || index > array.GetCapacity() {
		panic("out of range")
	}
	if array.size == len(array.data) {
		array.resize(2 * array.size)
	}
	for i := array.size - 1; i >= index; i-- {
		array.data[i+1] = array.data[i]
	}
	array.data[index] = value
	array.size++
}

func (array *Array) AddLast(e interface{}) {
	array.Add(array.GetSize(), e)
}
func (array *Array) AddFirst(e interface{}) {
	array.Add(0, e)
}

func (array *Array) Get(index int) interface{} {
	if index < 0 || index >= array.size {
		panic("out of range")
	}
	return array.data[index]
}

func (array *Array) GetFirst() interface{} {
	return array.Get(0)
}

func (array *Array) GetLast() interface{} {
	return array.Get(array.size - 1)
}

// 修改 index 索引位置的元素
func (array *Array) Set(index int, e interface{}) {
	if index < 0 || index >= array.size {
		panic("Set failed. Index is illegal.")
	}
	array.data[index] = e
}

//查找数组中是否有元素e
func (array *Array) Contains(e interface{}) bool {
	for i := 0; i < array.size; i++ {
		if Interfaces.Compare(array.data[i], e) == 0 {
			return true
		}
	}
	return false
}

// 查找数组中元素 e 所在的索引，不存在则返回 -1
func (array *Array) Find(e interface{}) int {
	for i := 0; i < array.size; i++ {
		if Interfaces.Compare(array.data[i], e) == 0 {
			return i
		}
	}
	return -1
}

// 查找数组中元素 e 所有的索引组成的切片，不存在则返回 -1
func (array *Array) FindAll(e interface{}) (indexes []int) {
	for i := 0; i < array.size; i++ {
		if Interfaces.Compare(array.data[i], e) == 0 {
			indexes = append(indexes, i)
		}
	}
	return
}

// 从数组中删除 index 位置的元素，返回删除的元素
func (array *Array) Remove(index int) interface{} {
	if index < 0 || index >= array.size {
		panic("Remove failed,Index is illegal.")
	}

	e := array.data[index]
	for i := index + 1; i < array.size; i++ {
		array.data[i-1] = array.data[i]
	}
	array.size--
	array.data[array.size] = nil //loitering object != memory leak

	// 考虑边界条件，避免长度为 1 时，resize 为 0
	if array.size == len(array.data)/4 && len(array.data)/2 != 0 {
		array.resize(len(array.data) / 2)
	}
	return e
}

// 从数组中删除第一个元素，返回删除的元素
func (array *Array) RemoveFirst() interface{} {
	return array.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (array *Array) RemoveLast() interface{} {
	return array.Remove(array.size - 1)
}

// 从数组中删除一个元素 e
func (array *Array) RemoveElement(e interface{}) {
	index := array.Find(e)
	if index != -1 {
		array.Remove(index)
	}
}

// 从数组中删除所有元素 e
func (array *Array) RemoveAllElement(e interface{}) bool {
	if array.Find(e) == -1 {
		return false
	}

	for i := 0; i < array.size; i++ {
		if Interfaces.Compare(array.data[i], e) == 0 {
			array.Remove(i)
		}
	}
	return true
}

func (array *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < array.size; i++ {
		newData[i] = array.data[i]
	}
	array.data = newData
}

func (array *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", array.size, len(array.data)))
	buffer.WriteString("[")
	for i := 0; i < array.size; i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(array.data[i]))
		if i != (array.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
