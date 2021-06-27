package leetcode

import "fmt"

//https://leetcode-cn.com/problems/largest-number/
func largestNumber(nums []int) string {
	if len(nums) == 1 {
		return fmt.Sprint(nums[0])
	} else if len(nums) == 0 {
		return ""
	}
	return ""
}
