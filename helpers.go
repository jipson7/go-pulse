package main

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

func printCollection(iter *firestore.DocumentIterator) {
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		d := doc.Data()
		fmt.Println(d)
	}
}

func printTrials(trials []*Trial) {
	fmt.Printf("Printing %d Trials\n", len(trials))
	for _, trial := range trials {
		fmt.Println(trial.date)
	}
}
