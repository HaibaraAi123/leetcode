package leetcode

//https://leetcode-cn.com/problems/non-decreasing-array/
//取 i-1 i i+1 i+2 四个元素 已知：i-1<= i、i>i+1
//则从四个数的所有大小关系中找出分别移动 i、i+1后满足条件的大小关系
//移动i  时为i+2>=i+1 && i+1>=i-1
//移动i+1时为i+2>=i
func checkPossibility(nums []int) bool {
	for i, count := 0, 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] && count == 1 {
			return false
		} else if nums[i] > nums[i+1] {
			if i == 0 || i+1 == len(nums)-1 {
				count++
			} else if nums[i+2] >= nums[i] {
				count++
			} else if nums[i+2] >= nums[i+1] && nums[i+1] >= nums[i-1] {
				count++
			} else {
				return false
			}
		}
	}
	return true
}
