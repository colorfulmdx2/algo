package __array

func LongestConsecutive(nums []int) int {
	hash := make(map[int]bool)

	for _, num := range nums {
		hash[num] = true
	}

	res := 0

	for _, num := range nums {
		if hash[num-1] {
			continue
		}
		sequence := 1
		nextVal := num + 1

		for hash[nextVal] {
			sequence++
			nextVal++
		}

		if sequence > res {
			res = sequence
		}
	}

	return res
}
