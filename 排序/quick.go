package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 7, 6, 5, 3, 2, 8, 1}
	sort1 := quickSort(arr, 1)
	fmt.Println(sort1)

	sort2 := quickSort(arr, 2)
	fmt.Println(sort2)
}

//快速排序 双边循环
func quickSort(arr []int, sortType int) []int {
	if len(arr) <= 0 {
		panic("数组长度小于0")
	}
	newArr := make([]int, len(arr))
	copy(newArr, arr)

	_quickSort(newArr, 0, len(arr)-1, sortType)
	return newArr
}

func _quickSort(arr []int, startIndex, endIndex int, sortType int) {
	if startIndex >= endIndex {
		return
	}
	//基准元素位置
	pivotIndex := 0
	switch sortType {
	case 1:
		pivotIndex = partition1(arr, startIndex, endIndex)
		_quickSort(arr, startIndex, pivotIndex-1, sortType)
		_quickSort(arr, pivotIndex, endIndex, sortType)
		break
	case 2:
		pivotIndex = partition2(arr, startIndex, endIndex)
		_quickSort(arr, startIndex, pivotIndex-1, sortType)
		_quickSort(arr, pivotIndex+1, endIndex, sortType)
		break
	default:
		panic("err type")
	}
}

/**
分治 （双边循环）
返回基准元素的小标
*/
func partition1(arr []int, startIndex int, endIndex int) int {
	pivot := arr[startIndex]
	left := startIndex
	right := endIndex

	for left != right {
		for left < right && arr[left] <= arr[pivot] {
			left++
		}

		for left < right && arr[right] >= arr[pivot] {
			right--
		}
		//交换left和right的位置
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	//基准元素和left、right位置重叠的位置交换
	arr[pivot], arr[left] = arr[left], arr[pivot]

	return left
}

/*
单边循环
*/
func partition2(arr []int, startIndex int, endIndex int) int {
	pivot := arr[startIndex]
	mark := startIndex
	for i := startIndex + 1; i <= endIndex; i++ {
		if arr[i] < pivot {
			mark++
			arr[i], arr[mark] = arr[mark], arr[i]
		}
	}
	arr[startIndex], arr[mark] = arr[mark], pivot
	return mark
}
