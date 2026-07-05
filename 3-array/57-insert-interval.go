package __array

import "sort"

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)

	for i, interval := range intervals {
		if newInterval[1] < interval[0] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		}

		if newInterval[0] > interval[1] {
			res = append(res, interval)
			continue
		}

		newInterval = []int{min(interval[0], newInterval[0]), max(interval[1], newInterval[1])}
	}

	res = append(res, newInterval)

	return res
}

/*
Дан диапазон портов min и max. Также есть список уже занятых портов busy.
Нужно написать функцию, возвращающую все диапазоны свободных портов внутри min и max.

### Пример

min = 30000
max = 32000
busy = [30100, 30200]

expected = [[30000, 30099], [30101, 30199], [30201, 32000]]
по факту нужно создать интервалы
*/

func GetFreePortsIntervals(min int, max int, busy []int) [][]int {
	res := make([][]int, 0)

	sort.Ints(busy)

	curr := min
	for i := 0; i < len(busy); i++ {
		if busy[i] < curr {
			continue
		}
		if busy[i] > max {
			break
		}
		if busy[i] > curr {
			res = append(res, []int{curr, busy[i] - 1})
		}
		curr = busy[i] + 1

	}

	if curr <= max {
		res = append(res, []int{curr, max})
	}

	return res
}
