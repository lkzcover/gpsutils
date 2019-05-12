package gpsutils

import (
	"errors"
)

// ValidateCoordinatesMask get mask and return internal type mask coordinates and error
func ValidateCoordinatesMask(mask string) (uint8, error) {
	for i, tCM := range typeCoordinatesMask {
		if mask == tCM {
			return uint8(i + 1), nil
		}
	}
	return 0, errors.New("data mask error")
}

// ValidateCoordinatesData get coordinates data and mask, and return the coordinates in the Coordinates data type and the validity of the pair
func ValidateCoordinatesData(coordinatesN, coordinatesE, mask string) (*Coordinates, error) {
	var c Coordinates
	var err error

	if c.formatType, err = ValidateCoordinatesMask(mask); err != nil {
		return nil, err
	}

	if len(coordinatesN) < len(mask) || len(coordinatesE) < len(mask) {
		return nil, errors.New("incorrect coordinates format")
	}

	if err = parseCoordinateStr(coordinatesN, c.formatType, &c.N); err != nil {
		return nil, err
	}

	if err = parseCoordinateStr(coordinatesE, c.formatType, &c.E); err != nil {
		return nil, err
	}

	return &c, nil
}
