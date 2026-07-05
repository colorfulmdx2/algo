package __array

import "sort"

//Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	res = append(res, intervals[0]) // to avoid edge case

	for i := 1; i < len(intervals); i++ {
		lastEnd := res[len(res)-1][1]
		start, end := intervals[i][0], intervals[i][1]

		if start <= lastEnd {
			res[len(res)-1][1] = max(lastEnd, end)
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}
