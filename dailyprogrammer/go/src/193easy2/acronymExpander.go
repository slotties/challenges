package main

import (
	"regexp"
)

/*
	Challenge: http://www.reddit.com/r/dailyprogrammer/comments/2ptrmp/20141219_challenge_193_easy_acronym_expander/
	During online gaming (or any video game that requires teamwork) , there is often times that you need to speak to your teammates. Given the nature of the game, it may be inconvenient to say full sentences and it's for this reason that a lot of games have acronyms in place of sentences that are regularly said.

	Example
		gg : expands to 'Good Game'
		brb : expands to 'be right back'
		and so on...

	This is even evident on IRC's and other chat systems.
	However, all this abbreviated text can be confusing and intimidating for someone new to a game. They're not going to instantly know what 'gl hf all'(good luck have fun all) means. It is with this problem that you come in.
	You are tasked with converting an abbreviated sentence into its full version.
*/
func ExpandAcronyms(text string, acronyms map[string]string) string {
	textBytes := []byte(text)
	wordsRe := regexp.MustCompile("\\w*")

	result := wordsRe.ReplaceAllFunc(textBytes, func(word []byte) []byte {
		expandedWord, isAcronym := acronyms[string(word)]
		if isAcronym {
			return []byte(expandedWord)
		} else {
			return word
		}
	})

	return string(result)
}