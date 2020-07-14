package BST

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBST_Remove(t *testing.T) {
	bst := New()

	n := 10
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		bst.Add(rand.Intn(10))
	}

	// 注意, 由于随机生成的数据有重复, 所以bst中的数据数量大概率是小于n的

	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}
	// 打乱切片
	for i := range nums {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
	fmt.Println(nums)
	fmt.Println(bst)

	// 乱序删除[0...n)范围里的所有元素
	for i := 0; i < n; i++ {
		if bst.Contains(nums[i]) {
			bst.Remove(nums[i])
			fmt.Println("After remove", nums[i], ", size = ", bst.GetSize())
		}
	}
	// 最终整个二分搜索树应该为空
	fmt.Println(bst.GetSize())

}
