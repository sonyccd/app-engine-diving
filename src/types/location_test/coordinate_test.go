package location_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sonyccd/app-engine-diving/src/types/location"
)

var _ = Describe("Coordinates", func() {
	var ()

	Describe("strings to coordinates", func() {
		var (
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
				It("Latitude is 1st part of gps attribute", func() {
					Expect(diveSite.GPS.Lat).To(Equal(35.910920))
				})

				It("Longitude is 2nd part of gps attribute", func() {
					Expect(diveSite.GPS.Long).To(Equal(-78.426640))
				})
			})

			Context("bad coordinates", func() {
				It("one number is mapped to latitude, longitude is zero", func() {
					siteStr = `<site uuid='12345' name='3601 Quarry Rd, Rolesville, NC 27571 / Lake Park' gps='35.910920'></site>`
					err = xml.Unmarshal([]byte(siteStr), &diveSite)

					Expect(diveSite.GPS.Lat).To(Equal(35.910920))
					Expect(diveSite.GPS.Long).To(Equal(0.0)) 
				})

				It("third number is ignored", func() {
					siteStr = `<site uuid='12345' name='3601 Quarry Rd, Rolesville, NC 27571 / Lake Park' gps='35.910920 -78.426640 43.34'></site>`
					err = xml.Unmarshal([]byte(siteStr), &diveSite)

					Expect(diveSite.GPS.Lat).To(Equal(35.910920))
					Expect(diveSite.GPS.Long).To(Equal(-78.426640))
				})

				It("Lat and Long are zero when there are no values", func() {
					siteStr = `<site uuid='12345' name='3601 Quarry Rd, Rolesville, NC 27571 / Lake Park' gps=''></site>`
					err = xml.Unmarshal([]byte(siteStr), &diveSite)

					Expect(diveSite.GPS.Lat).To(Equal(0.0))
					Expect(diveSite.GPS.Long).To(Equal(0.0))
				})

				It("Lat and Long are zero when the GPS is a string", func() {
					siteStr = `<site uuid='12345' name='3601 Quarry Rd, Rolesville, NC 27571 / Lake Park' gps='my magical location'></site>`
					err = xml.Unmarshal([]byte(siteStr), &diveSite)

					Expect(diveSite.GPS.Lat).To(Equal(0.0))
					Expect(diveSite.GPS.Long).To(Equal(0.0))
				})

				Context("Lat and Long are mapped when", func() {
					It("there is an extra space between", func() {
						siteStr = `<site uuid='12345' name='3601 Quarry Rd, Rolesville, NC 27571 / Lake Park' gps='35.910920  -78.426640'></site>`
						err = xml.Unmarshal([]byte(siteStr), &diveSite)
	
						Expect(diveSite.GPS.Lat).To(Equal(35.910920))
						Expect(diveSite.GPS.Long).To(Equal(-78.426640))
					})

					It("there is a leading space", func() {
						siteStr = `<site uuid='12345' name='3601 Quarry Rd, Rolesville, NC 27571 / Lake Park' gps=' 35.910920 -78.426640'></site>`
						err = xml.Unmarshal([]byte(siteStr), &diveSite)
	
						Expect(diveSite.GPS.Lat).To(Equal(35.910920))
						Expect(diveSite.GPS.Long).To(Equal(-78.426640))
					})

					It("there is a trailing space", func() {
						siteStr = `<site uuid='12345' name='3601 Quarry Rd, Rolesville, NC 27571 / Lake Park' gps='35.910920  -78.426640 '></site>`
						err = xml.Unmarshal([]byte(siteStr), &diveSite)
	
						Expect(diveSite.GPS.Lat).To(Equal(35.910920))
						Expect(diveSite.GPS.Long).To(Equal(-78.426640))
					})

					It("has a lot of space in between", func() {
						siteStr = `<site uuid='12345' name='3601 Quarry Rd, Rolesville, NC 27571 / Lake Park' gps='    35.910920     -78.426640     '></site>`
						err = xml.Unmarshal([]byte(siteStr), &diveSite)
	
						Expect(diveSite.GPS.Lat).To(Equal(35.910920))
						Expect(diveSite.GPS.Long).To(Equal(-78.426640))
					})
				})

				
			})
		})

	})

})
