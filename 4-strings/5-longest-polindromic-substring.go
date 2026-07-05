package __strings

func LongestPalindrome(s string) string { // babad
	polindrom := ""
	polindromLen := 0

	for i, _ := range s {
		// odd length
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			currSubstrLen := r - l + 1
			if currSubstrLen > polindromLen {
				polindrom = s[l : r+1]
				polindromLen = r - l + 1
			}
			l--
			r++
		}

		// even length
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			currSubstrLen := r - l + 1
			if currSubstrLen > polindromLen {
				polindrom = s[l : r+1]
				polindromLen = r - l + 1
			}
			l--
			r++
		}

	}

	return polindrom
}
