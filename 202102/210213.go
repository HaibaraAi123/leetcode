package leetcode

//https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, v := range nums {
		nums[v%(len(nums)+1)-1] += len(nums) + 1
	}
	for i, v := range nums {
		if v <= len(nums) {
			res = append(res, i+1)
		}
	}
	return res
}
