package main

import (
	"math"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTwoUnitCircleArea_NoIntersection(t *testing.T) {
	output := TwoUnitCircleArea(-1, 0, 1, 0)
	assert.Equal(t, 0, output)
}

func TestTwoUnitCircleArea_xIntersection(t *testing.T) {
	output := TwoUnitCircleArea(-0.5, 0, 0.5, 0)
	roundedOutput := math.Floor(output * 10000) / 10000
	assert.Equal(t, 5.0548, roundedOutput)
}

func TestTwoUnitCircleArea_yIntersection(t *testing.T) {
	output := TwoUnitCircleArea(0, 0.5, 0, -0.5)
	roundedOutput := math.Floor(output * 10000) / 10000
	assert.Equal(t, 5.0548, roundedOutput)
}