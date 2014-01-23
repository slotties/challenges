package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEnglishBraille26Letter(t *testing.T) {
	// http://en.wikipedia.org/wiki/English_Braille#Alphabet
	assert.Equal(t, "a", EnglishBraille26Letter("O....."))
	assert.Equal(t, "b", EnglishBraille26Letter("O.O..."))
	assert.Equal(t, "c", EnglishBraille26Letter("OO...."))
	assert.Equal(t, "d", EnglishBraille26Letter("OO.O.."))
	assert.Equal(t, "e", EnglishBraille26Letter("O..O.."))
	assert.Equal(t, "f", EnglishBraille26Letter("OOO..."))
	assert.Equal(t, "g", EnglishBraille26Letter("OOOO.."))
	assert.Equal(t, "h", EnglishBraille26Letter("O.OO.."))
	assert.Equal(t, "i", EnglishBraille26Letter(".OO..."))
	assert.Equal(t, "j", EnglishBraille26Letter(".OOO.."))
	assert.Equal(t, "k", EnglishBraille26Letter("O...O."))
	assert.Equal(t, "l", EnglishBraille26Letter("O.O.O."))
	assert.Equal(t, "m", EnglishBraille26Letter("OO..O."))
	assert.Equal(t, "n", EnglishBraille26Letter("OO.OO."))
	assert.Equal(t, "o", EnglishBraille26Letter("O..OO."))
	assert.Equal(t, "p", EnglishBraille26Letter("OOO.O."))
	assert.Equal(t, "q", EnglishBraille26Letter("OOOOO."))
	assert.Equal(t, "r", EnglishBraille26Letter("O.OOO."))
	assert.Equal(t, "s", EnglishBraille26Letter(".OO.O."))
	assert.Equal(t, "t", EnglishBraille26Letter(".OOOO."))
	assert.Equal(t, "u", EnglishBraille26Letter("O...OO"))
	assert.Equal(t, "v", EnglishBraille26Letter("O.O.OO"))
	assert.Equal(t, "w", EnglishBraille26Letter(".OOO.O"))
	assert.Equal(t, "x", EnglishBraille26Letter("OO..OO"))
	assert.Equal(t, "y", EnglishBraille26Letter("OO.OOO"))
	assert.Equal(t, "z", EnglishBraille26Letter("O..OOO"))
}

func TestEnglishBraille26Word_HelloWorld(t *testing.T) {
	input := 
			"O. O. O. O. O. .O O. O. O. OO\n" +
			"OO .O O. O. .O OO .O OO O. .O\n" +
			".. .. O. O. O. .O O. O. O. .."
			
	output := EnglishBraille26Word(input)
	
	assert.Equal(t, "helloworld", output)
}
