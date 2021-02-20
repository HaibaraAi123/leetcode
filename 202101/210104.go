package leetcode

//https://leetcode-cn.com/problems/fibonacci-number/submissions/
func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	x, y, result := 0, 1, 1
	for i := 0; i <= n-2; i++ {
		result = x + y
		x = y
		y = result
	}
	return result
}
