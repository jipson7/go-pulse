package main

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"log"
)

type Trial struct {
	ref     *firestore.DocumentRef
	devices []*Device
	start   int64
	end     int64
	date    string
}

type Trials []*Trial

func (trials *Trials) AddTrial(doc *firestore.DocumentSnapshot) {
	*trials = append(*trials, newTrial(doc))
}

func newTrial(doc *firestore.DocumentSnapshot) *Trial {
	t := new(Trial)
	t.ref = doc.Ref
	t.devices = nil
	docData := doc.Data()
	t.start = docData["start"].(int64)
	t.end = docData["end"].(int64)
	t.date = docData["date"].(string)
	return t
}

func (trials Trials) LoadDevices(ctx context.Context) {
	for _, trial := range trials {
		trial.LoadDevices(ctx)
	}
}

func (trial *Trial) LoadDevices(ctx context.Context) {
	iter := trial.ref.Collection(DeviceCollection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		device := NewDevice(doc)
		trial.devices = append(trial.devices, device)
	}
}

func (trial *Trial) Print() {
	fmt.Println(trial.date)
}

func (trials Trials) Print() {
	for _, trial := range trials {
		fmt.Println("\nTrial:")
		trial.Print()
		fmt.Println("Devices:")
		for _, device := range trial.devices {
			device.Print()
		}
	}
}
