package main

import (
	"cloud.google.com/go/firestore"
	"errors"
	"firebase.google.com/go"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"log"
)

const TrialCollection = "trials"
const DeviceCollection = "devices"
const DataCollection = "data"

func getFirestoreClient() *firestore.Client {
	ctx := context.Background()
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

func createTrialsSlice() Trials {
	ctx := context.Background()
	client := getFirestoreClient()
	defer client.Close()
	iter := client.Collection(TrialCollection).Documents(ctx)
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
	trials.LoadDevices()
	return trials
}

func promptForTrial(trials Trials) *Trial {
	fmt.Println("Select one of the following Trials:\n")
	for idx, trial := range trials {
		fmt.Printf("Press (%d) for the following Trial:\n", idx)
		fmt.Println(trial)
	}

	fmt.Printf("Enter a number: ")
	var result int
	_, err := fmt.Scanf("%d", &result)
	if err != nil {
		log.Fatalln(err)
	}
	if len(trials) <= result {
		err := errors.New("Selected Trial out of range")
		log.Fatalln(err)
	}
	return trials[result]
}

func main() {
	trials := createTrialsSlice()
	trial := promptForTrial(trials)
	fmt.Println(trial)
}
