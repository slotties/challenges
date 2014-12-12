package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
	"os"
)

func TestCheckSpelling_Success(t *testing.T) {
	reader := strings.NewReader("foo\nbar")
	check := NewSpellCheck(reader)

	assert.NotNil(t, check)
	assert.Equal(t, 100.0, check.CalculateCorrectness("foo"))
}

func TestCheckSpelling_CompleteFail(t *testing.T) {
	reader := strings.NewReader("foo")
	check := NewSpellCheck(reader)

	assert.NotNil(t, check)
	assert.Equal(t, 0.0, check.CalculateCorrectness("xyz"))
}

func TestCheckSpelling_PartialFail(t *testing.T) {
	reader := strings.NewReader("foo")
	check := NewSpellCheck(reader)

	assert.NotNil(t, check)
	assert.Equal(t, 50.0, check.CalculateCorrectness("fof"))
}

func TestCheckSpelling_DictBased(t *testing.T) {
	file, err := os.Open("../190intermediate/enable1.txt")
	assert.Nil(t, err)
	defer file.Close()

	check := NewSpellCheck(file)

	assert.NotNil(t, check)
	assert.True(t, check.CalculateCorrectness("zookeeper") < 0.001)
	assert.True(t, check.CalculateCorrectness("asdf") < 0.001)
}