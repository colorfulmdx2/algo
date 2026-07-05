package __array

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	codec := &Codec{}

	tests := []struct {
		encodedString string
		expected      []string
	}{
		{"4|word5|hello", []string{"word", "hello"}},
		{"3|foo3|bar3|baz", []string{"foo", "bar", "baz"}},
		{"0|3|one", []string{"", "one"}},
		{"1|a2|bc3|def", []string{"a", "bc", "def"}},
	}

	for _, test := range tests {
		result := codec.Decode1(test.encodedString)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Decode(%q) = %v; want %v", test.encodedString, result, test.expected)
		}
	}
}

func TestGetFreePortIntervals(t *testing.T) {
	tests := []struct {
		min      int
		max      int
		busy     []int
		expected [][]int
	}{
		{min: 30000, max: 32000, busy: []int{30100, 30200}, expected: [][]int{{30000, 30099}, {30101, 30199}, {30201, 32000}}},
	}

	for _, test := range tests {
		intervals := GetFreePortsIntervals(test.min, test.max, test.busy)
		if !reflect.DeepEqual(intervals, test.expected) {
			t.Errorf("intervals = %v; want %v", intervals, test.expected)
		}
	}
}
