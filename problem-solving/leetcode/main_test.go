package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "No Duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{},
		},
		{
			name:     "One Duplicate",
			input:    []int{1, 2, 3, 4, 2, 5},
			expected: []int{2},
		},
		{
			name:     "Multiple Duplicates",
			input:    []int{1, 2, 3, 4, 2, 5, 6, 4, 7, 7},
			expected: []int{2, 4, 7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			duplicates := findDuplicates(tc.input)
			if len(duplicates) != len(tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, duplicates)
			} else {
				for i := 0; i < len(tc.expected); i++ {
					if duplicates[i] != tc.expected[i] {
						t.Errorf("Expected %v, but got %v", tc.expected, duplicates)
						break
					}
				}
			}
		})
	}
}

func TestFindDuplicatesWithCount(t *testing.T) {
	testCases := []struct {
		name               string
		input              []int
		expectedDuplicates []int
		expectedCounts     map[int]int
	}{
		{
			name:               "No Duplicates",
			input:              []int{1, 2, 3, 4, 5},
			expectedDuplicates: []int{},
			expectedCounts:     map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1},
		},
		{
			name:               "One Duplicate",
			input:              []int{1, 2, 3, 4, 2, 5},
			expectedDuplicates: []int{2},
			expectedCounts:     map[int]int{1: 1, 2: 2, 3: 1, 4: 1, 5: 1},
		},
		{
			name:               "Multiple Duplicates",
			input:              []int{1, 2, 3, 4, 2, 5, 6, 4, 7, 7},
			expectedDuplicates: []int{2, 4, 7},
			expectedCounts:     map[int]int{1: 1, 2: 2, 3: 1, 4: 2, 5: 1, 6: 1, 7: 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			counts, duplicates := findDuplicatesWithCount(tc.input)

			if !reflect.DeepEqual(duplicates, tc.expectedDuplicates) {
				t.Errorf("Expected duplicates %v, but got %v", tc.expectedDuplicates, duplicates)
			}

			if !reflect.DeepEqual(counts, tc.expectedCounts) {
				t.Errorf("Expected counts %v, but got %v", tc.expectedCounts, counts)
			}
		})
	}
}

func TestFindItemsWithCount(t *testing.T) {
	counts := map[int]int{1: 1, 2: 2, 3: 1, 4: 3, 5: 1, 6: 1, 7: 2}

	testCases := []struct {
		name          string
		counts        map[int]int
		targetCount   int
		expectedItems []int
	}{
		{
			name:          "Target Count 1",
			counts:        counts,
			targetCount:   1,
			expectedItems: []int{1, 3, 5, 6},
		},
		{
			name:          "Target Count 2",
			counts:        counts,
			targetCount:   2,
			expectedItems: []int{2, 7},
		},
		{
			name:          "Target Count 3",
			counts:        counts,
			targetCount:   3,
			expectedItems: []int{4},
		},
		{
			name:          "Target Count 4 (Not Present)",
			counts:        counts,
			targetCount:   4,
			expectedItems: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			items := findItemsWithCount(tc.counts, tc.targetCount)

			if !reflect.DeepEqual(items, tc.expectedItems) {
				t.Errorf("Expected items with count %d: %v, but got %v", tc.targetCount, tc.expectedItems, items)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := twoSum(tc.nums, tc.target)
			if !compareSlices(result, tc.expected) {
				t.Errorf("For nums=%v, target=%d, expected %v, but got %v", tc.nums, tc.target, tc.expected, result)
			}
		})
	}
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{121, true},     // Palindrome
		{12321, true},   // Palindrome
		{1234321, true}, // Palindrome
		{123, false},    // Not a palindrome
		{12345, false},  // Not a palindrome
	}

	for _, tc := range testCases {
		t.Run(
			strconv.Itoa(tc.input), // Use the input as the subtest name
			func(t *testing.T) {
				result := isPalindrome(tc.input)
				if result != tc.expected {
					t.Errorf("Expected isPalindrome(%d) to be %v, but got %v", tc.input, tc.expected, result)
				}
			},
		)
	}
}

func TestReverseInt(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{1534236469, 0},  // Overflow case
		{-2147483648, 0}, // Underflow case
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := reverseInt(tc.input)
			if result != tc.expected {
				t.Errorf("reverseInt(%d) = %d; want %d", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAddTwoNumbers(t *testing.T) {
	// Test case 1
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	expectedResult := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}
	result := addTwoNumbers(l1, l2)
	if !isSameList(result, expectedResult) {
		t.Errorf("Test case 1 failed. Expected: %v, Got: %v", expectedResult, result)
	}

	// Test case 2
	l1 = &ListNode{Val: 0}
	l2 = &ListNode{Val: 0}
	expectedResult = &ListNode{Val: 0}
	result = addTwoNumbers(l1, l2)
	if !isSameList(result, expectedResult) {
		t.Errorf("Test case 2 failed. Expected: %v, Got: %v", expectedResult, result)
	}
}

func isSameList(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	// If one of the lists is not nil, they are not the same.
	return l1 == nil && l2 == nil
}
