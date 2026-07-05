package __binary_search

func MinEatingSpeed(piles []int, h int) int {
	maxSpeed := 1_000_000_000
	minSpeed := 1

	for l, r := minSpeed, maxSpeed; l <= r; {
		currSpeed := l + (r-l)/2
		time := h

		for i := 0; i < len(piles); i++ {
			pile := piles[i]
			timeForCurrentPile := (pile + currSpeed - 1) / currSpeed
			time = time - timeForCurrentPile
		}

		if time < 0 {
			l = currSpeed + 1
		}
		if time >= 0 {
			minSpeed = currSpeed
			r = currSpeed - 1
		}
		time = h

	}

	return minSpeed
}

func BSearch(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2 // avoid overflow

		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
