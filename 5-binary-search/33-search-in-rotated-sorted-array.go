package __binary_search

func SearchInRotated(nums []int, target int) int {
	bSearch := func(l, r int) int {
		for l <= r {
			mid := l + (r-l)/2
			if nums[mid] == target {
				return mid
			}

			if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return -1
	}

	l, r, pivot := 0, len(nums)-1, 0

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}

	}

	pivot = l

	leftResult := bSearch(0, pivot-1)

	if leftResult == -1 {
		return bSearch(pivot, len(nums)-1)
	}

	return leftResult
}
