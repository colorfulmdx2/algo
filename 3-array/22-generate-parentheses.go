package __array

import "strings"

func GenerateParenthesis(n int) []string {
	stack := &Stack[string]{items: []string{}}
	var res []string
	var backtrack func(open, closed int)

	backtrack = func(open, closed int) {
		if open+closed == n*2 {
			res = append(res, strings.Join(stack.items, ""))
			return
		}

		if open < n {
			stack.Push("(")
			backtrack(open+1, closed)
			stack.Pop()
		}

		if closed < open {
			stack.Push(")")
			backtrack(open, closed+1)
			stack.Pop()
		}
	}

	backtrack(0, 0)

	return res
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[T]) Top() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}
