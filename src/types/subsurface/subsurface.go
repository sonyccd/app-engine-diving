package subsurface

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// Coordinate latitude and longitude
type Coordinate struct {
	Lat  float64
	Long float64
}

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
	GPS  Coordinate `xml:"gps,attr"`
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

	coordStrings := strings.Split(tempSite.GPS, " ")
	var validStrings []string
	for _, coordString := range coordStrings {
		if len(coordString) > 0 {
			validStrings = append(validStrings, coordString)
		}
	}

	var coordinates = make([]float64, 2)
	for i := 0; i < 2; i++ {
		if i >= len(validStrings) {
			coordinates[i] = 0
		} else {
			if value, err := strconv.ParseFloat(validStrings[i], 64); err != nil {
				coordinates[i] = 0
			} else {
				coordinates[i] = value
			}
		}
	}

	*s = Site{
		UUID: tempSite.UUID,
		Name: tempSite.Name,
		GPS:  Coordinate{Lat: coordinates[0], Long: coordinates[1]}}

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
