package leetcode

import "container/list"

const BASE = 769

//https://leetcode-cn.com/problems/design-hashset/
type MyHashSet struct {
	Hash []list.List
}

/** Initialize your data structure here. */
func Constructor0313() MyHashSet {
	hash := make([]list.List, BASE)
	return MyHashSet{Hash: hash}
}

func (h *MyHashSet) Add(key int) {
	if !h.Contains(key) {
		h.Hash[key%BASE].PushFront(key)
	}

}

func (h *MyHashSet) Remove(key int) {
	for lp := h.Hash[key%BASE].Front(); lp != nil; lp = lp.Next() {
		if lp.Value.(int) == key {
			h.Hash[key%BASE].Remove(lp)
		}
	}
}

/** Returns true if this set contains the specified element */
func (h *MyHashSet) Contains(key int) bool {
	for lp := h.Hash[key%BASE].Front(); lp != nil; lp = lp.Next() {
		if lp.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
