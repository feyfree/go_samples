package simple

import "fmt"

type Chart interface {
	Display()
}

type HistogramChart struct {
}

type LineChart struct{}

func (h *HistogramChart) Display() {
	fmt.Println("This is a histogram chart")
}

func (l *LineChart) Display() {
	fmt.Println("This is a line chart")
}
