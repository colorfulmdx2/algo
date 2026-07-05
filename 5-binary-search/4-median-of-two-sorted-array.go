package __binary_search

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)

	if l%2 == 1 {
		i := l / 2
		if len(nums1) >= i {
			return float64(nums1[i])
		}
		return float64(nums2[i-len(nums1)])
	}
	return 0
}
