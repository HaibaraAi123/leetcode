package leetcode

import "fmt"

//https://leetcode-cn.com/problems/regions-cut-by-slashes/
//并查集
func regionsBySlashes(grid []string) int {
	n := len(grid)
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, n)
		for j, p := 0, 0; j < n; j++ {
			if grid[i][p] == ' ' {
				nums[i][j] = 0
			} else if grid[i][j] == '/' {
				nums[i][j] = 1
			} else if grid[i][j] == '\\' {
				nums[i][j] = -1

			}
			p++
		}
	}
	fmt.Println(nums)
	return 1
}
