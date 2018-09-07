package subsurface_test

import (
	"encoding/xml"
	"testing"

	"github.com/sonyccd/app-engine-diving/src/types/subsurface"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type SubsurfaceTest struct{}

var _ = check.Suite(&SubsurfaceTest{})

func (s SubsurfaceTest) SetUpTest(c *check.C) {

}

func (s SubsurfaceTest) TestSubsurfaceObject(c *check.C) {
	diveLogStr := `<divelog program='subsurface' version='3'></divelog>`

	var diveLog subsurface.Subsurface
	err := xml.Unmarshal([]byte(diveLogStr), &diveLog)

	c.Assert(diveLog, check.DeepEquals, subsurface.Subsurface{
		XMLName: xml.Name{Local: "divelog"},
	})
	c.Assert(err, check.IsNil)
}

func (s SubsurfaceTest) TestSubsurfaceObjectBad(c *check.C) {
	diveLogStr := `<divelogs program='subsurface' version='3'></divelog>`

	var diveLog subsurface.Subsurface
	err := xml.Unmarshal([]byte(diveLogStr), &diveLog)

	c.Assert(diveLog, check.DeepEquals, subsurface.Subsurface{})
	c.Assert(err, check.NotNil)
}

func (s SubsurfaceTest) TestZeroDiveSites(c *check.C) {
	diveLogStr := `<divelog program='subsurface' version='3'>
					<divesites></divesites>
				  </divelog>`

	var diveLog subsurface.Subsurface
	err := xml.Unmarshal([]byte(diveLogStr), &diveLog)

	expectedDiveLog := subsurface.Subsurface{
		XMLName: xml.Name{Local: "divelog"},
		DiveSites: nil,
	}

	c.Assert(diveLog, check.DeepEquals, expectedDiveLog)
	c.Assert(err, check.IsNil)
}

func (s SubsurfaceTest) TestOneDiveSite(c *check.C) {
	diveLogStr := `<divelog program='subsurface' version='3'>
	<divesites>
        <site uuid='12345' name='test name' gps='35.200 -78.400'>
        </site>
    </divesites>
	</divelog>`

	var diveLog subsurface.Subsurface
	err := xml.Unmarshal([]byte(diveLogStr), &diveLog)

	expectedDiveLog := subsurface.Subsurface{
		XMLName: xml.Name{Local: "divelog"},
		DiveSites: []subsurface.Site{
			{UUID: "12345", Name: "test name", GPS: "35.200 -78.400"},
		},
	}

	c.Assert(diveLog, check.DeepEquals, expectedDiveLog)
	c.Assert(err, check.IsNil)
}

func (s SubsurfaceTest) TestDiveSiteInvalid(c *check.C) {

	testCases := []struct {
		diveSite string
		expected subsurface.Site
	}{
		{"<site uuids='12345' name='test name' gps='35.200 -78.400'></site>", subsurface.Site{UUID: "", Name: "test name", GPS: "35.200 -78.400"}},
		{"<site uuid='12345' names='test name' gps='35.200 -78.400'></site>", subsurface.Site{UUID: "12345", Name: "", GPS: "35.200 -78.400"}},
		{"<site uuid='12345' name='test name' gpss='35.200 -78.400'></site>", subsurface.Site{UUID: "12345", Name: "test name", GPS: ""}},
		{"<site uuids='12345' names='test name' gps='35.200 -78.400'></site>", subsurface.Site{UUID: "", Name: "", GPS: "35.200 -78.400"}},
		{"<site uuid='12345' names='test name' gpss='35.200 -78.400'></site>", subsurface.Site{UUID: "12345", Name: "", GPS: ""}},
		{"<site uuids='12345' name='test name' gpss='35.200 -78.400'></site>", subsurface.Site{UUID: "", Name: "test name", GPS: ""}},
		{"<site uuids='12345' names='test name' gpss='35.200 -78.400'></site>", subsurface.Site{UUID: "", Name: "", GPS: ""}},
	}
	for _, testCase := range testCases {
		diveLogStr := `<divelog program='subsurface' version='3'>
		<divesites>` + testCase.diveSite + `</divesites></divelog>`

		var diveLog subsurface.Subsurface
		err := xml.Unmarshal([]byte(diveLogStr), &diveLog)

		expectedDiveLog := subsurface.Subsurface{
			XMLName: xml.Name{Local: "divelog"},
			DiveSites: []subsurface.Site{
				testCase.expected,
			},
		}

		c.Assert(diveLog, check.DeepEquals, expectedDiveLog)
		c.Assert(err, check.IsNil)
	}
}
