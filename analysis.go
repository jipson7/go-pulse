package main

import (
	"bytes"
)

type Analysis struct {
	trial *Trial
}

func NewAnalysis(trial *Trial) (a Analysis) {
	a.trial = trial
	return a
}

func (a Analysis) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("ANALYSIS\n")
	buffer.WriteString("================\n")
	trialString := a.trial.String()
	buffer.WriteString(trialString)
	return buffer.String()
}
