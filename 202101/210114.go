package leetcode

//https://leetcode-cn.com/problems/binary-prefix-divisible-by-5/
func prefixesDivBy5(A []int) []bool {
	var cmt int = 0
	result := make([]bool, len(A))
	for i, v := range A {
		//cmt=(cmt<<1|v)%5 我咋没看懂
		cmt = cmt<<1 + v
		cmt = cmt % 5
		result[i] = cmt == 0

	}
	return result
}
