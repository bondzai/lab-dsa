package main

import (
	"math"
	"sort"
	"strconv"
)

func findDuplicates(input []int) []int {
	found := make(map[int]bool)
	duplicates := []int{}

	for _, item := range input {
		if found[item] {
			duplicates = append(duplicates, item)
		} else {
			found[item] = true
		}
	}

	return duplicates
}

func findDuplicatesWithCount(input []int) (map[int]int, []int) {
	counts := make(map[int]int)
	duplicates := []int{}

	for _, item := range input {
		counts[item]++
		if counts[item] == 2 {
			duplicates = append(duplicates, item)
		}
	}

	return counts, duplicates
}

func findItemsWithCount(counts map[int]int, targetCount int) []int {
	items := []int{}
	for item, count := range counts {
		if count == targetCount {
			items = append(items, item)
		}
	}

	sort.Ints(items)

	return items
}

func reverseString(input string) (result string) {
	for _, v := range input {
		result = string(v) + result
	}
	return result
}

func isPalindrome(input int) bool {
	real := strconv.Itoa(input)
	mirror := reverseString(real)
	return real == mirror
}

func reverseInt(x int) int {
	reversed := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if reversed > math.MaxInt32/10 || (reversed == math.MaxInt32/10 && pop > 7) {
			return 0 // Return 0 if reversing x causes overflow
		}
		if reversed < math.MinInt32/10 || (reversed == math.MinInt32/10 && pop < -8) {
			return 0 // Return 0 if reversing x causes underflow
		}
		reversed = reversed*10 + pop
	}
	return reversed
}
