package leetcode

//https://leetcode-cn.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		if m[v]++; m[v] != 1 {
			return true
		}
	}
	return false
}
