package __array

import "fmt"

func TopKFrequent(nums []int, k int) (res []int) {
	countHash := map[int]int{}
	countSlice := make([][]int, len(nums)+1)

	for _, num := range nums {
		if count, ok := countHash[num]; ok {
			countHash[num] = count + 1
		} else {
			countHash[num] = 1
		}
	}

	for key, numCount := range countHash {
		countSlice[numCount] = append(countSlice[numCount], key)
	}
	fmt.Println(countHash)
	for i := len(countSlice) - 1; i > 0; i-- {
		res = append(res, countSlice[i]...)
		if len(res) == k {
			return
		}
	}

	return
}
