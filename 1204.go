package leetcode

//https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/solution/fen-ge-shu-zu-wei-lian-xu-zi-xu-lie-by-l-lbs5

//贪心+哈希
func isPossible(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	m := make(map[int]int)
	mend := make(map[int]int)
	// 不理解这个mend
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for _, v := range nums {
		if m[v] == 0 {
			continue
		}
		if mend[v-1] > 0 {
			//优先增长已有序列
			m[v]--
			mend[v-1]--
			mend[v]++
		} else if m[v+1] != 0 && m[v+2] != 0 {
			//否则若后续有三个连续值 则创建新序列
			m[v]--
			m[v+1]--
			m[v+2]--
			mend[v+2]++
		} else {
			return false
		}
	}
	return true
}
