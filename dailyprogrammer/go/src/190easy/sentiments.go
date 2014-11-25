package main

import (
	"strings"
)

/*
	Specific part of Challenge:
	Check for trigger words.
*/

// TODO: implement word boosting
// TODO: implement simple word normalization/stemming

type RankedWord struct {
	word string
	score int
}
type SentimentalAnalyzer struct {
	words []RankedWord
}

func New() (*SentimentalAnalyzer) {
	analyzer := new(SentimentalAnalyzer)
	analyzer.words = make([]RankedWord, 0)

	return analyzer
}

/**
 * Returns a positive value for a seemingly happy text and a negative
 * value for a seemingly sad text. Both values go to their extreme the
 * more extreme the text seems to be.
 */
func (self *SentimentalAnalyzer) Score(text string) int {
	score := 0
	for _, rankedWord := range self.words {
		count := strings.Count(text, rankedWord.word)
		score += (count * rankedWord.score)
	}

	return score
}

func (self *SentimentalAnalyzer) Add(word string, score int) {
	self.words = append(self.words, RankedWord{ word, score })
}