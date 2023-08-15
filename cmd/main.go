package main

import (
	"fmt"
	"time"

	epoch "epoch/pkg/epoch"
)

func main() {

	dateTime := time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC)

	//dateTime2 := time.Date(-1500, 1, 31, 0, 0, 0, 0, time.UTC)

	doc := epoch.NewDocument()
	e1 := doc.AddEpochWithDataRelativeEnd(dateTime, epoch.JDYear*930, "Adam")
	e2 := doc.AddRelativeEpochWithData(e1, epoch.JDYear*130, epoch.JDYear*912, "Seth")
	e4 := doc.AddRelativeEpochWithData(e2, epoch.JDYear*105, epoch.JDYear*905, "Enosh")
	e5 := doc.AddRelativeEpochWithData(e4, epoch.JDYear*90, epoch.JDYear*910, "Kenon")
	e6 := doc.AddRelativeEpochWithData(e5, epoch.JDYear*70, epoch.JDYear*895, "Mahalalel")
	e7 := doc.AddRelativeEpochWithData(e6, epoch.JDYear*65, epoch.JDYear*962, "Jared")
	e8 := doc.AddRelativeEpochWithData(e7, epoch.JDYear*162, epoch.JDYear*365, "Enoch")
	e9 := doc.AddRelativeEpochWithData(e8, epoch.JDYear*65, epoch.JDYear*969, "Methuselah")
	e10 := doc.AddRelativeEpochWithData(e9, epoch.JDYear*187, epoch.JDYear*777, "Lamech")
	e11 := doc.AddRelativeEpochWithData(e10, epoch.JDYear*182, epoch.JDYear*950, "Noah")

	e12 := doc.AddRelativeEventWithData(e11, epoch.JDYear*600, "potop")
	doc.SetGPS(e12, 45.45, 45.45)
	fmt.Println(doc)

}
