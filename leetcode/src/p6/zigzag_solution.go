package p6

import "strings"

func convert(s string, numRows int) string {
	// total length
	n := len(s)
	if n <= 0 || numRows == 1 {
		return s
	}

	// loop size
	loopSize := 2 * (numRows - 1)
	loops := n / loopSize
	var result strings.Builder
	for i := 0; i < numRows; i++ {
		for j := 0; j <= loops; j++ {
			// calculate index
			firstIndex := j*loopSize + i
			secondIndex := firstIndex + (numRows-i-1)*2
			if firstIndex < n {
				result.WriteByte(s[firstIndex])
			}
			// check second index whether need to append
			if i > 0 && firstIndex != secondIndex && secondIndex < n {
				result.WriteByte(s[secondIndex])
			}
		}
	}
	return result.String()
}
