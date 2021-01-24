package leetcode

//https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/submissions/
func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	length := 1
	for lp, rp := 0, 1; rp < len(nums); {
		if nums[rp] <= nums[rp-1] {
			if rp-lp > length {
				length = rp - lp
			}
			lp, rp = rp, rp+1
		} else {
			rp++
			if rp == len(nums) {
				if rp-lp > length {
					length = rp - lp
				}
				return length
			}
		}
	}
	return length
}
