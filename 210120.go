package leetcode

import (
	"fmt"
	"math"
	"sort"
)

//https://leetcode-cn.com/problems/maximum-product-of-three-numbers/
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	func1 := func(a ...int) int {
		tmp := math.MinInt32
		for _, v := range a {
			if v > tmp {
				tmp = v
			}
		}
		return tmp
	}
	length := len(nums)
	return func1(nums[length-1]*nums[length-2]*nums[length-3], nums[length-1]*nums[0]*nums[1], nums[0]*nums[1]*nums[2])
}
