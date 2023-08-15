package epoch

import (
	jd "epoch/internal/julian"
	"fmt"
	"time"
)

type PrintOptions struct {
	Flags    bool
	YearOnly bool
	Time     bool
	Duration bool
	GPS      bool
	Id       bool
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
	ret := "-"
	if e.Relative() {
		ret += "r"

	} else {
		ret += "a"
	}
	switch e.(type) {
	case *EpochStruct:
		ret += "_epoch"

	case *EventStruct:
		ret += "_event"

	}

	return ret
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
	return fmt.Sprintf("\t%-15s", e.GetEpoch().Title)

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

func (doc Document) String() string {
	text := ""
	for index, e := range doc.Events {
		text += doc.PrintId(index)
		text += doc.PrintFlags(e)
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
			if doc.PrintOptions.Time {
				text += fmt.Sprintf(" - %3s", "_")
			}
			text += fmt.Sprintf("%10s", "_")
			text += doc.PrintTitle(e)
		}
		text += doc.PrintGPS(e)
		text += "\n"
	}
	return text
}
