package __array

import (
	"slices"
	"strconv"
	"strings"
)

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Count, s2Count := [26]int{}, [26]int{}
	for i, _ := range s1 {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	matches := 0
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches += 1
		} else {
			matches += 0
		}
	}

	for r, l := len(s1), 0; r < len(s2); r++ {
		// Case if s1 == s2 or s2 starts from s1 directly (no need to slide window)
		if matches == 26 {
			return true
		}

		rPointerIndex := s2[r] - 'a'
		s2Count[rPointerIndex]++
		if s1Count[rPointerIndex] == s2Count[rPointerIndex] {
			matches++
		} else if s1Count[rPointerIndex]+1 == s2Count[rPointerIndex] {
			matches--
		}

		rPointerIndex = s2[l] - 'a'
		s2Count[rPointerIndex]--
		if s1Count[rPointerIndex] == s2Count[rPointerIndex] {
			matches++
		} else if s1Count[rPointerIndex]-1 == s2Count[rPointerIndex] {
			matches--
		}
		l++
	}
	return matches == 26
}

func runesToString(runes []rune) string {
	slices.Sort(runes)
	var builder strings.Builder
	for _, r := range runes {
		builder.WriteString(strconv.Itoa(int(r)))
	}
	return builder.String()
}
