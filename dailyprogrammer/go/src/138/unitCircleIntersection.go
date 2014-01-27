package main

import "math"

/*
	http://en.wikipedia.org/wiki/Circular_segment
	
	Given:
		radius is always 1 (because of unit circles)
	
	Formulas:
	angle = theta = 2 * arccos(height / radius)	
				  = 2 * arccos(height)
	area of circular segment = 0.5 * (theta - sin(theta)) * radius²
							 = 0.5 * (theta - sin(theta))
	area of circle = 2 * pi
	
	area of both circles:
		A = 2 * pi - 2 * (0.5 * (theta - sin(theta))	
		  = 2 * pi - (theta - sin(theta))
*/
func TwoUnitCircleArea(x float64, y float64, u float64, w float64) float64 {
	distance := distance(x, y, u, w)
	if distance < 2.0 {
		height := distance / 2.0
		theta := 2 * math.Acos(height)
		A := 2 * math.Pi - (theta - math.Sin(theta))
		
		return A
	} else {
		// No intersection
		return 0.0
	}
}

func distance(x float64, y float64, u float64, w float64) float64 {
	return math.Hypot(x - u, y - w)
	/*
	dx := x - u
	dy := y - w
	return math.Sqrt(dx * dx + dy * dy)
	*/
}