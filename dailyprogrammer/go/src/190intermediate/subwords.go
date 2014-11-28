package main

import (
	"io"
	"bufio"
)

/*
	Challenge:
		Given the wordlist enable1.txt, you must find the word in that file which also contains the greatest number of words within that word.
		For example, the word 'grayson' has the following words in it
		Grayson
		Gray
		Grays
		Ray
		Rays
		Son
		On
		Here's another example, the word 'reports' has the following
		reports
		report
		port
		ports
		rep
		You're tasked with finding the word in that file that contains the most words.
		NOTE : If you have a different wordlist you would like to use, you're free to do so.
*/

func GetWordWithMostSubwords(reader io.Reader) string {
	words := loadWords(reader)

	var biggestWord string
	biggestSubwordCount := 0

	for word, _ := range words {
		subwordCount := countSubwords(word, words)

		if subwordCount > biggestSubwordCount {
			biggestWord = word
			biggestSubwordCount = subwordCount
		}
	}

	return biggestWord
}

func countSubwords(word string, words map[string]bool) int {
	wordLength := len(word)
	subwordCount := 0

	for i := 0; i < wordLength - 1; i++ {
		for j := i + 1; j < wordLength; j++ {
			subword := word[i:j+1]
			if words[subword] {
				subwordCount++
			}
		}
	}

	return subwordCount
} 

func loadWords(reader io.Reader) map[string]bool {
	scanner := bufio.NewScanner(reader)

	words := make(map[string]bool)

	for scanner.Scan() {
		word := scanner.Text()
		words[word] = true
	}	

	return words
}