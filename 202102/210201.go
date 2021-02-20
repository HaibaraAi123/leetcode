package leetcode

//https://leetcode-cn.com/problems/fair-candy-swap/
func fairCandySwap(A []int, B []int) []int {
	m, sumA, sumB := make(map[int]bool), 0, 0
	for _, v := range A {
		m[v] = true
		sumA += v
	}
	for _, v := range B {
		sumB += v
	}
	for _, v := range B {
		if m[(sumA-sumB+2*v)/2] == true {
			return []int{(sumA - sumB + 2*v) / 2, v}
		}
	}

	return nil
}
