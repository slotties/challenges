package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFletcher16_1(t *testing.T) {
	output := Fletcher16("Fletcher")
	var expected uint16 = 0xD330
	
	assert.Equal(t, expected, output)
}

func TestFletcher16_2(t *testing.T) {
	output := Fletcher16("Sally sells seashells by the seashore.")
	var expected uint16 = 0xD23E
	
	assert.Equal(t, expected, output)
}

func TestFletcher16_3(t *testing.T) {
	output := Fletcher16("Les chaussettes de l'archi-duchesse, sont-elles seches ou archi-seches ?")
	var expected uint16 = 0x404D
	
	assert.Equal(t, expected, output)
}