package __array

import (
	"math"
)

func LengthOfLongestSubstring(s string) int {
	substringLen := 0
	runes := []rune(s)
	hash := make(map[rune]bool)

	for l, r := 0, 0; r < len(runes); {
		// case when current char is in the prev substring
		// so it's end of substring and we move left pointer
		if hash[runes[r]] {
			delete(hash, runes[l])
			l++
		} else {
			// case when current char not in the hash and substring is continue
			hash[runes[r]] = true
			substringLen = int(math.Max(float64(substringLen), float64(r-l+1)))
			r++
		}
	}

	return substringLen
}
