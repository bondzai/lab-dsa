package main

// Input: nums = [3,2,4], target = 6
// Output: [1,2]

func twoSum(nums []int, target int) []int {
	// Create a map to store the values and their indices
	numIndices := make(map[int]int)

	for i, num := range nums {
		// Calculate the complement needed to reach the target
		complement := target - num

		// Check if the complement exists in the map
		if j, found := numIndices[complement]; found {
			// Return the indices of the two numbers that add up to the target
			return []int{j, i}
		}

		// If the complement doesn't exist in the map, add the current number to the map
		numIndices[num] = i
	}

	// If no solution is found, return an empty slice or an error value
	return nil
}
