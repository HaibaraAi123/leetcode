package leetcode

//https://leetcode-cn.com/problemset/all/
func makeConnected(n int, connections [][]int) int {
	fa := make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx == fy {
			return
		}
		fa[fx] = fa[fy]
		return
	}
	for _, v := range connections {
		merge(v[0], v[1])
	}
	head := 0
	for x, v := range fa {
		if v == x {
			head++
		}
	}
	if len(connections) >= n-1 {
		return head - 1
	}
	return -1

}
