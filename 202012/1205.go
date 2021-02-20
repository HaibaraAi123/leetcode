package leetcode

//https://leetcode-cn.com/problems/task-scheduler/submissions/
func leastInterval(tasks []byte, n int) int {
	m := make([]int, 26)
	endm := make([]int, 26)
	time := 0
	for _, v := range tasks {
		m[v-'A']++
		endm[v-'A'] = -n - 1
	}
	for count := len(tasks); count > 0; {
		var flag bool = false
		index := 0
		exist := 0
		for i := 0; i < len(m); i++ {
			if m[i] > exist && time-endm[i] > n {
				index = i
				exist = m[i]
				flag = true
			}
		}
		if flag {
			count--
			m[index]--
			endm[index] = time
			time++
		} else {
			time++
		}
	}
	return time
}
