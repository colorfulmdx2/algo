package __array

func countNegatives(grid [][]int) int {
	count := 0
	mLen := len(grid)
	rLen := len(grid[0])

	for i := 0; i < mLen; i++ {
		for j := 0; j < rLen; j++ {
			if grid[i][j] < 0 {
				count++
			}
		}
	}

	return count
}
