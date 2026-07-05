package __array

func Trap(height []int) int {
	res := 0

	//for l, r := 0, 1; r < len(height)-1; r++ {
	//	area := maxValue(height[l], height[r]) -
	//	res += area
	//}
	return res
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
