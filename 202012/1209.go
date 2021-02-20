package leetcode

//https://leetcode-cn.com/problems/unique-paths/\

func uniquePaths(m int, n int) int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	result[m-1][n-1] = 1
	//动态规划
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+1 < m {
				result[i][j] += result[i+1][j]
			}
			if j+1 < n {
				result[i][j] += result[i][j+1]
			}
		}
	}
	return result[0][0]
}
