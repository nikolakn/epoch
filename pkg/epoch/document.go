package epoch

import (
	"fmt"
	"sort"
	"time"

	"epoch/internal/gps"
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
			isEndRelative: true,
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
	e := &EpochStruct{
		EventStruct{
			Start:         relative,
			Description:   "",
			Title:         title,
			isRelative:    true,
			Parent:        parent,
			isEndRelative: true,
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
	e.GetEpoch().isEndRelative = true
	return e
}

func (doc *Document) SetGPS(e Event, l1, l2 gps.Degrees) Event {
	if e == nil {
		return nil
	}
	e.SetGPS(gps.NewGPS(l1, l2))
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
			text += fmt.Sprintf("--r %12s\t%d \t%12s", e.GetParent().GetTitle(), time.Year(), e.GetTitle()) // time.Month(), time.Day(), time.Hour(), time.Minute(),

		} else {
			text += fmt.Sprintf("%25s %d\t%12s", "", time.Year(), e.GetTitle()) //time.Month(), time.Day(), time.Hour(), time.Minute(),

		}
		if e.GetDuration() > 0 {
			if e.IsEndRelative() {
				text += fmt.Sprintf("\tduration: %6.0f godina", e.GetDuration()/JDYear)

				time := jd.JDToTime(e.GetStart() + e.GetDuration())
				text += fmt.Sprintf("\t| end %d", time.Year()) //, time.Month(), time.Day(), time.Hour(), time.Minute())

			} else {
				time := jd.JDToTime(e.GetDuration())
				text += fmt.Sprintf("\t| end %d", time.Year()) //, time.Month(), time.Day(), time.Hour(), time.Minute()
			}
		}
		if e.GetGPS().Latitude != 0 || e.GetGPS().Longitude != 0 {
			text += fmt.Sprintf("%s", e.GetGPS())
		}
		text += "\n"
	}
	return text
}
