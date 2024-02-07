package slidingwindow

import "fmt"

func AddUpSubArrays(arr []int, n int) [][]int {
	currentSum := 0
	currentSumStartIndex := 0
	subArrays := [][]int{}


	for i, v := range arr {
		currentSum += v

		for currentSum > n {
			currentSum -= arr[currentSumStartIndex]
			currentSumStartIndex++
		}
		
		if currentSum == n {
			subArrays = append(subArrays, arr[currentSumStartIndex:i+1])
		}
	}

	return subArrays
}

func DynamicallyDemo() {
	fmt.Println(AddUpSubArrays([]int{4, 1, 2, 5, 3, 1, 3, 2}, 5))
}
