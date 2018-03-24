package main

import (
	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"log"
)

type Device struct {
	ref         *firestore.DocumentRef
	name        string
	description string
	data        DeviceData
}

type DeviceData struct {
	hr, oxygen, red_led, ir_led map[string]int64
}

func NewDevice(doc *firestore.DocumentSnapshot) *Device {
	d := new(Device)
	d.ref = doc.Ref
	docData := doc.Data()
	d.name = docData["name"].(string)
	d.description = docData["description"].(string)
	return d
}

func (device *Device) String() string {
	return device.description
}

func (device *Device) initMaps() {
	device.data.hr = make(map[string]int64)
	device.data.oxygen = make(map[string]int64)
	device.data.red_led = make(map[string]int64)
	device.data.ir_led = make(map[string]int64)
}

func (device *Device) GetData() {
	ctx := context.Background()
	device.initMaps()
	iter := device.ref.Collection(DataCollection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		docData := doc.Data()
		timestamp := doc.Ref.ID
		if hr, ok := docData["hr"]; ok {
			device.data.hr[timestamp] = hr.(int64)
		}
		if oxygen, ok := docData["oxygen"]; ok {
			device.data.oxygen[timestamp] = oxygen.(int64)
		}
		if red_led, ok := docData["red_led"]; ok {
			device.data.red_led[timestamp] = red_led.(int64)
		}
		if ir_led, ok := docData["ir_led"]; ok {
			device.data.ir_led[timestamp] = ir_led.(int64)
		}
	}
}
