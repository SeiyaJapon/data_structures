package search_insert

func SearchInsert(nums []int, target int, low int, high int) int {
	if low == high {
		if nums[low] < target {
			return low + 1
		} else {
			return low
		}
	}

	if nums[(low+high)/2] < target {
		return SearchInsert(nums, target, (low+high)/2+1, high)
	}

	return SearchInsert(nums, target, low, (low+high)/2)
}
