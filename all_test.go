package gpsutils

import "testing"

func TestValidateCoordinatesMask(t *testing.T) {
	masks := [3]string{"DD.DDDDDD°", "DD°MM.MMMM'", "DD°MM'SS.SSS\""}

	for i, mask := range masks {
		typeMask, err := ValidateCoordinatesMask(mask)
		if err != nil || typeMask != uint8(i+1) {
			t.Errorf("validator error with mask %s", mask)
		}
	}

	_, err := ValidateCoordinatesMask("bilebrda")
	if err == nil {
		t.Errorf("validator error with mask bileberda")
	}
}

func TestValidateCoordinatesData(t *testing.T) {
	masks := [3]string{"DD.DDDDDD°", "DD°MM.MMMM'", "DD°MM'SS.SSS\""}
	datasN := [3]string{"11.111111°", "11°22.2222'", "11°22'33.333\""}

	for i, mask := range masks {
		c, err := ValidateCoordinatesData(datasN[i], datasN[i], mask)
		t.Logf("%+v", c)

		if err != nil {
			t.Errorf("validate coordinates data error: %s", err)
		}

	}
}

func TestConvertCoordinates(t *testing.T) {
	masks := [3]string{"DD.DDDDDD°", "DD°MM.MMMM'", "DD°MM'SS.SSS\""}
	masks2 := [3]string{"DD°MM.MMMM'", "DD°MM'SS.SSS\"", "DD.DDDDDD°"}
	datasN := [3]string{"23.755575°", "23°45.3345'", "23°45'20.100\""}
	datasE := [3]string{"12.120575°", "12°67.2345'", "12°67'14.100\""}

	for i, mask := range masks {
		c, err := ConvertCoordinates(mask, masks2[i], datasN[i], datasE[i])
		t.Logf("%+v", c)

		if err != nil {
			t.Errorf("convert coordinates data error: %s", err)
		}

	}
}

func TestParseCoordinates(t *testing.T) {
	masks := [3]string{"DD.DDDDDD°", "DD°MM.MMMM'", "DD°MM'SS.SSS\""}
	datasN := [3]string{"23.755575°", "23°45.3345'", "23°45'20.100\""}
	datasE := [3]string{"12.120575°", "12°67.2345'", "12°67'14.100\""}

	c, err := ParseCoordinates(datasN[0], datasE[0], masks[0])
	if err != nil {
		t.Errorf("parse coordinate error: %s", err)
	}
	t.Logf("%+v", c)
}

func TestDistanceCalculation(t *testing.T) {
	mask := "DD.DDDDDD°"
	datasN := [2]string{"50.879874°", "50.879840°'"}
	datasE := [2]string{"34.808982°", "34.808827°"}

	c1, err := ParseCoordinates(datasN[0], datasE[0], mask)
	if err != nil {
		t.Errorf("parse coordinate error: %s", err)
	}

	c2, err := ParseCoordinates(datasN[1], datasE[1], mask)
	if err != nil {
		t.Errorf("parse coordinate error: %s", err)
	}
	distance, err := DistanceCalculation(c1, c2)
	if err != nil {
		t.Errorf("distance calcultion error: %s", err)
	}
	t.Log(distance)
}
