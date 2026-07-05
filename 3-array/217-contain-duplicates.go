package __array

func containsDuplicate(nums []int) bool {
	hash := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			return true
		}
		hash[nums[i]] = true
	}

	return false
}
