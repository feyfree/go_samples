package flow

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	val := 0

	for {
		fmt.Print("Enter number: ")
		_, _ = fmt.Scanf("%d", &val)

		switch {
		case val < 0:
			panic("You entered a negative number!")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}
