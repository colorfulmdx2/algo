package __array

func LuckyNumbers(matrix [][]int) []int {
	rLen := len(matrix)
	cLen := len(matrix[0])

	luckySlice := make([]int, 0)

	rowMins := make([]int, rLen)
	colMaxs := make([]int, cLen)

	for i := 0; i < rLen; i++ {
		rowMin := matrix[i][0]
		for j := 0; j < cLen; j++ {
			if matrix[i][j] < rowMin {
				rowMin = matrix[i][j]
			}
		}
		rowMins[i] = rowMin
	}

	for j := 0; j < cLen; j++ {
		colMax := matrix[0][j]
		for i := 0; i < rLen; i++ {
			if matrix[i][j] > colMax {
				colMax = matrix[i][j]
			}
		}
		colMaxs[j] = colMax
	}

	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if matrix[i][j] == rowMins[i] && matrix[i][j] == colMaxs[j] {
				luckySlice = append(luckySlice, matrix[i][j])
			}
		}
	}

	return luckySlice
}
