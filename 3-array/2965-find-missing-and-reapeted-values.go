package __array

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	nSquared := n * n
	count := make([]int, nSquared+1)

	for _, row := range grid {
		for _, num := range row {
			count[num]++
		}
	}

	repeated, missing := -1, -1
	for i := 1; i <= nSquared; i++ {
		if count[i] == 2 {
			repeated = i
		}
		if count[i] == 0 {
			missing = i
		}
	}

	return []int{repeated, missing}
}
