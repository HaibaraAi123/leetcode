package leetcode

//https://leetcode-cn.com/problems/counting-bits/
func countBits(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[(i-1)/2] + 1
		}
	}
	return res
}
