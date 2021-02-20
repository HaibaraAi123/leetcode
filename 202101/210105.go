package leetcode

//https://leetcode-cn.com/problems/positions-of-large-groups/
func largeGroupPositions(s string) [][]int {
	result := [][]int{}
	for i := 0; i < len(s); {
		if i == len(s)-1 {
			break
		}
		for start, end := i, i+1; end < len(s); {
			if s[end] != s[start] {
				if end-start >= 3 {
					result = append(result, []int{start, end - 1})
				}
				i = end
				break
			} else {
				if end == len(s)-1 {
					if end-start >= 2 {
						result = append(result, []int{start, end})
					}
					return result
				}
				end++
			}
		}
	}
	return result
}
