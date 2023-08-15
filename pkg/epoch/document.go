package epoch

import (
	"sort"
	"time"

	"epoch/internal/gps"
	jd "epoch/internal/julian"
)

type Document struct {
	Events       []Event
	printOptions PrintOptions
}

func NewDocument(po PrintOptions) *Document {
	doc := &Document{
		Events: make([]Event, 0),
	}
	doc.printOptions = po
	return doc
}

func (doc *Document) AddEvent(e Event) Event {
	doc.Events = append(doc.Events, e)
	doc.docSort()
	return e
}

func (doc *Document) AddEventWithData(date time.Time, title string) Event {
	e := &EventStruct{
		Start: jd.TimeToJD(date),
		Title: title,
	}
	doc.Events = append(doc.Events, e)
	doc.docSort()
	return e
}

func (doc *Document) AddEpochWithDataRelativeEnd(date time.Time, end float64, title string) Event {
	e := &EpochStruct{
		EventStruct{
			Start:         jd.TimeToJD(date),
			Title:         title,
			IsEndRelative: true,
		},
		end,
	}
	doc.Events = append(doc.Events, e)
	doc.docSort()
	return e
}

func (doc *Document) AddEpochWithData(date time.Time, end time.Time, title string) Event {
	e := &EpochStruct{
		EventStruct{
			Start: jd.TimeToJD(date),
			Title: title,
		},
		jd.TimeToJD(end),
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
	e := &EventStruct{
		Start:      relative,
		Title:      title,
		Parent:     parent,
		IsRelative: true,
	}
	doc.Events = append(doc.Events, e)
	doc.docSort()
	return e
}

func (doc *Document) AddRelativeEpochWithData(parent Event, relative float64, end float64, title string) Event {
	if parent == nil {
		return nil
	}
	e := &EpochStruct{
		EventStruct{
			Start:       relative,
			Description: "",
			Title:       title,
			IsRelative:  true,
			Parent:      parent,
		},
		end,
	}
	doc.Events = append(doc.Events, e)
	doc.docSort()
	return e
}

func (doc *Document) SetEndDate(e Event, end time.Time) Event {
	if e == nil {
		return nil
	}
	julianDay := jd.TimeToJD(end)
	e.SetEnd(julianDay)
	return e
}

func (doc *Document) SetEndJD(e Event, jd float64) Event {
	if e == nil {
		return nil
	}
	e.SetEnd(jd)
	return e
}

func (doc *Document) SetRelativeEndDate(e Event, jd float64) Event {
	if e == nil {
		return nil
	}
	e.SetEnd(jd)
	e.GetEpoch().IsEndRelative = true
	return e
}

func (doc *Document) SetGPS(e Event, l1, l2 gps.Degrees) Event {
	if e == nil {
		return nil
	}
	e.GetEpoch().GPS = gps.NewGPS(l1, l2)
	return e
}

func (doc *Document) docSort() {
	sort.Slice(doc.Events, func(i, j int) bool {
		return doc.Events[i].GetStart() < doc.Events[j].GetStart()
	})
}
