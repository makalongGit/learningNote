package main

import "fmt"

func main() {
	arr := []int{5, 8, 6, 3, 9, 2, 1, 7}
	sort1 := bubbleSort1(arr)
	fmt.Println(sort1)

	sort2 := bubbleSort2(arr)
	fmt.Println(sort2)

	sort3 := bubbleSort3(arr)
	fmt.Println(sort3)

}

//普通冒泡排序
func bubbleSort1(arr []int) []int {
	if len(arr) <= 0 {
		panic("数组长度小于0")
	}
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	for i := 0; i < len(newArr); i++ {
		for j := 0; j < len(newArr)-i-1; j++ {
			if newArr[j] > newArr[j+1] {
				newArr[j], newArr[j+1] = newArr[j+1], newArr[j]
			}
		}
	}
	return newArr
}

//如果排列几轮后整个数组就已经是有序的了，就可提前结束工作
func bubbleSort2(arr []int) []int {
	if len(arr) <= 0 {
		panic("数组长度小于0")
	}
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	for i := 0; i < len(newArr); i++ {
		isSorted := true
		for j := 0; j < len(newArr)-i-1; j++ {
			if newArr[j] > newArr[j+1] {
				newArr[j], newArr[j+1] = newArr[j+1], newArr[j]
				isSorted = false
			}
		}
		//fmt.Println(newArr)
		if isSorted {
			fmt.Printf("第%d就已经有序了\n", i)
			break
		}
	}
	return newArr
}

//如果刚开始排列是，数组就已经部分有序，则有序区的长度和排序的轮数相等
func bubbleSort3(arr []int) []int {
	if len(arr) <= 0 {
		panic("数组长度小于0")
	}
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	//记录最后1次交换的位置
	lastExchangeIndex := 0
	//无序数列的边界，每次比较就只要比到这里为止
	sortBorder := len(arr) - 1
	for i := 0; i < len(newArr)-1; i++ {
		isSorted := true
		for j := 0; j < sortBorder; j++ {
			if newArr[j] > newArr[j+1] {
				newArr[j], newArr[j+1] = newArr[j+1], newArr[j]
				isSorted = false
				lastExchangeIndex = j
			}
		}
		sortBorder = lastExchangeIndex
		//fmt.Println(newArr)
		if isSorted {
			fmt.Printf("第%d就已经有序了\n", i)
			break
		}
	}
	return newArr

}
