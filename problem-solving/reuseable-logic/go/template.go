package main

func CloneAndProcessSlice(input []int) []int {
	buffer := make([]int, len(input))
	for i, v := range input {
		buffer[i] = v * 2
	}
	return buffer
}

func ProcessSlice(input []int) {
	for i := range input {
		input[i] *= 2
	}
}
