package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestChallenge(t *testing.T) {
	file, err := os.Open("enable1.txt")
	assert.Nil(t, err)
	defer file.Close()

	word := GetWordWithMostSubwords(file)
	assert.Equal(t, "ethylenediaminetetraacetates", word)
}