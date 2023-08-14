package main

import (
	"fmt"
	"time"

	epoch "epoch/pkg/epoch"
)

func main() {
	// Define the date and time you want to convert
	year := -4713
	month := time.November
	day := 24
	hour := 12
	minute := 0
	second := 0

	// Create a time.Time object for the specified date and time
	dateTime := time.Date(year, month, day, hour, minute, second, 0, time.UTC)
	dateTime2 := time.Date(-1500, 1, 31, 0, 0, 0, 0, time.UTC)

	doc := epoch.NewDocument()

	e1 := doc.AddEventWithData(dateTime2, "dogadjaj")
	doc.AddEventWithData(dateTime, "pocetak")
	doc.AddRelativeEventWithData(e1, epoch.JDYear*10, "relativni")

	fmt.Println(doc)

}
