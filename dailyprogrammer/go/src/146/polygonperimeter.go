package main

import "math"

/*
	Regular polygon:
	http://de.wikipedia.org/wiki/Polygon#Regelm.C3.A4.C3.9Fige_Polygone
	
	Formula for the perimeter:
	r * 2n * sin(PI / n)
	
	r = circumradius
	n = number of sides
*/
func PolygonPerimeter(n uint8, r float64) float64 {
	if n < 3 || n > 100 {
		return -1.0; // FIXME: through error
	}

	return r * 2 * float64(n) * math.Sin(math.Pi / float64(n))
}