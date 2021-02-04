package leetcode

//https://leetcode-cn.com/problems/maximum-average-subarray-i/
func findMaxAverage(nums []int, k int) float64 {
	maxSum, sum := 0, 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum = sum
	for lp, rp := 1, k; rp < len(nums); {
		sum = sum + nums[rp] - nums[lp-1]
		if maxSum < sum {
			maxSum = sum
		}
		lp++
		rp++
	}
	return float64(maxSum) / float64(k)
}
