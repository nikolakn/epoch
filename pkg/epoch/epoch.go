package epoch

const (
	JDYear     = 365.25       // days
	JDHalfYear = JDYear / 2.0 // days
	JDCentury  = 36525        // days
)

type Event interface {
	GetStart() float64
	GetDescription() string
	GetTitle() string
	IsRelative() bool
	GetParent() Event
	SetEnd(end float64)
	GetDuration() float64
}

type EventStruct struct {
	Start       float64
	Description string
	Title       string
	isRelative  bool
	Parent      Event
}

type EpochStruct struct {
	EventStruct
	End float64
}

func (e EventStruct) GetStart() float64 {
	if e.isRelative {
		sp := e.Parent.GetStart()
		return sp + e.Start
	}
	return e.Start
}

func (e EpochStruct) GetStart() float64 {
	if e.isRelative {
		sp := e.Parent.GetStart()
		return sp + e.Start
	}
	return e.Start
}

func (e EventStruct) GetDescription() string {
	return e.Description
}

func (e EpochStruct) GetDescription() string {
	return e.Description
}

func (e EventStruct) GetTitle() string {
	return e.Title
}

func (e EpochStruct) GetTitle() string {
	return e.Title
}

func (e EventStruct) IsRelative() bool {
	return e.isRelative
}

func (e EpochStruct) IsRelative() bool {
	return e.isRelative
}

func (e EventStruct) GetParent() Event {
	return e.Parent
}

func (e EpochStruct) GetParent() Event {
	return e.Parent
}

func (e EventStruct) SetEnd(end float64) {
	return
}

func (e EpochStruct) SetEnd(end float64) {
	e.End = end
}

func (e EventStruct) GetDuration() float64 {
	return 0
}

func (e EpochStruct) GetDuration() float64 {
	return e.End
}
