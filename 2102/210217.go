package leetcode

//https://leetcode-cn.com/problems/reshape-the-matrix/
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums)*len(nums[0]) != r*c {
		return nums
	}
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}
	cnt := 0
	for _, v := range nums {
		for _, v1 := range v {
			res[cnt/c][cnt%c] = v1
			cnt++
		}
	}
	return res
}
