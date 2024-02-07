package slidingwindow

// import "fmt"

// WRONG
// func MaxSequenceOf1sFlipping(arr []int, k int) []int {
// 	flipped := 0
// 	start := 0
// 	end := 0

// 	maxSubArray := map[string]int{
// 		"start": 0,
// 		// "end":   0,
// 		"len": 0,
// 	}

// 	for i, v := range arr {
// 		end = i

// 		if v == 0 {
// 			if flipped < k {
// 				flipped++
// 			} else {
// 				currLen := 1 + end - start

// 				if currLen > maxSubArray["len"] {
// 					maxSubArray["len"] = currLen
// 					maxSubArray["start"] = start
// 				}

// 				for arr[start] == 1 {
// 					start++
// 				}
// 			}
// 		}
// 	}

// 	return arr[maxSubArray["start"] : maxSubArray["start"]+maxSubArray["len"]]
// }

func FlippingSwappingDemo() {
	// fmt.Println(MaxSequenceOf1sFlipping([]int{1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 0, 1}, 3))
}
