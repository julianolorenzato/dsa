package slidingwindow

func DesiredCharaters(str string, chars []rune) string {
	shortestStart := 0
	shortestEnd := 0

	start := 0

	needed := countUniques(chars)

	for i, v := range str {
		if contains(chars, v) {
			needed--

			if needed == 0 {
				for !contains(chars, rune(str[start])) {
					start++
				}
			}
		}
	}
}

func contains(chars []rune, char rune) bool {
	for _, c := range chars {
		if c == char {
			return true
		}
	}

	return false
}

func countUniques(chars []rune) int {
	counter := make(map[rune]bool)

	for _, c := range chars {
		counter[c] = true
	}

	return len(counter)
}
