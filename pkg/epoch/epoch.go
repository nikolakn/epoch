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
	Relative() bool
	EndRelative() bool
	GetParent() Event
	SetEnd(end float64)
	GetDuration() float64
	GetEpoch() *EventStruct
}

type EventStruct struct {
	Id            int     `json:"id"`
	ParentId      int     `json:"parent_id"`
	Start         float64 `json:"start"`
	Description   string  `json:"Description"`
	Title         string  `json:"titile"`
	IsRelative    bool    `json:"relative"`
	IsEndRelative bool    `json:"end_relative"`
	Parent        Event   `json:"-"`
	Type          int     `json:"type"`
	Importance    int     `json:"importance"`
	Y             int     `json:"y"`
	GPS           gps.GPS `json:"gps"`
	Url           string  `json:"url"`
}

type EpochStruct struct {
	EventStruct
	End float64 `json:"end"`
}

func (e *EventStruct) GetStart() float64 {

	if e.IsRelative {
		sp := e.Parent.GetStart()
		return sp + e.Start
	}
	return e.Start
}

func (e *EpochStruct) GetStart() float64 {
	if e.IsRelative {
		sp := e.Parent.GetStart()
		return sp + e.Start
	}
	return e.Start
}

func (e *EventStruct) Relative() bool {
	return e.IsRelative
}

func (e *EpochStruct) Relative() bool {
	return e.IsRelative
}

func (e *EventStruct) EndRelative() bool {
	return e.IsEndRelative
}

func (e *EpochStruct) EndRelative() bool {
	return e.IsEndRelative
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

func (e *EventStruct) GetEpoch() *EventStruct {
	return e
}

func (e *EpochStruct) GetEpoch() *EventStruct {
	return &e.EventStruct
}
