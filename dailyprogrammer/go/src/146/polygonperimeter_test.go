package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
)

func TestPolygonPerimeter_1(t *testing.T) {
	output := PolygonPerimeter(5, 3.7)
	output = math.Floor(output * 1000) / 1000
	
	assert.Equal(t, 21.748, output)
}

func TestPolygonPerimeter_2(t *testing.T) {
	output := PolygonPerimeter(100, 1.0)
	output = math.Floor(output * 1000) / 1000
	
	assert.Equal(t, 6.282, output)
}
