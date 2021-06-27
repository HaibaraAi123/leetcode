package leetcode

//https://leetcode-cn.com/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	res = append(res, nums)
	m, cnt := make(map[int]int), 0
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			cnt++
		}
		m[v]++
	}
	for i := 1; i < len(nums); i++ {

	}

	return res
}
