package epoch

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"epoch/internal/gps"
	jd "epoch/internal/julian"
)

type Document struct {
	Events       []Event
	PrintOptions PrintOptions
	FileName     string
}

func NewDocument(po PrintOptions, fileName string) *Document {
	doc := &Document{
		Events: make([]Event, 0),
	}
	doc.PrintOptions = po
	doc.FileName = fileName
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
	for i, e := range doc.Events {
		e.GetEpoch().Id = i
	}
}

func (doc *Document) GetEventbuId(id int) Event {
	if id < 0 || id > len(doc.Events)-1 {
		fmt.Println("id out of range")
		return nil
	}
	return doc.Events[id]
}

func (doc *Document) GetEventbyTitle(title string) Event {
	for _, e := range doc.Events {
		if e.GetEpoch().Title == title {
			return e
		}
	}
	return nil
}

func (doc *Document) MoveStartAps(e Event, date time.Time) {
	start := jd.TimeToJD(date)
	e.GetEpoch().Start = start
	doc.docSort()
}

func (doc *Document) MoveStartRel(e Event, rel float64) {
	e.GetEpoch().Start = rel
	doc.docSort()
}
func (doc *Document) DeleteEvent(ev Event) {
	for _, e := range doc.Events {
		if e.GetParent() == ev {
			e.GetEpoch().Start = e.GetStart()
			e.GetEpoch().IsRelative = false //must go after setting start bavause secound getstart must be apsolute
			if e.GetDuration() != 0 {
				e.SetEnd(e.GetStart() + e.GetDuration())
			}

		}
	}
	for index, e := range doc.Events {
		if e == ev {
			l := len(doc.Events) - 1
			doc.Events[index] = doc.Events[l]
			doc.Events = doc.Events[:l]
		}
	}
	doc.docSort()

}

func (doc *Document) SearchEventsByTitle(title string) []Event {
	events := make([]Event, 0)
	for id, e := range doc.Events {
		if strings.Contains(e.GetEpoch().Title, title) {
			e.GetEpoch().Id = id
			events = append(events, e)
		}
	}
	return events
}

func (doc *Document) SearchEventsByDescription(des string) []Event {
	events := make([]Event, 0)
	for id, e := range doc.Events {
		if strings.Contains(e.GetEpoch().Description, des) {
			e.GetEpoch().Id = id
			events = append(events, e)
		}
	}
	return events
}
