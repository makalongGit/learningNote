package main

import "fmt"

func main1() {
	x := []string{"a","b","c"}
	for v:= range x{
		fmt.Println(v)
	}
}


// 正确示例
func main() {
	x := []string{"a", "b", "c"}
	for _, v := range x {    // 使用 _ 丢弃索引
		fmt.Println(v)
	}
}
