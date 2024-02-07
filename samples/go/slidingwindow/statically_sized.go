package slidingwindow

import "fmt"

// https://www.youtube.com/watch?v=jM2dhDPYMQM&ab_channel=QuanticDev

// Given an array of integers, find maximum
// sum subarray of the required size.

var exampleInput = []int{-1, 2, 3, 1, -3, 2}

// Sliding Window Technique, do not recalculate previous values
func MaximumSumSubArray(arr []int, size int) []int {
	currentSum := 0
	maxSum := 0
	maxSumStartIndex := 0

	for i, v := range arr {
		currentSum += v

		if i < size {
			maxSum = currentSum
		} else {
			currentSum -= arr[i-size]

			if currentSum > maxSum {
				maxSum = currentSum
				maxSumStartIndex = i - size + 1
			}
		}
	}

	return arr[maxSumStartIndex : maxSumStartIndex+size]
}

func Demo() {
	fmt.Println(MaximumSumSubArray(exampleInput, 3))
}
