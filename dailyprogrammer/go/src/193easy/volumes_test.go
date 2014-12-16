package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
)

func TestCube(t *testing.T) {
	cube := CubeByVolume(27)

	assert.Equal(t, 3.0, float32(cube.height))
	assert.Equal(t, 3.0, float32(cube.width))
	assert.Equal(t, 3.0, float32(cube.depth))
}

func TestCylinder(t *testing.T) {
	cylinder := CylinderByVolume(27)

	assert.Equal(t, 3.0, float32(cylinder.height))
	assert.Equal(t, 3.38, roundToPrecision(cylinder.diameter, 2))
}

func TestSphere(t *testing.T) {
	sphere := SphereByVolume(27)

	assert.Equal(t, 1.86, roundToPrecision(sphere.radius, 2))
}

func TestCone(t *testing.T) {
	cone := ConeByVolume(27)

	assert.Equal(t, 1.69, roundToPrecision(cone.radius, 2))
	assert.Equal(t, 9, float32(cone.height))
}

func roundToPrecision(x float64, precision int) float32 {
	y := math.Pow10(precision)
	return float32(float32(int(x * y)) / float32(y))
}