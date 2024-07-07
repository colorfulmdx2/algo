package __array

func rowAndMaximumOnes(mat [][]int) []int {
	rowIndex := 0
	onesCount := 0

	mLen := len(mat)

	for i := mLen - 1; i >= 0; i-- {
		currOnesCount := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				currOnesCount++
			}
		}
		if currOnesCount >= onesCount {
			rowIndex = i
			onesCount = currOnesCount
		}
	}

	return []int{rowIndex, onesCount}
}
