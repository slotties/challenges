package main

import (
	"os"
	"io/ioutil"
	"io"
	"bufio"
	"path/filepath"
	"strings"
	"sort"
)

/*
		Bruteforce guessing using maps and lookup hit counts.
 */
type WordList struct {
	dictionary map[string]bool
}
func (self *WordList) Add(word string) {
	self.dictionary[word] = true
}
func (self *WordList) Contains(word string) bool {
	_, found := self.dictionary[word]	
	return found
}

type LanguageRating struct {
	Language string
	Rating uint8
}
type LanguageRatingList []LanguageRating
func (l LanguageRatingList) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l LanguageRatingList) Len() int {
	return len(l)
}
func (l LanguageRatingList) Less(i int, j int) bool {
	return l[i].Rating > l[j].Rating
}
 
type BruteforceLanguageGuess struct {
	dictionaries map[string]WordList
}

func (self *BruteforceLanguageGuess) Guess(text string) LanguageRatingList {
	ratings := make(LanguageRatingList, 0)
	minHitRatio := 0.5
	
	for language, words := range self.dictionaries {
		rating, tokencount := rate(text, words)
		ratio := float64(rating) / float64(tokencount)
		
		if ratio >= minHitRatio {		
			ratings = append(ratings, LanguageRating{ Language: language, Rating: rating})
		}
	}
	
	sort.Sort(ratings)
	
	return ratings
}
func rate(text string, words WordList) (uint8, int) {
	count := uint8(0)
	tokens := strings.Split(text, " ")
	
	for _, word := range tokens {
		if words.Contains(word) {
			count++
		}
	}
	
	return count, len(tokens)
}

func NewBruteforceLanguageGuess(dictionaryFolder string) *BruteforceLanguageGuess {
	guess := BruteforceLanguageGuess{}
	guess.dictionaries = make(map[string]WordList)
	
	files, _ := ioutil.ReadDir(dictionaryFolder)
	for _, file := range files {
		dict, err := readDictionary(dictionaryFolder, file.Name())
		if err == nil {
			guess.dictionaries[file.Name()] = dict
		}
	}
	
	return &guess
}

func readDictionary(folder string, file string) (WordList, error) {
	words := WordList{}
	words.dictionary = make(map[string]bool)
	
	f, err := os.Open(filepath.Join(folder, file))
	if err != nil {
		return words, err
	}
	defer f.Close()
	
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return words, err
		}
		
		if line[0] != '%' && len(line) > 1 {
			words.Add(string(line))
		}
	}
	
	return words, nil
}