package main

import (
	"cloud.google.com/go/firestore"
	"errors"
	"firebase.google.com/go"
	"fmt"
	"github.com/wcharczuk/go-chart"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"image"
	"image/png"
	"log"
	"os"
)

const TrialCollection = "trials"
const DeviceCollection = "devices"
const DataCollection = "data"

func catch(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getFirestoreClient() *firestore.Client {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "pulseoximeterapp"}
	app, err := firebase.NewApp(ctx, conf)
	catch(err)

	client, err := app.Firestore(ctx)
	catch(err)
	return client
}

func createTrialsSlice(client *firestore.Client) Trials {
	ctx := context.Background()
	iter := client.Collection(TrialCollection).Documents(ctx)
	var trials Trials
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		catch(err)
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
	catch(err)
	if len(trials) <= result {
		err := errors.New("Selected Trial out of range")
		log.Fatalln(err)
	}
	return trials[result]
}

func createGraphImage(trial *Trial) (img image.Image) {
	for _, device := range trial.devices {
		dataset := device.GetDataset("hr")
		dataset.DropFirst(1)
		graph := chart.Chart{
			XAxis: chart.XAxis{
				Style: chart.Style{
					Show: true, //enables / displays the x-axis
				},
			},
			YAxis: chart.YAxis{
				Style: chart.Style{
					Show: true, //enables / displays the y-axis
				},
			},
			Series: []chart.Series{
				dataset.CreateChartSeries(),
			},
		}
		collector := &chart.ImageWriter{}
		graph.Render(chart.PNG, collector)

		var err error
		img, err = collector.Image()
		catch(err)
	}
	return
}

func saveImage(img image.Image, filename string) {
	outfile, err := os.Create(filename)
	catch(err)
	defer outfile.Close()
	png.Encode(outfile, img)
}

func main() {
	client := getFirestoreClient()
	defer client.Close()
	trials := createTrialsSlice(client)
	trial := promptForTrial(trials)
	trial.FetchAllData()
	img := createGraphImage(trial)
	saveImage(img, "./graphs/test.png")
}
