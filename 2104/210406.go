package leetcode

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/
func removeDuplicates(nums []int) int {
	res, length := 0, len(nums)
	if len(nums) <= 2 {
		return len(nums)
	}
	for lp := 0; lp+2 < len(nums); {
		if nums[lp] == nums[lp+1] && nums[lp+1] == nums[lp+2] {
			res++
			nums = append(nums[0:lp+2], nums[lp+3:]...)
		} else {
			lp++
		}
	}
	return length - res
}
