package epoch

import (
	jd "epoch/internal/julian"
	"fmt"
	"strconv"
	"time"
)

type PrintOptions struct {
	Flags       bool `json:"flags"`
	YearOnly    bool `json:"yearonly"`
	Time        bool `json:"time"`
	Duration    bool `json:"duration"`
	GPS         bool `json:"gps"`
	Id          bool `json:"id_option"`
	Description bool `json:"description"`
}

func (doc Document) PrintId(index int) string {
	if !doc.PrintOptions.Id {
		return ""
	}
	return fmt.Sprintf("%-4d", index)

}
func (doc Document) PrintFlags(e Event) string {
	if !doc.PrintOptions.Flags {
		return ""
	}
	ret := ""
	if e.Relative() {
		ret += "r"

	} else {
		ret += "a"
	}
	if e.GetEpoch().Importance > 0 {
		s2 := strconv.Itoa(e.GetEpoch().Importance)
		ret += "_I" + s2
	}
	if e.GetEpoch().Type > 0 {
		s2 := strconv.Itoa(e.GetEpoch().Type)
		ret += "_T" + s2
	}
	switch e.(type) {
	case *EpochStruct:
		ret += "_epoch"

	case *EventStruct:
		ret += "_event"

	}

	return ret + " "
}

func (doc Document) PrintStart(e Event) string {
	time := jd.JDToTime(e.GetStart())
	if doc.PrintOptions.YearOnly {
		return fmt.Sprintf("%4d", time.Year())

	} else {
		if doc.PrintOptions.Time {
			res := fmt.Sprintf("%d.%d.%d %d:%d", time.Day(), time.Month(), time.Year(), time.Hour(), time.Minute())
			return fmt.Sprintf("\t%16s", res)
		} else {
			res := fmt.Sprintf("%d.%d.%d", time.Day(), time.Month(), time.Year())
			return fmt.Sprintf("\t%10s", res)
		}

	}

}
func (doc Document) PrintTitle(e Event) string {
	return fmt.Sprintf("\t%-20.15s", e.GetEpoch().Title)

}

func (doc Document) PrintDuration(duration float64) string {
	if !doc.PrintOptions.Duration {
		return ""
	}
	return fmt.Sprintf("\t( %.2f years )", duration)
}

func (doc Document) PrintEnd(end time.Time) string {
	if doc.PrintOptions.YearOnly {
		return fmt.Sprintf(" - %4d", end.Year())
	} else {
		if doc.PrintOptions.Time {
			res := fmt.Sprintf(" - %d.%d.%d %d:%d", end.Day(), end.Month(), end.Year(), end.Hour(), end.Minute())
			return fmt.Sprintf("%-16s", res)
		} else {
			res := fmt.Sprintf(" - %d.%d.%d", end.Day(), end.Month(), end.Year())
			return fmt.Sprintf("%-10s", res)
		}
	}
}
func (doc Document) PrintGPS(e Event) string {
	if !doc.PrintOptions.GPS {
		return ""
	}
	g := e.GetEpoch().GPS
	if g.Latitude != 0 || g.Longitude != 0 {
		return fmt.Sprintf("\t( %s )", g)
	}
	return ""
}
func (doc Document) PrintDescription(e Event) string {
	if !doc.PrintOptions.Description {
		return ""
	}
	d := e.GetEpoch().Description

	return fmt.Sprintf("\t%s ", d)

}

func (doc Document) String() string {
	text := ""
	for index, e := range doc.Events {
		text += doc.PrintId(index)
		text += fmt.Sprintf("%16s", doc.PrintFlags(e))
		text += doc.PrintStart(e)
		if e.GetDuration() != 0 {
			if e.Relative() || e.EndRelative() {
				time := jd.JDToTime(e.GetStart() + e.GetDuration())
				text += doc.PrintEnd(time)
				text += doc.PrintTitle(e)
				text += doc.PrintDuration(e.GetDuration() / JDYear)
			} else {
				time := jd.JDToTime(e.GetDuration())
				text += doc.PrintEnd(time)
				text += doc.PrintTitle(e)
				text += doc.PrintDuration((e.GetDuration() - e.GetStart()) / JDYear)
			}
		} else {
			if doc.PrintOptions.Time && !doc.PrintOptions.YearOnly {
				text += fmt.Sprintf(" - %3s", "_")
			}
			if doc.PrintOptions.YearOnly {
				text += fmt.Sprintf("%4s", "_")
			} else {
				text += fmt.Sprintf("%10s", "_")
			}
			text += doc.PrintTitle(e)
		}
		text += doc.PrintGPS(e)
		text += doc.PrintDescription(e)
		text += "\n"
	}
	return text
}
