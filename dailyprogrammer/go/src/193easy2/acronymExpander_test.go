package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	acronyms := map[string]string {
		"lol": "laugh out loud",
		"dw": "don't worry",
		"hf": "have fun",
		"gg": "good game",
		"brb": "be right back",
		"g2g": "got to go",
		"wtf": "what the fuck",
		"wp": "well played",
		"gl": "good luck",
		"imo": "in my opinion",
	}

	text := ExpandAcronyms("imo that was wp. Anyway I've g2g", acronyms)
	assert.Equal(t, "in my opinion that was well played. Anyway I've got to go", text)

	text = ExpandAcronyms("wtf that was unfair", acronyms)
	assert.Equal(t, "what the fuck that was unfair", text)

	text = ExpandAcronyms("gl all hf", acronyms)
	assert.Equal(t, "good luck all have fun", text)	
}