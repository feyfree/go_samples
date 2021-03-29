package snippets

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestFib(t *testing.T) {
	var result []int
	result = append(result, 1)
	result = append(result, 1)
	n := rand.Intn(20)
	for i := 2; i < n; i++ {
		result = append(result, result[i-1]+result[i-2])
	}
	fmt.Println(result)
}
