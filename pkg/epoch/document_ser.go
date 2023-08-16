package epoch

import (
	"encoding/json"
	"epoch/internal/gps"
	"io"
	"log"
	"os"
)

type MarshalEventStruct struct {
	Id            int     `json:"id"`
	ParentId      int     `json:"parent_id"`
	Start         float64 `json:"start"`
	Description   string  `json:"Description"`
	Title         string  `json:"titile"`
	IsRelative    bool    `json:"relative"`
	IsEndRelative bool    `json:"end_relative"`
	Type          int     `json:"type"`
	Importance    int     `json:"importance"`
	Y             int     `json:"y"`
	GPS           gps.GPS `json:"gps"`
	Url           string  `json:"url"`
	End           float64 `end:"start"`
}

type MarshalDoc struct {
	Events []MarshalEventStruct `json:"events"`
	Po     PrintOptions         `json:"print_options"`
}

func (doc *Document) Savejson(name string) {
	for index, e := range doc.Events {
		e.GetEpoch().Id = index
	}

	tempArr := make([]MarshalEventStruct, 0)
	for index, e := range doc.Events {
		var es MarshalEventStruct
		es.Description = e.GetEpoch().Description
		es.Start = e.GetEpoch().Start
		es.Id = index
		es.Title = e.GetEpoch().Title
		es.IsRelative = e.GetEpoch().IsRelative
		es.IsEndRelative = e.GetEpoch().IsEndRelative
		es.Type = e.GetEpoch().Type
		es.Importance = e.GetEpoch().Importance
		es.Y = e.GetEpoch().Y
		es.GPS = e.GetEpoch().GPS
		es.Url = e.GetEpoch().Url
		es.End = e.GetDuration()
		if e.GetEpoch().Parent != nil {
			es.ParentId = e.GetEpoch().Parent.GetEpoch().Id
		} else {
			es.ParentId = -1
		}

		tempArr = append(tempArr, es)

	}
	md := MarshalDoc{
		Events: tempArr,
		Po:     doc.PrintOptions,
	}

	file, err_m := json.MarshalIndent(md, "", "  ")
	if err_m != nil {
		log.Println("error marshal: ", err_m)
		return
	}
	err_write := os.WriteFile(name, file, 0644)
	if err_write != nil {
		log.Println("error write to file:  ", err_write)
		return
	}
}

func (doc *Document) LoadFromJson(name string) {
	//tempArr :=
	md := MarshalDoc{
		Events: make([]MarshalEventStruct, 0),
		Po:     doc.PrintOptions,
	}

	jsonFile, err := os.Open(name)
	if err != nil {
		log.Println("error read to file:  ", err)
		return
	}
	byteValue, _ := io.ReadAll(jsonFile)
	erru := json.Unmarshal(byteValue, &md) //&md.Events for old files
	if erru != nil {
		log.Println("error read to file:  ", erru)
		return
	}
	tempArr := md.Events
	doc.Events = make([]Event, 0)
	doc.PrintOptions = md.Po //comment out for old files
	for _, e := range tempArr {
		if e.End != 0 {
			es := &EpochStruct{}
			es.Start = e.Start
			es.Description = e.Description
			es.Id = e.Id
			es.End = e.End
			es.ParentId = e.ParentId
			es.GPS = e.GPS
			es.Url = e.Url
			es.IsRelative = e.IsRelative
			es.IsEndRelative = e.IsEndRelative
			es.Url = e.Url
			es.Title = e.Title
			es.Type = e.Type
			es.Importance = e.Importance
			doc.Events = append(doc.Events, es)
		} else {
			es := &EventStruct{}
			es.Start = e.Start
			es.Description = e.Description
			es.Id = e.Id
			es.ParentId = e.ParentId
			es.GPS = e.GPS
			es.Url = e.Url
			es.IsRelative = e.IsRelative
			es.IsEndRelative = e.IsEndRelative
			es.Url = e.Url
			es.Title = e.Title
			es.Type = e.Type
			es.Importance = e.Importance
			doc.Events = append(doc.Events, es)

		}
	}

	for _, e := range doc.Events {
		if e.GetEpoch().ParentId >= 0 {
			e.GetEpoch().Parent = doc.Events[e.GetEpoch().ParentId]
		} else {
			e.GetEpoch().Parent = nil
		}
	}
}
