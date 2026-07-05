package __binary_search

func SearchMatrix(matrix [][]int, target int) bool {
	yLen := len(matrix)
	xLen := len(matrix[1])
	fulLen := (xLen * yLen) - 1

	low, high := 0, fulLen

	for low <= high {
		mid := low + (high-low)/2
		firstIndex := mid / xLen
		secondIndex := mid - (firstIndex * xLen)

		if matrix[firstIndex][secondIndex] < target {
			low = mid + 1
		} else if matrix[firstIndex][secondIndex] > target {
			high = mid - 1
		} else {
			return true
		}
	}

	return false

}
