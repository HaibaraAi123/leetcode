package leetcode

//https://leetcode-cn.com/problems/toeplitz-matrix/
func isToeplitzMatrix(matrix [][]int) bool {
	row, col := len(matrix), len(matrix[0])

	for i, j := 0, 0; i < row; {
		for x, y, tmp := i, j, matrix[i][j]; x < row && y < col; {
			if matrix[x][y] != tmp {
				return false
			}
			tmp, x, y = matrix[x][y], x+1, y+1
		}
		i++
	}
	for i, j := 0, 0; j < col; {
		for x, y, tmp := i, j, matrix[i][j]; x < row && y < col; {
			if matrix[x][y] != tmp {
				return false
			}
			tmp, x, y = matrix[x][y], x+1, y+1
		}
		j++
	}
	return true
}
