package factory_test

import "testing"

import "../../samples/factory/simple"

func TestSimpleFactory(t *testing.T) {
	chartFactory := new(simple.ChartFactory)
	chart1 := chartFactory.GetChart("line")
	chart2 := chartFactory.GetChart("line")
	t.Log(&chart1 == &chart2)
}
