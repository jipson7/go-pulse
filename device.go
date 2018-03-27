package main

import (
	"cloud.google.com/go/firestore"
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"log"
	"strconv"
)

type Device struct {
	ref         *firestore.DocumentRef
	name        string
	description string
	data        DeviceData
}

type DeviceData struct {
	hr, oxygen, red_led, ir_led map[int64]int64
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
	device.data.hr = make(map[int64]int64)
	device.data.oxygen = make(map[int64]int64)
	device.data.red_led = make(map[int64]int64)
	device.data.ir_led = make(map[int64]int64)
}

func (device *Device) FetchData() {
	ctx := context.Background()
	device.initMaps()
	iter := device.ref.Collection(DataCollection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		catch(err)
		docData := doc.Data()
		timestamp, err := strconv.ParseInt(doc.Ref.ID, 10, 64)
		catch(err)
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

func (device *Device) GetDataset(s string) (*Dataset, bool) {
	var x, y []int64
	var data map[int64]int64
	switch s {
	case "hr":
		data = device.data.hr
	case "oxygen":
		data = device.data.oxygen
	case "red_led":
		data = device.data.red_led
	case "ir_led":
		data = device.data.ir_led
	default:
		log.Fatalln(errors.New("Invalid Data Selection " + s))
	}
	if len(data) == 0 {
		return nil, false
	} else {
		for timestamp, val := range data {
			x = append(x, timestamp)
			y = append(y, val)
		}
	}
	return NewDataset(x, y), true
}
