package location

import (
	"strconv"
	"strings"
)

// Coordinate latitude and longitude
type Coordinate struct {
	Lat  float64
	Long float64
}

// StringToCoordinate will try to transform the string into a set of coordinates
// Invalid fields will be returned as 0.0
func StringToCoordinate(coord string) Coordinate {
	coordStrings := strings.Split(coord, " ")
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

	return Coordinate{
		Lat: coordinates[0],
		Long: coordinates[1],
	}
}