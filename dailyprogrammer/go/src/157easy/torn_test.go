package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func TestIsTornNumber(t *testing.T) {
	for i := 1; i <= 9999; i++ {
		if i == 3025 || i == 9801 {
			assert.True(t, IsTornNumber(i), strconv.Itoa(i) + " should be a torn number")
		} else {
			assert.False(t, IsTornNumber(i), strconv.Itoa(i) + " should NOT be a torn number")
		}
	}
}