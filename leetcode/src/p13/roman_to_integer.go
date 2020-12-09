package p13

func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	initial := 0
	n := len(s)
	for i := 0; i < n-1; i++ {
		valueF := romanMap[rune(s[i])]
		valueB := romanMap[rune(s[i+1])]
		if valueF < valueB {
			initial -= valueF
		} else {
			initial += valueF
		}
	}
	initial += romanMap[rune(s[n-1])]
	return initial
}
