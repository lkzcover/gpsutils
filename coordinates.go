package gpsutils

import "fmt"

// Coordinates basic structure to work with this package
type Coordinates struct {
	N           coordinate
	E           coordinate
	CoordinateN string // N coordinates in the specified format
	CoordinateE string // E coordinates in the specified format
	Format      string // format coordinates
	formatType  uint8  // internal identifier coordinates type
}

type coordinate struct {
	DD float64 // degree
	MM float64 // min
	SS float64 // sec
}

// buildCoordinates the method collects coordinates in the structure parametres CoordinateN and CoordinateE
func (obj *Coordinates) buildCoordinates() {
	switch obj.formatType {
	case 1:
		{
			//"DD.DDDDDD°"
			obj.CoordinateN = fmt.Sprintf("%f°", obj.N.DD)
			obj.CoordinateE = fmt.Sprintf("%f°", obj.E.DD)

			//TODO протестировать, рассмотреть возможность перехода на math.Mod
			obj.N.MM = 0
			obj.E.MM = 0
			obj.N.SS = 0
			obj.E.SS = 0
		}
	case 2:
		{
			//"DD°MM.MMMM'"
			obj.CoordinateN = fmt.Sprintf("%.0f°%.4f'", obj.N.DD, obj.N.MM)
			obj.CoordinateE = fmt.Sprintf("%.0f°%.4f'", obj.E.DD, obj.E.MM)

			//TODO протестировать, рассмотреть возможность перехода на math.Mod
			obj.N.DD = round(obj.N.DD, 0)
			obj.E.DD = round(obj.N.DD, 0)
			obj.N.SS = 0
			obj.E.SS = 0
		}
	case 3:
		{
			//"DD°MM'SS.SSS\""
			obj.CoordinateN = fmt.Sprintf("%.0f°%.0f'%.3f\"", obj.N.DD, obj.N.MM, obj.N.SS)
			obj.CoordinateE = fmt.Sprintf("%.0f°%.0f'%.3f\"", obj.E.DD, obj.E.MM, obj.E.SS)

			//TODO протестировать, рассмотреть возможность перехода на math.Mod
			obj.N.DD = round(obj.N.DD, 0)
			obj.E.DD = round(obj.N.DD, 0)
			obj.N.MM = round(obj.N.MM, 0)
			obj.E.MM = round(obj.N.MM, 0)
		}
	}
}
