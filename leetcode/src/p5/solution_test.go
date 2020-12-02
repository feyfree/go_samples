package p5

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	n := 10
	var dp [][]bool
	for i := 0; i < n; i++ {
		dp = append(dp, []bool{})
		for j := 0; j < n; j++ {
			dp[i] = append(dp[i], i == j)
		}
	}
	fmt.Println(dp)
}
