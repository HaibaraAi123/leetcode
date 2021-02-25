package leetcode

//https://leetcode-cn.com/problems/transpose-matrix/
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	if m != n {
		newm := make([][]int, n)
		for i := range newm {
			tmp := make([]int, m)
			for j := range tmp {
				tmp[j] = matrix[j][i]
			}
			newm[i] = tmp
		}
		return newm
	}
	for i, j := 0, 0; i < m && j < n; {
		matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		if j == n-1 {
			i, j = i+1, i+1
			continue
		}
		j++
	}
	return matrix
}
