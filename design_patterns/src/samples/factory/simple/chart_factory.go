package simple

type ChartFactory struct {
}

func (*ChartFactory) GetChart(m string) Chart {
	if m == "line" {
		// allocates a struct and assigns a pointer to that
		return &LineChart{}
	}
	if m == "histogram" {
		return &HistogramChart{}
	}
	return nil
}
