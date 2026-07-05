package __array

func isValid(s string) bool {
	stack := make([]byte, 0)
	hash := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, v := range []byte(s) {
		if val, ok := hash[v]; ok {
			if len(stack) == 0 {
				return false
			}
			head := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if head != val {
				return false
			}
		} else {
			stack = append(stack, v)
		}

	}

	return len(stack) == 0
}
