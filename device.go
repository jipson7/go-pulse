package main

import (
	"cloud.google.com/go/firestore"
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
