package __array

func diagonalSum(mat [][]int) int {
	sum := 0
	matrixLen := len(mat[0])

	for row, columnLeft, columnRight := 0, 0, matrixLen-1; row < matrixLen; row, columnLeft, columnRight = row+1, columnLeft+1, columnRight-1 {
		if columnRight == columnLeft {
			sum += mat[row][columnLeft]
			continue
		}
		sum += mat[row][columnRight]
		sum += mat[row][columnLeft]
	}

	return sum
}
