package leetcode

//https://leetcode-cn.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; ; {
		if nums[i] == val && i != len(nums)-1 {
			nums = append(nums[0:i], nums[i+1:]...)
		} else if nums[i] == val && i == len(nums)-1 {
			nums = nums[0:i]
			break
		} else {
			i++
		}
		if i == len(nums) {
			break
		}
	}
	return len(nums)
}
