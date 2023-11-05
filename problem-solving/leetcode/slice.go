package main

import (
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

func reverse(input string) (result string) {
	for _, v := range input {
		result = string(v) + result
	}
	return result
}

func isPalindrome(input int) bool {
	real := strconv.Itoa(input)
	mirror := reverse(real)
	return real == mirror
}
