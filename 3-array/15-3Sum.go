package __array

import (
	"slices"
)

func ThreeSum(nums []int) [][]int {
	slices.Sort(nums)

	result := make([][]int, 0)

	// len(nums)-2 to make sure that there is space for triplet
	for i := 0; i < len(nums)-2; i++ {
		first := nums[i]

		// Handle duplicates for the fist num
		if i > 0 && first == nums[i-1] {
			continue
		}
		for l, r := i+1, len(nums)-1; l < r; {
			second := nums[l]
			third := nums[r]

			sum := first + second + third

			if sum == 0 {
				result = append(result, []int{first, second, third})

				// Handle duplicates in the triplet(!) for the left pointer
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				//Handle duplicates in the triplet(!) for the right pointer
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return result
}
