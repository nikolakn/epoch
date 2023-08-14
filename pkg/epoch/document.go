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

func (doc *Document) AddEventWithData(date time.Time, title string) Event {
	julianDay := jd.TimeToJD(date)
	e := EventStruct{
		Start: julianDay,
		Title: title,
	}
	doc.Events = append(doc.Events, e)
	doc.docSort()
	return e
}

func (doc *Document) AddEpochWithData(date time.Time, end float64, title string) Event {
	julianDay := jd.TimeToJD(date)
	e := EpochStruct{
		EventStruct{
			Start: julianDay,
			Title: title,
		},
		end,
	}
	doc.Events = append(doc.Events, e)
	doc.docSort()
	return e
}

// AddRelativeEventWithData returns error in case parent is nil.
//
// relative is how many days event is far from parent start

func (doc *Document) AddRelativeEventWithData(parent Event, relative float64, title string) Event {
	if parent == nil {
		return nil
	}
	e := EventStruct{
		Start:      relative,
		Title:      title,
		Parent:     parent,
		isRelative: true,
	}
	doc.Events = append(doc.Events, e)
	doc.docSort()
	return e
}

func (doc *Document) AddRelativeEpochWithData(parent Event, relative float64, end float64, title string) Event {
	if parent == nil {
		return nil
	}
	e := EpochStruct{
		EventStruct{
			Start:       relative,
			Description: "",
			Title:       title,
			isRelative:  true,
			Parent:      parent,
		},
		end,
	}
	doc.Events = append(doc.Events, e)
	doc.docSort()
	return e
}

func (doc *Document) SetEndDate(e Event, end float64) Event {
	if e == nil {
		return nil
	}
	e.SetEnd(end)
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
		if e.IsRelative() {
			text += fmt.Sprintf("--r %12s\t%d-%d-%d %d:%d\t%12s", e.GetParent().GetTitle(), time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), e.GetTitle())

		} else {
			text += fmt.Sprintf("%d-%d-%d %d:%d\t%15s", time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute(), e.GetTitle())

		}
		if e.GetDuration() > 0 {
			text += fmt.Sprintf("\tduration: %6.0f godina", e.GetDuration()/JDYear)
		}
		text += "\n"
	}
	return text
}
