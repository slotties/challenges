package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBruteforceGuess1(t *testing.T) {
	guess := NewBruteforceLanguageGuess("F:\\Download\\dictionaries")
	languageOrder := guess.Guess("l'école a été classé meilleure école de cinéma d'europe par la revue professionnelle de référence the hollywood reporter et 7e meilleure école de cinéma du monde juste derrière le california institute of the arts et devant l'université columbia")
	
	assert.Equal(t, 2, len(languageOrder))
	assert.Equal(t, "fr.dic", languageOrder[0].Language)
	assert.Equal(t, "en.dic", languageOrder[1].Language)
}

func TestBruteforceGuess2(t *testing.T) {
	guess := NewBruteforceLanguageGuess("F:\\Download\\dictionaries")
	languageOrder := guess.Guess("few things are harder to put up with than the annoyance of a good example")
	
	assert.Equal(t, 1, len(languageOrder))
	assert.Equal(t, "en.dic", languageOrder[0].Language)
}