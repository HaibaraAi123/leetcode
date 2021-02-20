package leetcode

import "fmt"

func min1(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func minCostClimbingStairs(cost []int) int {
	m := make([]int, len(cost))
	m[len(cost)-1], m[len(cost)-2] = 0, 0
	for i := len(cost) - 3; i >= 0; i-- {
		m[i] = min1(m[i+1]+cost[i+1], m[i+2]+cost[i+2])
	}
	fmt.Println(m)
	return min1(m[0]+cost[0], m[1]+cost[1])
}
