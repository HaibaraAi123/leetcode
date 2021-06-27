package leetcode

//https://leetcode-cn.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) {
	for lp, rp := 0, 0; rp < n; {
		if lp >= m {
			nums1[lp] = nums2[rp]
			rp++
			lp++
			continue
		}
		if nums1[lp] > nums2[0] {
			nums1[lp], nums2[0] = nums2[0], nums1[lp]
			for i := 0; i < n-1; i++ {
				if nums2[i] > nums2[i+1] {
					nums2[i], nums2[i+1] = nums2[i+1], nums2[i]
				} else {
					break
				}
			}
		}
		lp++
	}
}
