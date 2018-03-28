package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
)

type Analysis struct {
	trial *Trial
	rmse  map[string]float64
	mae   map[string]float64
}

func Analyze(trial *Trial) (a Analysis) {
	a.trial = trial
	a.rmse = getErrors(trial, "rmse")
	a.mae = getErrors(trial, "mae")
	return a
}

// Get error between 2 datasets. errorType can be rmse or mae
func getErrors(trial *Trial, errorType string) map[string]float64 {
	results := make(map[string]float64)
	if len(trial.devices) != 2 {
		err := errors.New("Require 2 devices for similarity scoring.")
		catch(err)
	}
	for _, dataType := range DataTypes {
		d1, exists1 := trial.devices[0].GetDataset(dataType)
		d2, exists2 := trial.devices[1].GetDataset(dataType)
		if !exists1 || !exists2 {
			continue
		}
		start, end := d1.GetCommonBounds(d2)
		// i is a sampling point
		numSamples := 0
		result := 0.0
		for i := start; i <= end; i = i + 10 {
			x1 := d1.Interpolate(i)
			x2 := d2.Interpolate(i)
			switch errorType {
			case "rmse":
				result += math.Pow((x1 - x2), 2.0)
			case "mae":
				result += math.Abs(x1 - x2)
			}
			numSamples++
		}
		result = result / float64(numSamples)
		if errorType == "rmse" {
			result = math.Sqrt(result)
		}
		results[dataType] = result
	}
	return results
}

func (a Analysis) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("\nANALYSIS\n")
	buffer.WriteString("================\n")
	buffer.WriteString("\nTRIAL INFO\n")
	trialString := a.trial.String()
	buffer.WriteString(trialString)
	buffer.WriteString("\nROOT MEAN SQUARED ERROR\n")
	for key, val := range a.rmse {
		scoresString := fmt.Sprintf("%s : %.2f\n", key, val)
		buffer.WriteString(scoresString)
	}
	buffer.WriteString("\nMEAN ABSOLUTE ERROR\n")
	for key, val := range a.mae {
		scoresString := fmt.Sprintf("%s : %.2f\n", key, val)
		buffer.WriteString(scoresString)
	}
	return buffer.String()
}
