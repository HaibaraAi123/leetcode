package leetcode

//https://leetcode-cn.com/problems/find-pivot-index/
func pivotIndex(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	for lp, tmp := 0, 0; lp < len(nums); lp++ {
		if 2*tmp+nums[lp] == sum {
			return lp
		}
		tmp += nums[lp]
	}
	return -1
}
