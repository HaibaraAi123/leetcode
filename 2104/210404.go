package leetcode

//https://leetcode-cn.com/problems/rabbits-in-forest/
func numRabbits(answers []int) int {
	m, cnt := make(map[int]int), 0
	for _, v := range answers {
		m[v]++
	}
	for v, num := range m {
		var k int = 0
		if num%(v+1) == 0 {
			k = num / (v + 1)
		} else {
			k = num/(v+1) + 1
		}
		cnt += (1 + v) * k
	}
	return cnt
}
