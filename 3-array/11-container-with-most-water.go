package __array

func maxArea(height []int) int {
	res := 0

	for l, r := 0, len(height)-1; l < r; {
		// jst a formula of area square a * b
		// but min side coz we limited by shorter wall, (r - l) it's width of container
		area := min(height[l], height[r]) * (r - l)

		if area > res {
			res = area
		}
		// we got to find the highest side, so we're moving lowest
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return res

}
