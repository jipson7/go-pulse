package main

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"log"
)

type Device struct {
	ref         *firestore.DocumentRef
	name        string
	description string
}

func NewDevice(doc *firestore.DocumentSnapshot) *Device {
	d := new(Device)
	d.ref = doc.Ref
	docData := doc.Data()
	d.name = docData["name"].(string)
	d.description = docData["description"].(string)
	return d
}

func (device *Device) Print() {
	fmt.Println(device.description)
}

func (device *Device) GetData(ctx context.Context) {
	iter := device.ref.Collection(DataCollection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(doc.Ref.ID)
	}
}
