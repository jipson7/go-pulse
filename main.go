package main

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"log"
)

const TrialCollection = "trials"
const DeviceCollection = "devices"

func getFirestoreClient(ctx context.Context) *firestore.Client {
	conf := &firebase.Config{ProjectID: "pulseoximeterapp"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func createTrialsSlice(iter *firestore.DocumentIterator) Trials {
	var trials Trials
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		trials.AddTrial(doc)
	}
	return trials
}

func main() {
	ctx := context.Background()
	client := getFirestoreClient(ctx)
	defer client.Close()
	iter := client.Collection(TrialCollection).Documents(ctx)
	trials := createTrialsSlice(iter)
	trials.LoadDevices(ctx)
	trials.Print()
}
