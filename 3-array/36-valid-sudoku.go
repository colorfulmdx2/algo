package __array

const colKey = "column"
const rowKey = "row"
const gridKey = "grid"

func isValidSudoku(board [][]byte) bool {
	//hash := make(map[string]bool)
	//
	//for row := 0; row < len(board); row++ {
	//	for col := 0; col < len(board); col++ {
	//		currValue := string(board[row][col])
	//
	//		if currValue == "." {
	//			continue
	//		}
	//
	//		_, okRow := hash[currValue+rowKey+string(row)]
	//		_, okColumn := hash[currValue+colKey+string(col)]
	//		_, okGrid := hash[currValue+gridKey+string(row/3)+"-"+string(col/3)]
	//
	//		if okRow || okColumn || okGrid {
	//			return false
	//		}
	//
	//		hash[currValue+rowKey+string(row)] = true
	//		hash[currValue+colKey+string(col)] = true
	//		hash[currValue+gridKey+string(row/3)+"-"+string(col/3)] = true
	//	}
	//}
	return true
}
