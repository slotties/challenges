package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func assertLetter(t *testing.T, expected string, input string) {
	output, err := EnglishBraille26Letter(input)
	assert.Nil(t, err)
	assert.Equal(t, expected, output)
}

func TestEnglishBraille26Letter(t *testing.T) {
	// http://en.wikipedia.org/wiki/English_Braille#Alphabet
	
	
	assertLetter(t, "a", "O.....")
	assertLetter(t, "b", "O.O...")
	assertLetter(t, "c", "OO....")
	assertLetter(t, "d", "OO.O..")
	assertLetter(t, "e", "O..O..")
	assertLetter(t, "f", "OOO...")
	assertLetter(t, "g", "OOOO..")
	assertLetter(t, "h", "O.OO..")
	assertLetter(t, "i", ".OO...")
	assertLetter(t, "j", ".OOO..")
	assertLetter(t, "k", "O...O.")
	assertLetter(t, "l", "O.O.O.")
	assertLetter(t, "m", "OO..O.")
	assertLetter(t, "n", "OO.OO.")
	assertLetter(t, "o", "O..OO.")
	assertLetter(t, "p", "OOO.O.")
	assertLetter(t, "q", "OOOOO.")
	assertLetter(t, "r", "O.OOO.")
	assertLetter(t, "s", ".OO.O.")
	assertLetter(t, "t", ".OOOO.")
	assertLetter(t, "u", "O...OO")
	assertLetter(t, "v", "O.O.OO")
	assertLetter(t, "w", ".OOO.O")
	assertLetter(t, "x", "OO..OO")
	assertLetter(t, "y", "OO.OOO")
	assertLetter(t, "z", "O..OOO")
}

func TestEnglishBraille26Letter_BadInput(t *testing.T) {
	// Does not match any letter.
	_, error := EnglishBraille26Letter("xxxxxx")
	assert.NotNil(t, error)
	
	// One char too much.
	_, error = EnglishBraille26Letter("OO.....")
	assert.NotNil(t, error)
}

func TestEnglishBraille26Word_HelloWorld(t *testing.T) {
	input := 
			"O. O. O. O. O. .O O. O. O. OO\n" +
			"OO .O O. O. .O OO .O OO O. .O\n" +
			".. .. O. O. O. .O O. O. O. .."
			
	output := EnglishBraille26Word(input)
	
	assert.Equal(t, "helloworld", output)
}
