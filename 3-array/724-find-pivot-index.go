package __array

func pivotIndex(nums []int) int {
	sum := 0
	leftSum := 0
	for _, e := range nums {
		sum += e
	}
	for i := 0; i < len(nums); i++ {
		rightSum := sum - nums[i] - leftSum

		if rightSum == leftSum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}
