package leetcode

import "fmt"

//https://leetcode-cn.com/problems/degree-of-an-array/
func findShortestSubArray(nums []int) int {
	res, maxcnt, m := 0, 0, make(map[int][3]int)
	minfunc := func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}
	// cnt lp rp
	for i, v := range nums {
		if _, ok := m[v]; !ok {
			//m[v][0] := [3]int{1, i, i}
			//m[v][0], m[v][1] = m[v][0]+1, i
			//以上两种均错 每次都要完整给value:[3]int赋值
			m[v] = [3]int{1, i, i}
			continue
		}
		m[v] = [3]int{m[v][0] + 1, m[v][1], i}
	}
	for _, v := range m {
		if v[0] > maxcnt {
			res, maxcnt = v[2]-v[1]+1, v[0]
		} else if v[0] == maxcnt {
			res = minfunc(res, v[2]-v[1]+1)
		}
	}
	fmt.Println(m)
	return res
}
