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
