package __array

import "math"

func characterReplacement(s string, k int) int {
	res := 0
	maxCharSequenceInWindow := 0
	str := []rune(s)
	count := make(map[rune]int)

	for l, r := 0, 0; r < len(str); r++ {
		rightChar := str[r]
		count[rightChar] = 1 + count[rightChar]
		maxCharSequenceInWindow = int(math.Max(float64(maxCharSequenceInWindow), float64(count[rightChar])))
		windowSize := r - l + 1

		if windowSize-maxCharSequenceInWindow > k {
			count[str[l]] -= 1
			l++
		}

		newWindowSize := r - l + 1

		res = int(math.Max(float64(res), float64(newWindowSize)))
	}

	return res
}
