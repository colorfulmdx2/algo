package __strings

import "unicode"

func isPalindrome(s string) bool {
	runes := []rune(s)
	for l, r := 0, len(runes)-1; l < r; {
		left := unicode.ToLower(runes[l])
		right := unicode.ToLower(runes[r])

		if !isValidChar(left) {
			l++
			continue
		}
		if !isValidChar(right) {
			r--
			continue
		}

		if left != right {
			return false
		}
		l++
		r--
	}
	return true
}

func isValidChar(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsDigit(char)
}
