package __array

func Transpose(matrix [][]int) [][]int {
	rLen := len(matrix)
	cLen := len(matrix[0])
	newMatrix := make([][]int, rLen)

	for i := range newMatrix {
		newMatrix[i] = make([]int, cLen)
	}

	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			newMatrix[j][i] = matrix[i][j]
		}
	}
	return newMatrix
}
