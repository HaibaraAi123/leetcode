package leetcode

//https://leetcode-cn.com/problems/rotate-image/submissions/
func rotate(matrix [][]int) {
	n := len(matrix) - 1
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j <= n-2*i-1; j++ {
			matrix[i][i+j], matrix[i+j][n-i], matrix[n-i][n-j-i], matrix[n-j-i][i] = matrix[n-j-i][i], matrix[i][i+j], matrix[i+j][n-i], matrix[n-i][n-j-i]

		}
	}
}
