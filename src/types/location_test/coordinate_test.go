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
			coord location.Coordinate
		)

		AfterEach(func() {
			coord = location.Coordinate{}
		})

		Context("valid", func() {
			BeforeEach(func() {
				siteStr = "35.910920 -78.426640"
				coord = location.StringToCoordinate(siteStr)
			})

			Context("attributes", func() {
				It("Latitude is 1st part of gps attribute", func() {
					Expect(coord.Lat).To(Equal(35.910920))
				})

				It("Longitude is 2nd part of gps attribute", func() {
					Expect(coord.Long).To(Equal(-78.426640))
				})
			})

			Context("bad coordinates", func() {
				It("one number is mapped to latitude, longitude is zero", func() {
					siteStr = "35.910920"
					coord = location.StringToCoordinate(siteStr)

					Expect(coord.Lat).To(Equal(35.910920))
					Expect(coord.Long).To(Equal(0.0)) 
				})

				It("third number is ignored", func() {
					siteStr = "35.910920 -78.426640 43.34"
					coord = location.StringToCoordinate(siteStr)

					Expect(coord.Lat).To(Equal(35.910920))
					Expect(coord.Long).To(Equal(-78.426640))
				})

				It("Lat and Long are zero when there are no values", func() {
					siteStr = ""
					coord = location.StringToCoordinate(siteStr)

					Expect(coord.Lat).To(Equal(0.0))
					Expect(coord.Long).To(Equal(0.0))
				})

				It("Lat and Long are zero when the GPS is a string", func() {
					siteStr = "my magical location"
					coord = location.StringToCoordinate(siteStr)

					Expect(coord.Lat).To(Equal(0.0))
					Expect(coord.Long).To(Equal(0.0))
				})

				Context("Lat and Long are mapped when", func() {
					It("there is an extra space between", func() {
						siteStr = "35.910920  -78.426640"
						coord = location.StringToCoordinate(siteStr)
	
						Expect(coord.Lat).To(Equal(35.910920))
						Expect(coord.Long).To(Equal(-78.426640))
					})

					It("there is a leading space", func() {
						siteStr = " 35.910920 -78.426640"
						coord = location.StringToCoordinate(siteStr)
	
						Expect(coord.Lat).To(Equal(35.910920))
						Expect(coord.Long).To(Equal(-78.426640))
					})

					It("there is a trailing space", func() {
						siteStr = "35.910920  -78.426640 "
						coord = location.StringToCoordinate(siteStr)
	
						Expect(coord.Lat).To(Equal(35.910920))
						Expect(coord.Long).To(Equal(-78.426640))
					})

					It("has a lot of space in between", func() {
						siteStr = "    35.910920     -78.426640     "
						coord = location.StringToCoordinate(siteStr)
	
						Expect(coord.Lat).To(Equal(35.910920))
						Expect(coord.Long).To(Equal(-78.426640))
					})
				})

				
			})
		})

	})

})
