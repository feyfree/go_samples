package map_test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	for _, age := range studentsAge {
		fmt.Printf("Ages %d\n", age)
	}
}
