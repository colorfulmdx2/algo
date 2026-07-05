package __array

func DailyTemperatures(temperatures []int) []int {
	stack := &Stack[int]{items: []int{}}
	res := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		curr := temperatures[i]

		for tIndex, ok := stack.Top(); ok && temperatures[tIndex] < curr; {
			stack.Pop()
			res[tIndex] = i - tIndex
			tIndex, ok = stack.Top()
		}
		stack.Push(i)
	}

	return res
}
