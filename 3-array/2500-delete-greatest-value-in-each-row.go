package __array

func deleteGreatestValue(grid [][]int) int {
	hLen := len(grid[0])
	sum := 0

	for i := 0; i < hLen; i++ {
		currMax := 0
		for j := 0; j < len(grid); j++ {
			row := grid[j]
			rowMax := findMax(row)
			if rowMax > currMax {
				currMax = rowMax
			}
			grid[j] = deleteElement(row, rowMax)
		}
		sum += currMax
	}

	return sum
}

func findMax(row []int) int {
	if len(row) == 0 {
		panic("row is empty")
	}

	maxValue := row[0]

	for _, num := range row {
		if num > maxValue {
			maxValue = num
		}
	}

	return maxValue
}

func deleteElement(row []int, element int) []int {
	for i, v := range row {
		if v == element {
			return append(row[:i], row[i+1:]...)
		}
	}
	return row
}
