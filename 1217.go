package leetcode

//https://leetcode-cn.com/problems/search-insert-position/submissions/
//sort.SearchInts() sort.Search(n int, f func(i int)bool)
func search(n int, f func(i int) bool) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) //avoid overflow
		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}
func searchInts(nums []int, target int) int {
	return search(len(nums), func(i int) bool { return nums[i] >= target })
}
func searchInsert(nums []int, target int) int {
	return searchInts(nums, target)
}
