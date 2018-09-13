package subsurface

import (
	"encoding/xml"

	"github.com/sonyccd/app-engine-diving/src/types/location"
)

type Subsurface struct {
	XMLName   xml.Name       `xml:"divelog"`
	Version   string         `xml:"version,attr"`
	Program   string         `xml:"program,attr"`
	Settings  DiveComputerId `xml:"settings>divecomputerid"`
	DiveSites []Site         `xml:"divesites>site"`
	Dives     []Dive         `xml:"dives>dive"`
}

type DiveComputerId struct {
	Model        string `xml:"model,attr"`
	DeviceId     string `xml:"deviceid,attr"`
	SerialNumber string `xml:"serial,attr"`
}

type Site struct {
	UUID string     `xml:"uuid,attr"`
	Name string     `xml:"name,attr"`
	GPS  location.Coordinate `xml:"gps,attr"`
}

// UnmarshalXML will map Site and the GPS property to a coordinate
func (s *Site) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tempSite := struct {
		UUID string `xml:"uuid,attr"`
		Name string `xml:"name,attr"`
		GPS  string `xml:"gps,attr"`
	}{}

	if err := d.DecodeElement(&tempSite, &start); err != nil {
		return err
	}

	*s = Site{
		UUID: tempSite.UUID,
		Name: tempSite.Name,
		GPS:  location.StringToCoordinate(tempSite.GPS),
	}

	return nil
}

type Dive struct {
	Number   uint   `xml:"number,attr"`
	Date     string `xml:"date,attr"`
	Time     string `xml:"time,attr"`
	Duration string `xml:"duration,attr"`

	Cylinder     Cylinder     `xml:"cylinder"`
	DiveComputer DiveComputer `xml:"divecomputer"`
}

type Cylinder struct {
	Size         string `xml:"size,attr"`
	WorkPressure string `xml:"workpressure,attr"`
	Description  string `xml:"description,attr"`
}

type DiveComputer struct {
	Model    string `xml:"model,attr"`
	DeviceId string `xml:"deviceid,attr"`

	Depth       Depth       `xml:"depth"`
	Temperature Temperature `xml:"temperature"`
	Events      []Event     `xml:"event"`
	Samples     []Sample    `xml:"sample"`
}

type Depth struct {
	Max  string `xml:"max,attr"`
	Mean string `xml:"mean,attr"`
}

type Temperature struct {
	Air   string `xml:"air,attr"`
	Water string `xml:"water,attr"`
}

type Event struct {
	Time     string `xml:"time,attr"`
	Type     string `xml:"type,attr"`
	Value    string `xml:"value,attr"`
	Name     string `xml:"name,attr"`
	Cylinder string `xml:"cylinder,attr"`
}

type Sample struct {
	Time  string `xml:"time,attr"`
	Depth string `xml:"depth,attr"`
	Temp  string `xml:"temp,attr"`
}
