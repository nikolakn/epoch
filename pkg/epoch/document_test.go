package epoch

import (
	"fmt"
	"testing"
	"time"
)

func TestDocumetn(t *testing.T) {

	dateTime := time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC)

	doc := NewDocument()
	e1 := doc.AddEpochWithDataRelativeEnd(dateTime, JDYear*930, "Adam")
	e2 := doc.AddRelativeEpochWithData(e1, JDYear*130, JDYear*912, "Seth")
	e4 := doc.AddRelativeEpochWithData(e2, JDYear*105, JDYear*905, "Enosh")
	e5 := doc.AddRelativeEpochWithData(e4, JDYear*90, JDYear*910, "Kenon")
	e6 := doc.AddRelativeEpochWithData(e5, JDYear*70, JDYear*895, "Mahalalel")
	e7 := doc.AddRelativeEpochWithData(e6, JDYear*65, JDYear*962, "Jared")
	e8 := doc.AddRelativeEpochWithData(e7, JDYear*162, JDYear*365, "Enoch")
	e9 := doc.AddRelativeEpochWithData(e8, JDYear*65, JDYear*969, "Methuselah")
	e10 := doc.AddRelativeEpochWithData(e9, JDYear*187, JDYear*777, "Lamech")
	e11 := doc.AddRelativeEpochWithData(e10, JDYear*182, JDYear*950, "Noah")

	e12 := doc.AddRelativeEventWithData(e11, JDYear*600, "potop")
	doc.SetGPS(e12, 45.45, 45.45)
	fmt.Println(doc)
	doc.Savejson("../../test_data/test.json")
}
