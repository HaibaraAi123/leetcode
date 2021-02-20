package leetcode

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
	size int
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) {
	h.size++
	heap.Push(h, v)
}
func (h *hp) pop() int {
	h.size--
	return heap.Pop(h).(int)
}
func (h *hp) prune() {
	for h.Len() > 0 {
		num := h.IntSlice[0]
		if h == small {
			num = -num

		}
		if d, has := delayed[num]; has {
			if d > 1 {
				delayed[num]--
			} else {
				delete(delayed, num)
			}
			heap.Pop(h)

		} else {
			break
		}
	}
}

var delayed map[int]int
var small, large *hp

//https://leetcode-cn.com/problems/sliding-window-median/
func medianSlidingWindow(nums []int, k int) []float64 {

	return []float64{0.0}
}
