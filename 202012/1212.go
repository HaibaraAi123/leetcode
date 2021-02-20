package leetcode

//https://leetcode-cn.com/problems/wiggle-subsequence/'
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	result := 0
	for lp, rp := 0, 1; rp < len(nums)-1; {

		if nums[rp]-nums[rp-1] == 0 {
			result = max(result, rp-lp-1)
			rp++
			lp = rp
		}
		cnt := nums[rp] > nums[rp-1]
		if (nums[rp+1] > nums[rp] && !cnt) || (nums[rp+1] < nums[rp] && cnt) {
			cnt = nums[rp+1] > nums[rp]
			rp++
			if rp == len(nums)-1 {
				result = max(result, rp-lp+1)
			}

		} else {
			result = max(result, rp-lp+1)
			rp++
			lp = rp
		}
	}
	return result
}
