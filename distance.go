package gpsutils

import (
	"errors"
	"math"
)

const sdM = 111.194926645 // Удельное расстрояние по мередиану

// DistanceCalculation calculates the distance between two GPS coordinates
func DistanceCalculation(c1, c2 *Coordinates) (distance float64, err error) {
	if c1.formatType != c2.formatType {
		return 0, errors.New("coordinates format are not equal")
	}

	distance = math.Sqrt(math.Pow((c1.N.DD-c2.N.DD)*sdM, 2) + math.Pow((c1.E.DD-c2.E.DD)*(math.Cos(c1.N.DD)*sdM), 2))

	return
}
