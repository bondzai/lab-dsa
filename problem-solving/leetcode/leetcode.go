package main

// Input: nums = [3,2,4], target = 6
// Output: [1,2]

func twoSum(nums []int, target int) []int {
	numIndices := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, found := numIndices[complement]; found {
			return []int{j, i}
		}

		numIndices[num] = i
	}

	return nil
}
