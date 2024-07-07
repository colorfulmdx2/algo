package __array

func flipAndInvertImage(image [][]int) [][]int {
	matrixLen := len(image[0])
	for row := 0; row < matrixLen; row++ {
		for l, r := 0, matrixLen-1; l < r; l, r = l+1, r-1 {
			image[row][l], image[row][r] = image[row][r], image[row][l]
		}

		for i := 0; i < matrixLen; i++ {
			if image[row][i] == 0 {
				image[row][i] = 1
			} else {
				image[row][i] = 0
			}
		}
	}

	return image
}
