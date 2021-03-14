package leetcode

import "container/list"

//https://leetcode-cn.com/problems/design-hashmap/
type MyHashMap struct {
	Hash []list.List
}

/** Initialize your data structure here. */
func Constructor0314() MyHashMap {
	hash := make([]list.List, BASE)
	return MyHashMap{Hash: hash}
}

/** value will always be non-negative. */
func (h *MyHashMap) Put(key int, value int) {
	v := []int{key, value}
	for lp := h.Hash[key%BASE].Front(); lp != nil; lp = lp.Next() {
		if lp.Value.([]int)[0] == key {
			lp.Value = v
			return
		}
	}
	h.Hash[key%BASE].PushFront(v)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (h *MyHashMap) Get(key int) int {
	for lp := h.Hash[key%BASE].Front(); lp != nil; lp = lp.Next() {
		if lp.Value.([]int)[0] == key {
			return lp.Value.([]int)[1]
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (h *MyHashMap) Remove(key int) {
	for lp := h.Hash[key%BASE].Front(); lp != nil; lp = lp.Next() {
		if lp.Value.([]int)[0] == key {
			h.Hash[key%BASE].Remove(lp)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
