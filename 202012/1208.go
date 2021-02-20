package leetcode

import "math"

// https://leetcode-cn.com/problems/split-array-into-fibonacci-sequence/
// 回溯剪枝
func splitIntoFibonacci(S string) []int {
	n := len(S)
	if n < 3 {
		return []int{}
	}
	result := make([]int, 0, len(S))
	var backtrack func(index, prev, sum int) bool
	// why backtrack = func (index, prev, sum int) bool{...backtrack()...}
	// error
	backtrack = func(index, prev, sum int) bool {
		if index == n {
			return len(result) >= 3
		}
		cur := 0
		for i := index; i < n; i++ {

			if i > index && S[index] == '0' {
				break
			}
			cur = cur*10 + int(S[i]-'0')
			if cur > math.MaxInt32 {
				break
			}
			if len(result) >= 2 {
				if cur < sum {
					continue
				} else if cur > sum {
					return false
				}
			}
			result = append(result, cur)
			if backtrack(i+1, cur, cur+prev) {
				return true
			}
			// 剪枝
			result = result[:len(result)-1]
		}
		return false
	}
	backtrack(0, 0, 0)
	return result
}
