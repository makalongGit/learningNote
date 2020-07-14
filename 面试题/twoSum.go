package 面试题

import "fmt"

/*
给定一个整数数列，找出其中和为特定值的那两个数。

你可以假设每个输入都只会有一种答案，同样的元素不能被重用。

示例:
给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func twoSum(nums []int, sum int) {
	tmpMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := tmpMap[nums[i]]; ok {
			fmt.Println(nums[i], tmpMap[nums[i]])
			break
		}
		k := sum - nums[i]
		tmpMap[k] = nums[i]
	}
}
