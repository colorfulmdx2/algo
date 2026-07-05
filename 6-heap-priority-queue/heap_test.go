package __heap_priority_queue

import (
	"reflect"
	"testing"
)

func TestFindRelativeRanks(t *testing.T) {

	tests := []struct {
		score    []int
		expected []string
	}{
		{[]int{5, 4, 3, 2, 1}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
	}

	for _, test := range tests {
		result := FindRelativeRanks(test.score)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("got = %s; want %s", result, test.expected)
		}
	}
}
