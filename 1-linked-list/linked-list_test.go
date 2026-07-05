package linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBrowserHistoryLeetCodeCase(t *testing.T) {
	// Create a mostResent BrowserHistory instance with the initial homepage
	browser := Constructor("leetcode.com")

	// Define a list to capture the results of each operation
	results := []string{}

	// Visit "google.com"
	browser.Visit("google.com")
	results = append(results, browser.Cur.Url)
	fmt.Println(browser)

	// Visit "facebook.com"
	browser.Visit("facebook.com")
	results = append(results, browser.Cur.Url)
	fmt.Println(browser)

	// Visit "youtube.com"
	browser.Visit("youtube.com")
	results = append(results, browser.Cur.Url)
	fmt.Println(browser)

	// Back 1 step, should go to "facebook.com"
	results = append(results, browser.Back(1))
	fmt.Println(browser)

	// Back 1 step, should go to "google.com"
	results = append(results, browser.Back(1))
	fmt.Println(browser)

	// Forward 1 step, should go to "facebook.com"
	results = append(results, browser.Forward(1))
	fmt.Println(browser)

	// Visit "linkedin.com"
	browser.Visit("linkedin.com")
	results = append(results, browser.Cur.Url)
	fmt.Println(browser)

	// Forward 2 steps, should stay on "linkedin.com" (no forward steps available)
	results = append(results, browser.Forward(2))
	fmt.Println(browser)

	// Back 2 steps, should go to "google.com"
	results = append(results, browser.Back(2))
	fmt.Println(browser)

	// Back 7 steps, should go to "leetcode.com" (moving beyond history goes to the start)
	results = append(results, browser.Back(7))
	fmt.Println(browser)

	// Expected results
	expected := []string{
		"google.com",   // After visit "google.com"
		"facebook.com", // After visit "facebook.com"
		"youtube.com",  // After visit "youtube.com"
		"facebook.com", // After back 1 step
		"google.com",   // After back 1 step
		"facebook.com", // After forward 1 step
		"linkedin.com", // After visit "linkedin.com"
		"linkedin.com", // After forward 2 steps (stays on "linkedin.com")
		"google.com",   // After back 2 steps
		"leetcode.com", // After back 7 steps (moves beyond history)
	}

	// Check if the results match the expected values
	for i, res := range results {
		if res != expected[i] {
			t.Errorf("At index %d, expected '%s', got '%s'", i, expected[i], res)
		}
	}
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    *ListNode
		expected *ListNode
	}{
		{
			name:     "reverse 3 elements",
			input:    &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			expected: &ListNode{3, &ListNode{2, &ListNode{1, nil}}},
		},
		{
			name:     "empty list",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single element",
			input:    &ListNode{42, nil},
			expected: &ListNode{42, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseList(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("reverseList() = %v, want %v", listToSlice(got), listToSlice(tt.expected))
			}
		})
	}
}

// утилита для вывода списка в виде []int, чтобы проще смотреть в логах
func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
