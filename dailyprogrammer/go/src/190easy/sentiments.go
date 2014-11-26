package main

import (
	"regexp"
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
	wordCounts := countWords(getWords(text))

	for _, rankedWord := range self.words {
		count := getWordCount(rankedWord.word, wordCounts)
		score += (count * rankedWord.score)
	}

	return score
}

func (self *SentimentalAnalyzer) Add(word string, score int) {
	self.words = append(self.words, RankedWord{ word, score })
}

func getWordCount(searchWord string, wordCounts map[string]int) int {
	for word, count := range wordCounts {
		if word == searchWord {
			return count
		}
	}

	return 0
}

func getWords(text string) []string {
    words := regexp.MustCompile("\\w+")
    return words.FindAllString(text, -1)
}

func countWords(words []string) map[string]int {
    counts := make(map[string]int)
    for _, word := range words {
        counts[word]++
    }
    return counts
}