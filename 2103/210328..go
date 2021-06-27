package leetcode

//https://leetcode-cn.com/problems/binary-search-tree-iterator/
type TreeNode0328 struct {
	Val         int
	Left, Right *TreeNode0328
}
type BSTIterator struct {
	tree []int
	p    int
}

func Tree2slice(root *TreeNode0328) []int {
	res := []int{}
	if root.Left != nil {
		res = append(res, Tree2slice(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, Tree2slice(root.Right)...)
	}
	return res
}
func Constructor0328(root *TreeNode0328) BSTIterator {
	res := Tree2slice(root)
	return BSTIterator{tree: res, p: -1}
}

func (b *BSTIterator) Next() int {
	b.p += 1
	return b.tree[b.p]
}

func (b *BSTIterator) HasNext() bool {
	return len(b.tree) != b.p+1
}
