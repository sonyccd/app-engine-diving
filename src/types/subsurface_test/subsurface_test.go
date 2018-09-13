package subsurface_test

import (
	"encoding/xml"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sonyccd/app-engine-diving/src/types/subsurface"
)

var _ = Describe("XML un-marshalling", func() {
	var ()

	BeforeEach(func() {

	})

	JustBeforeEach(func() {

	})

	Context("Divelog to subsurface", func() {
		var (
			diveLog subsurface.Subsurface
			err     error
		)

		AfterEach(func() {
			diveLog = subsurface.Subsurface{}
		})

		Context("valid tag and attributes", func() {
			BeforeEach(func() {
				diveLogStr := `<divelog program='subsurface' version='3'></divelog>`
				err = xml.Unmarshal([]byte(diveLogStr), &diveLog)
			})

			It("no error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("contains xml name divelog", func() {
				name := xml.Name{Local: "divelog"}
				Expect(diveLog.XMLName).To(Equal(name))
			})

			It("contains version attribute", func() {
				Expect(diveLog.Version).To(Equal("3"))
			})

			It("contains program attribute", func() {
				Expect(diveLog.Program).To(Equal("subsurface"))
			})

			It("zero dive sites", func() {
				Expect(diveLog.DiveSites).To(BeNil())
			})

			It("zero dives", func() {
				Expect(diveLog.Dives).To(BeNil())
			})

			It("dive computer is zero object", func() {
				Expect(diveLog.Settings).To(Equal(subsurface.DiveComputerId{}))
			})
		})

		Describe("settings to dive computer", func() {
			Context("valid", func() {
				BeforeEach(func() {
					diveLogStr := `
					<divelog program='subsurface' version='3'>
						<settings>
							<divecomputerid model='Scubapro Matrix' deviceid='ffffffff' serial='63005620'/>
						</settings>
					</divelog>`
					err = xml.Unmarshal([]byte(diveLogStr), &diveLog)
				})

				Specify("deviceID is deviceid attribute", func() {
					Expect(diveLog.Settings.DeviceId).To(Equal("ffffffff"))
				})

				Specify("model is model attribute", func() {
					Expect(diveLog.Settings.Model).To(Equal("Scubapro Matrix"))
				})

				Specify("serialNumber is serial attribute", func() {
					Expect(diveLog.Settings.SerialNumber).To(Equal("63005620"))
				})
			})

			Context("missing attributes", func() {
				BeforeEach(func() {
					diveLogStr := `
					<divelog program='subsurface' version='3'>
						<settings>
							<divecomputerid/>
						</settings>
					</divelog>`
					err = xml.Unmarshal([]byte(diveLogStr), &diveLog)
				})

				Specify("deviceID is empty string", func() {
					Expect(diveLog.Settings.DeviceId).To(BeEmpty())
				})

				Specify("model is empty string", func() {
					Expect(diveLog.Settings.Model).To(BeEmpty())
				})

				Specify("serialNumber is empty string", func() {
					Expect(diveLog.Settings.SerialNumber).To(BeEmpty())
				})
			})

		})

		It("invalid tag throws error", func() {
			diveLogStr := `<divelogs program='subsurface' version='3'></divelogs>`
			err = xml.Unmarshal([]byte(diveLogStr), &diveLog)
			Expect(err).To(HaveOccurred())
		})

		Context("missing attribute", func() {
			It("version is empty string", func() {
				diveLogStr := `<divelog program='subsurface' versions='3'></divelog>`
				err = xml.Unmarshal([]byte(diveLogStr), &diveLog)
				Expect(diveLog.Version).To(BeEmpty())
			})

			It("program is empty string", func() {
				diveLogStr := `<divelog programs='subsurface' version='3'></divelog>`
				err = xml.Unmarshal([]byte(diveLogStr), &diveLog)
				Expect(diveLog.Program).To(BeEmpty())
			})
		})

	})

	Describe("Dive Sites", func() {
		var (
			diveSite subsurface.Site
			siteStr string
			err     error
		)
		
		AfterEach(func() {
			diveSite = subsurface.Site{}
		})

		Context("valid", func() {
			BeforeEach(func() {
				siteStr = `<site uuid='12345' name='3601 Quarry Rd, Rolesville, NC 27571 / Lake Park' gps='35.910920 -78.426640'></site>`
				err = xml.Unmarshal([]byte(siteStr), &diveSite)
			})

			It("no error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			Context("attributes", func() {
				Specify("UUID is uuid attribute", func() {
					Expect(diveSite.UUID).To(Equal("12345"))
				})

				Specify("Name is name attribute", func() {
					Expect(diveSite.Name).To(Equal("3601 Quarry Rd, Rolesville, NC 27571 / Lake Park"))
				})

				It("Latitude is 1st part of gps attribute", func() {
					Expect(diveSite.GPS.Lat).To(Equal(35.910920))
				})

				It("Longitude is 2nd part of gps attribute", func() {
					Expect(diveSite.GPS.Long).To(Equal(-78.426640))
				})
			})
		})

	})

})
