package main

import (
	"cloud.google.com/go/firestore"
	"fmt"
)

type Trial struct {
	ref     *firestore.DocumentRef
	devices []Device
	start   int64
	end     int64
	date    string
}

type Trials []*Trial

func NewTrial(doc *firestore.DocumentSnapshot) *Trial {
	t := new(Trial)
	t.ref = doc.Ref
	t.devices = nil
	docData := doc.Data()
	t.start = docData["start"].(int64)
	t.end = docData["end"].(int64)
	t.date = docData["date"].(string)
	return t
}

func (trials Trials) LoadDevices(client *firestore.Client) {
	for _, trial := range trials {
		trial.LoadDevices(client)
	}
}

func (trial *Trial) LoadDevices(client *firestore.Client) {
	iter := trial.ref.Collection(DeviceCollection)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		device := NewDevice(doc)
		fmt.Println(device)
	}
}
