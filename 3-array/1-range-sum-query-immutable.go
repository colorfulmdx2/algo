package __array

type NumArray struct {
	prefixArray []int
}

func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums)+1)

	for i, e := range nums {
		arr[i+1] = arr[i] + e
	}

	return NumArray{prefixArray: arr}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixArray[right+1] - this.prefixArray[left]
}
