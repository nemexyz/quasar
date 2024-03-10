package domain

import (
	"math"
)

type Satellite struct {
	name        string
	distance    float64
	message     []string
	coordinates [2]float64
}

func NewSatellite(name string, coordinates [2]float64) *Satellite {
	return &Satellite{
		name:        name,
		coordinates: coordinates,
	}
}

func (s *Satellite) ReceiveMessage(distance float64, message []string) {
	s.distance = distance
	s.message = message
}

func (s *Satellite) FixMsgDelay(realLength int) {
	s.message = s.message[len(s.message)-realLength:]
}

func (s *Satellite) IntersectionWith(satellite *Satellite) [][2]float64 {
	return s.GetRadiusIntersection(s.coordinates, s.distance, satellite.coordinates, satellite.distance)
}

func (s *Satellite) GetRadiusIntersection(p0 [2]float64, r0 float64, p1 [2]float64, r1 float64) [][2]float64 {
	x0 := p0[0]
	y0 := p0[1]
	x1 := p1[0]
	y1 := p1[1]
	dx := x1 - x0
	dy := y1 - y0
	d := math.Hypot(dx, dy)

	// No solution. circles do not intersect
	if d > (r0 + r1) {
		panic("Error: cannot find location with given data")
	}
	// No solution. one circle is contained in the other
	if d < math.Abs(r0-r1) {
		panic("Error: cannot find location with given data")
	}
	// No solution. circles are the same
	if d == 0 && r0 == r1 {
		panic("Error: cannot find location with given data")
	}

	a := ((r0 * r0) - (r1 * r1) + (d * d)) / (2.0 * d)
	x2 := x0 + (dx * a / d)
	y2 := y0 + (dy * a / d)

	// h is the distance from point (x2, y2) to either of the intersection points
	h := math.Sqrt((r0 * r0) - (a * a))
	rx := -dy * (h / d)
	ry := dx * (h / d)

	// Get the intersection points
	xi := round(x2 + rx)
	xiP := round(x2 - rx)
	yi := round(y2 + ry)
	yiP := round(y2 - ry)

	return [][2]float64{{xi, yi}, {xiP, yiP}}
}

func round(num float64) float64 {
	return math.Round(num*100) / 100
}
