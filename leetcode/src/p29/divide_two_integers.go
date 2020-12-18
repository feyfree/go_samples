package p29

import "math"

func divide(dividend int, divisor int) int {
	// 判断结果的标志位 true 表示为正数
	sign := (dividend < 0 && divisor < 0) || (dividend > 0 && divisor > 0)
	bigDividend := int64(math.Abs(float64(dividend)))
	bigDivisor := int64(math.Abs(float64(divisor)))
	res := 0
	for bigDividend >= bigDivisor {
		for i := 0; (bigDivisor << i) <= bigDividend; i++ {
			bigDividend -= bigDivisor << i
			res += 1 << i
		}
	}
	if !sign {
		res = -res
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}
