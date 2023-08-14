package epoch

import (
	"fmt"
	"sort"
	"time"

	jd "epoch/internal/julian"
)

type Document struct {
	Events []Event
}

func NewDocument() *Document {
	doc := &Document{
		Events: make([]Event, 0),
	}
	return doc
}

func (doc *Document) AddEvent(e Event) Event {
	doc.Events = append(doc.Events, e)
	doc.docSort()
	return e
}

func (doc *Document) AddEventWithData(date time.Time, des string) Event {
	julianDay := jd.TimeToJD(date)
	e := EventStruct{
		Start:       julianDay,
		Description: des,
	}
	doc.Events = append(doc.Events, e)
	doc.docSort()
	return e
}

// AddRelativeEventWithData returns error in case parent is nil.
//
// relative is how many days event is far from parent start

func (doc *Document) AddRelativeEventWithData(parent Event, relative float64, des string) Event {
	if parent == nil {
		return nil
	}
	e := EventStruct{
		Start:       relative,
		Description: des,
		Parent:      parent,
		isRelative:  true,
	}
	doc.Events = append(doc.Events, e)
	doc.docSort()
	return e
}

func (doc *Document) docSort() {
	sort.Slice(doc.Events, func(i, j int) bool {
		return doc.Events[i].GetStart() < doc.Events[j].GetStart()
	})
}

func (doc Document) String() string {
	text := ""
	for _, e := range doc.Events {
		time := jd.JDToTime(e.GetStart())

		date := fmt.Sprintf("%d %d %d %s\n", time.Year(), time.Month(), time.Day(), e.GetDescription())
		text += date
	}
	return text
}
