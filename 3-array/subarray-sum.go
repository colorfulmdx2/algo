package __array

func subarraySum(nums []int, k int) int {
	var count int
	var prefixSum int

	hash := make(map[int]int)

	hash[0] = 1

	for _, num := range nums {
		prefixSum += num

		if value, ok := hash[prefixSum-k]; ok {
			count += value
		}

		hash[prefixSum] = hash[prefixSum] + 1
	}
	return count
}
