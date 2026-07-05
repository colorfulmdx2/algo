package __strings

import (
	"reflect"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {

	tests := []struct {
		str      string
		expected string
	}{
		{"cbbd", "bab"},
	}

	for _, test := range tests {
		result := LongestPalindrome(test.str)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Polindrome from(%s) = %s; want %s", test.str, result, test.expected)
		}
	}
}
