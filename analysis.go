package main

import (
	"bytes"
	"errors"
	"fmt"
)

type Analysis struct {
	trial            *Trial
	similarityScores map[string]float64
}

func Analyze(trial *Trial) (a Analysis) {
	a.trial = trial
	a.similarityScores = getSimilarityScores(trial)
	return a
}

func getSimilarityScore(d1 *Dataset, d2 *Dataset) float64 {
	//TODO
	return -1.0
}

func getSimilarityScores(trial *Trial) map[string]float64 {
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
		results[dataType] = getSimilarityScore(d1, d2)
	}
	return results
}

func (a Analysis) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("ANALYSIS\n")
	buffer.WriteString("================\n")
	buffer.WriteString("\nTRIAL INFO\n")
	trialString := a.trial.String()
	buffer.WriteString(trialString)
	buffer.WriteString("\nSIMILARITY SCORES\n")
	for key, val := range a.similarityScores {
		scoresString := fmt.Sprintf("%s : %.2f\n", key, val)
		buffer.WriteString(scoresString)
	}
	return buffer.String()
}
