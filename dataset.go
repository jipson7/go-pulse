package main

import (
	"github.com/wcharczuk/go-chart"
	"sort"
)

type Dataset struct {
	x []int64
	y []int64
}

func NewDataset(x []int64, y []int64) *Dataset {
	d := new(Dataset)
	d.x = x
	d.y = y
	sort.Sort(d)
	return d
}

func (d *Dataset) GetStartTime() int64 {
	return d.x[0]
}

// Create chart.Series compatible with
// wcharczuk/go-chart
func (d *Dataset) CreateChartSeries() chart.ContinuousSeries {
	d.DropFirst(1)
	return chart.ContinuousSeries{
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
		},
		XValues: convertSliceIntToFloat(d.x),
		YValues: convertSliceIntToFloat(d.y),
	}
}

// Drops the first n elements from the dataset
func (d *Dataset) DropFirst(n int) {
	if n > d.Len() {
		n = d.Len()
	}
	d.x = d.x[n:]
	d.y = d.y[n:]
}

func convertSliceIntToFloat(l []int64) (r []float64) {
	for _, num := range l {
		r = append(r, float64(num))
	}
	return
}

func (d *Dataset) Len() int {
	return len(d.x)
}
func (d *Dataset) Swap(i, j int) {
	d.x[i], d.x[j] = d.x[j], d.x[i]
	d.y[i], d.y[j] = d.y[j], d.y[i]
}
func (d *Dataset) Less(i, j int) bool {
	return d.x[i] < d.x[j]
}
