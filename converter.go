package gpsutils

import (
	"errors"
	"fmt"
)

// ConvertCoordinates get the mask for the input and output data, get the coordinates, and return the coordinates in the Coordinates data type
func ConvertCoordinates(maskIn, maskOut, coordinatesN, coordinatesE string) (*Coordinates, error) {

	input, err := ValidateCoordinatesData(coordinatesN, coordinatesE, maskIn)
	if err != nil {
		return nil, fmt.Errorf("input data %s", err)
	}

	if maskIn == maskOut {
		return nil, errors.New("masks are equal")
	}

	var output Coordinates

	if output.formatType, err = ValidateCoordinatesMask(maskOut); err != nil {
		return nil, fmt.Errorf("output %s", err)
	}
	output.Format = maskOut

	switch input.formatType {
	case 1:
		{
			calcFormat1(&output.N, &input.N)
			calcFormat1(&output.E, &input.E)
		}
	case 2:
		{
			calcFormat2(&output.N, &input.N)
			calcFormat2(&output.E, &input.E)
		}
	case 3:
		{
			calcFormat3(&output.N, &input.N)
			calcFormat3(&output.E, &input.E)
		}
	}

	output.buildCoordinates()

	return &output, nil
}
