package linked_list

func findDuplicate(nums []int) int {
	hash := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := hash[v]; ok {
			return v
		}
		hash[v] = struct{}{}
	}
	return -1
}
