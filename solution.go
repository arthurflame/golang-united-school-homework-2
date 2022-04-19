package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

//func CalcSquare(sideLen float64, sidesNum #yourTypeNameHere#) float64 {
//}

type smartType int

const (
	SidesCircle   smartType = 0
	SidesTriangle           = 3
	SidesSquare             = 4
)

func CalcSquare(sideLen float64, sidesNum smartType) float64 {
	switch sidesNum {
	case 0:
		return math.Pi * (math.Pow(sideLen, 2))
	case 3:
		thirdSideLen := sideLen
		perimeter := (sideLen*2 + thirdSideLen) / 2
		area := math.Sqrt(perimeter * (perimeter - sideLen))
		return area
	case 4:
		return sideLen * sideLen
	default:
		return 0
	}
}
