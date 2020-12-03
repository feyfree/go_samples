package p7

import "math"

//Given a 32-bit signed integer, reverse digits of an integer.
//Note:
//Assume we are dealing with an environment that could only
//store integers within the 32-bit signed integer range:
//For the purpose of this problem,
//assume that your function returns 0 when the reversed integer overflows.
func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > (math.MaxInt32/10) || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
