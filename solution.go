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

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

func CalcSquare(sideLen float64, sidesNum int) float64 {
	var total float64
	switch sidesNum {
	case SidesTriangle:
		total = (math.Sqrt(3) * sideLen * sideLen) / 4
	case SidesSquare:
		total = sideLen * sideLen
	case SidesCircle:
		total = math.Pi * sideLen * sideLen
	default:
		total = 0
	}
	return total
}

func main() {
	CalcSquare(10.0, SidesTriangle)
	CalcSquare(10.0, SidesSquare)
	CalcSquare(10.0, SidesCircle)
}
