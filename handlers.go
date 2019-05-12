package gpsutils

import (
	"errors"
	"math"
	"strconv"
)

type coordinatesFormat struct {
	dd be
	mm be
	ss be
}

type be struct {
	b uint8
	e uint8
}

var typeCoordinatesMask = [3]string{"DD.DDDDDD°", "DD°MM.MMMM'", "DD°MM'SS.SSS\""}
var typeCoordinatesData = map[uint8]coordinatesFormat{
	1: coordinatesFormat{be{0, 9}, be{0, 0}, be{0, 0}},
	2: coordinatesFormat{be{0, 2}, be{4, 11}, be{0, 0}},
	3: coordinatesFormat{be{0, 2}, be{4, 6}, be{7, 13}},
}

func round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func parseCoordinateStr(strCoordinate string, coordinateType uint8, c *coordinate) (err error) {
	c.DD, err = strconv.ParseFloat(strCoordinate[typeCoordinatesData[coordinateType].dd.b:typeCoordinatesData[coordinateType].dd.e], 32)
	if err != nil {
		return errors.New("coordinates DD format error")
	}
	c.DD = round(c.DD, 6)
	if coordinateType > 1 {
		c.MM, err = strconv.ParseFloat(strCoordinate[typeCoordinatesData[coordinateType].mm.b:typeCoordinatesData[coordinateType].mm.e], 32)
		if err != nil {
			return errors.New("coordinates MM format error")
		}
		c.MM = round(c.MM, 4)
	}
	if coordinateType > 2 {
		c.SS, err = strconv.ParseFloat(strCoordinate[typeCoordinatesData[coordinateType].ss.b:typeCoordinatesData[coordinateType].ss.e], 32)
		if err != nil {
			return errors.New("coordinates SS format error")
		}
		c.SS = round(c.SS, 3)
	}

	return nil
}

func calcFormat1(out, in *coordinate) {
	out.DD = in.DD
	out.MM = round(math.Mod(in.DD, 1)*60, 4)
	out.SS = round(math.Mod(out.MM, 1)*60, 3)
}

func calcFormat2(out, in *coordinate) {
	out.MM = in.MM
	out.DD = round(in.DD+in.MM/60, 4)
	out.SS = round(math.Mod(in.MM, 1)*60, 3)
}

func calcFormat3(out, in *coordinate) {
	out.SS = in.SS
	out.MM = round(in.MM+in.SS/60, 4)
	out.DD = round(in.DD+out.MM/60, 6)
}
