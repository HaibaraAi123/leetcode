package leetcode

func findBoud(nums []int, target int, flag bool) int {
	lp, rp := 0, len(nums)
	for lp < rp {
		mid := lp + (rp-lp)/2
		if nums[mid] == target {
			if flag {
				//Right bound
				lp = mid + 1
			} else {
				//left
				rp = mid
			}
		} else if nums[mid] < target {
			lp = mid + 1
		} else if nums[mid] > target {
			rp = mid
		}
	}
	if flag {
		if lp == 0 {
			return -1
		}
		if nums[lp-1] != target {
			return -1
		}
		return lp - 1
	} else {
		if lp == len(nums) {
			return -1
		}
		if nums[lp] != target {
			return -1
		}
		return lp
	}
}
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	} else if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	lp := findBoud(nums, target, false)
	if lp == -1 {
		return []int{-1, -1}
	}
	rp := findBoud(nums, target, true)
	return []int{lp, rp}

}
