package __array

import "strconv"

func EvalRPN(tokens []string) int {
	stack := &Stack[int]{items: []int{}}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if r, ok := isToken(token); ok {
			b := stack.Pop()
			a := stack.Pop()
			sum := calc(a, b, r)
			stack.Push(sum)
			continue
		}
		intToken, err := strconv.Atoi(token)
		if err != nil {
			return 0
		}
		stack.Push(intToken)
	}

	return stack.items[0]
}

func calc(a int, b int, o rune) int {
	switch o {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}

func isToken(token string) (rune, bool) {
	switch token {
	case "+":
		return rune(token[0]), true
	case "-":
		return rune(token[0]), true
	case "*":
		return rune(token[0]), true
	case "/":
		return rune(token[0]), true
	}
	return '!', false
}
