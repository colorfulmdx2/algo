package __binary_search

func findMin(nums []int) int {

	l, r, m := 0, len(nums)-1, nums[0]

	for l <= r {
		if nums[l] < nums[r] {
			if nums[l] < m {
				m = nums[l]
			}
			break

		}

		mid := l + (r-l)/2
		m = min(m, nums[mid])

		if nums[mid] >= nums[l] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return m
}
