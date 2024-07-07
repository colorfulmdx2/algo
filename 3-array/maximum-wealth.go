package __array

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	for i := 0; i < len(accounts); i++ {
		currWealth := 0
		for j := 0; j < len(accounts[i]); j++ {
			currWealth += accounts[i][j]
		}
		if currWealth > maxWealth {
			maxWealth = currWealth
		}
	}
	return maxWealth
}
