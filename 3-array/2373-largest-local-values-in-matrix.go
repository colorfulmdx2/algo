package __array

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	res := make([][]int, n-2)

	for i := range res {
		res[i] = make([]int, n-2)
	}

	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			maxVal := grid[i][j]
			for r := i; r < i+3; r++ {
				for c := j; c < j+3; c++ {
					if grid[r][c] > maxVal {
						maxVal = grid[r][c]
					}
				}
			}
			res[i][j] = maxVal
		}
	}

	return res
}
