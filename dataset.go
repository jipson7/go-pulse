package main

type Dataset struct {
	x []float64
	y []float64
}

func NewDataset(x []int64, y []int64) *Dataset {
	d := new(Dataset)
	d.x = convertSliceIntToFloat(x)
	d.y = convertSliceIntToFloat(y)
	return d
}

func convertSliceIntToFloat(l []int64) (r []float64) {
	for _, num := range l {
		r = append(r, float64(num))
	}
	return
}
