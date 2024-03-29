package coordinategen

import (
	"math"
	"math/rand"
)

type CoordinateGen struct {
	centralLatitude  float64
	centralLongitude float64
	radius           float64
}

func New(centralLatitude, centralLongitude, radius float64) *CoordinateGen {
	return &CoordinateGen{
		centralLatitude:  centralLatitude,
		centralLongitude: centralLongitude,
		radius:           radius,
	}
}

func (c *CoordinateGen) randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func (c *CoordinateGen) round(x float64, prec int) float64 {
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

func (c *CoordinateGen) GenCoordinates() (float64, float64) {
	radiusInDegrees := c.radius / 111000.0

	u := c.randFloat(0, 1)
	v := c.randFloat(0, 1)

	w := radiusInDegrees * math.Sqrt(u)
	tN := 2 * math.Pi * v
	x := w * math.Cos(tN)
	y := w * math.Sin(tN)

	newX := x / math.Cos(math.Pi*c.centralLongitude/180.0)

	foundLongitude := c.round(newX+c.centralLongitude, 6)
	foundLatitude := c.round(y+c.centralLatitude, 6)

	return foundLatitude, foundLongitude
}
