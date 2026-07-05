package __array

func twoSum(numbers []int, target int) []int {
	hash := make(map[int]int)

	for i, e := range numbers {
		addendum := target - e

		if index, ok := hash[addendum]; ok {
			return []int{index, i + 1}
		}
		hash[e] = i + 1
	}

	return []int{}
}

func TwoSumTwoPointers(numbers []int, target int) []int {
	for l, r := 0, len(numbers)-1; l < r; {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}
		if sum < target {
			l++
			continue
		}
		if sum > target {
			r--
			continue
		}
	}
	return []int{}
}
