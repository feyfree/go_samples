package p50

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return myPowImpl(x, n)
	}
	return 1 / myPowImpl(x, n)
}

func myPowImpl(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		return myPowImpl(x*x, n/2)
	} else {
		return myPowImpl(x*x, n/2) * x
	}
}
