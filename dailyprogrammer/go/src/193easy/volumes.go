package main

import (
	"math"
)

/*
 	Challenge 193 easy:
 	An international shipping company is trying to figure out how to manufacture various types of containers. Given a volume they want to figure out the dimensions of various shapes that would all hold the same volume.
 	Input:
		A volume in cubic meters.
	Output:
		Dimensions of containers of various types that would hold the volume. The following containers are possible.
		
		Cube
		Ball (Sphere)
		Cylinder
		Cone
	Example Input:
		27
	Example Output:
		Cube: 3.00m width, 3.00m, high, 3.00m tall
		Cylinder: 3.00m tall, Diameter of 3.38m
		Sphere: 1.86m Radius
		Cone: 9.00m tall, 1.69m Radius
	Some Inputs to test.
		27, 42, 1000, 2197
 */

type Cube struct {
	width float64
	height float64
	depth float64
}
type Cylinder struct {
	height float64
	diameter float64
}
type Sphere struct {
	radius float64
}
type Cone struct {
	height float64
	radius float64
}

func CubeByVolume(volume int) (Cube) {
	edgeLen := cbrt(float64(volume))
	return Cube{ edgeLen, edgeLen, edgeLen }
}

func CylinderByVolume(volume int) (Cylinder) {
	height := cbrt(float64(volume))
	/* http://de.wikipedia.org/wiki/Zylinder_(Geometrie)
		volume = area * height
		area = PI * radius²
		=> volume = PI * radius² * height
		=> PI * radius² = volume / height
		=> radius² = volume / height / PI
		=> radius = sqrt(volume / height / PI)
	*/
	diameter := math.Sqrt(float64(volume) / height / math.Pi) * 2

	return Cylinder{ height, diameter }
}

func SphereByVolume(volume int) (Sphere) {
	/* http://www.schulminator.com/mathematik/kugel
		volume = (4/3) * PI * r³
		=> volume / PI = 4/3 * r³
		=> (3 * volume / PI) / 4 = r³
		=> cbrt(3 * volume / PI / 4) = r
	*/
	radius := cbrt(3.0 * float64(volume) / math.Pi / 4.0)

	return Sphere{ radius }
}

func ConeByVolume(volume int) (Cone) {
	// TODO: why?!
	height := math.Pow(cbrt(float64(volume)), 2.0)
	/*
		http://www.mathepower.com/kegel.php
		volume = (area * height) / 3
		area = PI * radius²
		=> volume = 1/3 * PI * radius² * height
		=> volume / height = 1/3 * PI * radius²
		=> volume / height / PI = 1/3 * radius²
		=> 3 * volume / height / PI = radius²
		=> 3 * volume / height / PI = radius²
		=> sqrt(3 * volume / height / PI) = radius
	*/
	radius := math.Sqrt(3.0 * float64(volume) / height / math.Pi)

	return Cone{ height, radius }
}

func cbrt(x float64) float64 {	
	// http://stackoverflow.com/questions/13263477/how-to-take-a-cube-root-in-go
	return math.Pow(float64(x), 1.0 / 3)
}