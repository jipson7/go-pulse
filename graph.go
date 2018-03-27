package main

import (
	"github.com/wcharczuk/go-chart"
	"image"
	"image/png"
	"os"
)

type Graph struct {
	trial *Trial
}

func (g Graph) createSeriesSlice() (seriesSlice []chart.Series) {
	for _, device := range g.trial.devices {
		for _, dataType := range DataTypes {
			dataset, exists := device.GetDataset(dataType)
			if exists {
				chartSeries := dataset.CreateChartSeries()
				seriesSlice = append(seriesSlice, chartSeries)
			}
		}
	}
	return
}

func (g Graph) createGraphImage() (img image.Image) {
	seriesSlice := g.createSeriesSlice()
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true, //enables / displays the x-axis
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true, //enables / displays the y-axis
			},
		},
		Series: seriesSlice,
	}
	collector := &chart.ImageWriter{}
	graph.Render(chart.PNG, collector)

	var err error
	img, err = collector.Image()
	catch(err)
	return
}

func (g Graph) SaveImageToFile(filename string) {
	img := g.createGraphImage()
	outfile, err := os.Create(filename)
	catch(err)
	defer outfile.Close()
	png.Encode(outfile, img)
}
