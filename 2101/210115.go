package leetcode

//https://leetcode-cn.com/problems/most-stones-removed-with-same-row-or-column/
//并查集
func removeStones(stones [][]int) int {

	fa, rank := map[int]int{}, map[int]int{}
	var find func(x int) int
	find = func(x int) int {
		if _, ok := fa[x]; !ok {
			fa[x], rank[x] = x, 1
		}
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) {
		fax, fay := find(x), find(y)
		if fax == fay {
			return
		}
		if rank[fax] < rank[fay] {
			fax, fay = fay, fax
		}
		rank[fax] += rank[fay]
		fa[fay] = fax
	}
	for _, v := range stones {
		merge(v[0], v[1]+10001)
	}
	result := len(stones)
	for x, fx := range fa {
		if x == fx {
			result--
		}
	}
	return result
}
