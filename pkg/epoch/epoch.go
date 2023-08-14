package epoch

const (
	JDYear     = 365.25       // days
	JDHalfYear = JDYear / 2.0 // days
	JDCentury  = 36525        // days
)

type Event interface {
	GetStart() float64
	GetDescription() string
}

type EventStruct struct {
	Start       float64
	Description string
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
	return e.GetStart()
}

func (e EventStruct) GetDescription() string {
	return e.Description
}

func (e EpochStruct) GetDescription() string {
	return e.GetDescription()
}
