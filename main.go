package main

import (
	"aavev3-raw-events-decoder/internal/datalab"
	"aavev3-raw-events-decoder/internal/decoder"
	"fmt"
	"os"
	"time"
)

func main() {
	accessKeyID := os.Getenv("ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("SECRET_ACCESS_KEY")

	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	stop := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	for day := start; day.Before(stop); day = day.AddDate(0, 0, 1) {
		DailyEtl(day, accessKeyID, secretAccessKey)
	}

}

func DailyEtl(day time.Time, accessKeyID, secretAccessKey string) {
	fmt.Printf("Starting Job for day %v...\n", day)

	endpoint := "minio-simple.lab.groupe-genes.fr"
	bucket := "projet-datalab-group-jprat"

	day_str := fmt.Sprint(day)[:10]

	input_path := "aavev3-raw-datasource/daily-raw-events/raw_events_snapshot_date=" + day_str + "/raw_events.json"
	output_path := "aavev3-raw-datasource/daily-decoded-events/decoded_events_snapshot_date=" + day_str + "/"

	fmt.Println("Reading raw events...")
	rawEvents, err := datalab.ReadRawEvents(endpoint, bucket, input_path, accessKeyID, secretAccessKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("   --> Found %v events to decode\n", len(rawEvents))
	// fmt.Println(rawEvents[3725])

	// eventsCodes := decoder.GetAllEventsCodes()
	// k := "BalanceTransfer(address,address,uint256,uint256)"
	// v := eventsCodes[k]
	// fmt.Println(v)
	// for j, event := range rawEvents {
	// 	if v == event.Topics[0].String() {
	// 		fmt.Println(j)
	// 		fmt.Println(strings.Split(k, "(")[0])
	// 	}
	// }

	// d := decoder.DecodeBalanceTransfer(rawEvents[3725])
	// fmt.Println(d)
	fmt.Println("Decoding raw events...")
	allDecodedEvents := decoder.DecodedEventsCollection{}
	eventsCodes := decoder.GetAllEventsCodes()
	decoder.DecodeRawEvents(rawEvents, &allDecodedEvents, eventsCodes)
	fmt.Printf("   --> Found %v active users\n", len(allDecodedEvents.ActiveUsers))
	fmt.Printf("   --> Found %v Initialized events\n", allDecodedEvents.NumberEventsNotDetected)

	fmt.Println("Generating and saving outputs...")
	datalab.SaveOutputs(endpoint, bucket, output_path, accessKeyID, secretAccessKey, allDecodedEvents)
	fmt.Println("Done!")
}
