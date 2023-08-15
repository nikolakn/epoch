package epoch

import (
	"epoch/internal/gps"
)

const (
	JDYear     = 365.25       //2425     // days
	JDHalfYear = JDYear / 2.0 // days
	JDCentury  = 36525        // days
)

type Event interface {
	GetStart() float64
	GetDescription() string
	GetTitle() string
	IsRelative() bool
	IsEndRelative() bool
	GetParent() Event
	SetEnd(end float64)
	GetDuration() float64
	GetGPS() gps.GPS
	SetGPS(g gps.GPS)
	GetEpoch() *EventStruct
}

type EventStruct struct {
	Id            int64
	Start         float64
	Description   string
	Title         string
	isRelative    bool
	isEndRelative bool
	Parent        Event
	Type          int
	Importance    int
	Y             int
	GPS           gps.GPS
	url           string
}

type EpochStruct struct {
	EventStruct
	End float64
}

func (e *EventStruct) GetStart() float64 {

	if e.isRelative {
		sp := e.Parent.GetStart()
		return sp + e.Start
	}
	return e.Start
}

func (e *EpochStruct) GetStart() float64 {
	if e.isRelative {
		sp := e.Parent.GetStart()
		return sp + e.Start
	}
	return e.Start
}

func (e *EventStruct) GetDescription() string {
	return e.Description
}

func (e *EpochStruct) GetDescription() string {
	return e.Description
}

func (e *EventStruct) GetTitle() string {
	return e.Title
}

func (e *EpochStruct) GetTitle() string {
	return e.Title
}

func (e *EventStruct) IsRelative() bool {
	return e.isRelative
}

func (e *EpochStruct) IsRelative() bool {
	return e.isRelative
}

func (e *EventStruct) IsEndRelative() bool {
	return e.isEndRelative
}

func (e *EpochStruct) IsEndRelative() bool {
	return e.isEndRelative
}

func (e *EventStruct) GetParent() Event {
	return e.Parent
}

func (e *EpochStruct) GetParent() Event {
	return e.Parent
}

func (e *EventStruct) SetEnd(end float64) {
	return
}

func (e *EpochStruct) SetEnd(end float64) {
	e.End = end
}

func (e *EventStruct) GetDuration() float64 {
	return 0
}

func (e *EpochStruct) GetDuration() float64 {
	return e.End
}

func (e *EventStruct) GetGPS() gps.GPS {
	return e.GPS
}

func (e *EpochStruct) GetGPS() gps.GPS {
	return e.GPS
}

func (e *EventStruct) SetGPS(ev gps.GPS) {
	e.GPS = ev
}

func (e *EpochStruct) SetGPS(ev gps.GPS) {
	e.GPS = ev
}

func (e *EventStruct) GetEpoch() *EventStruct {
	return e
}

func (e *EpochStruct) GetEpoch() *EventStruct {
	return &e.EventStruct
}
