package main

import "errors"

/*
	A parser/mapper of braille letters: 
	http://en.wikipedia.org/wiki/English_Braille#Alphabet
*/

/*
	This function parses a single letter from the english alphabet (the common 26 letters A to Z).
	The input format is a single line, 6 character string. The original braille is simply collapsed into a single line.
	Every dark dot is a O. Every white dot is a . (dot).
	
	Example:
	O.
	..
	..
	=> O..... => A
	
	O.
	O.
	..
	=> O.O... => B
	
	O.
	.O
	OO
	=> O..OOO => Z
*/
func EnglishBraille26Letter(input string) (string, error) {
	if len(input) != 6 {
		return "", errors.New("The input string must have a length of 6 characters")
	}

	letterMask := 0
	
	for i, char := range input {
		if char == 'O' {
			letterMask = letterMask | (1 << uint(i))
		}
	}
	
	letterMapping := " a c bif e d hjg k m lsp o n rtq              w  u x v   z y"
	if (letterMask >= len(letterMapping)) {
		return "", errors.New("Unknown letter")
	}
	
	letter := letterMapping[letterMask]
	if letter == ' ' {
		return "", errors.New("Unknown letter")
	}
	
	return string(letterMapping[letterMask]), nil
}

/*
	Expects a whole braille word/line containing multiple letters, such as:
	OO O. O.
	O. .O .O
	.. O. O.
	=> foo
*/
func EnglishBraille26Word(input string) string {
	// 9 characters per braille letter:
	// 3*2 characters to represent a single letter
	// plus 3 characters spacing to the next letter (or newlines to end a line)
	letterCount := len(input) / 9
	secondLineOffset := (letterCount + 1) * 3
	thirdLineOffset := (letterCount + 1) * 6
	
	word := ""
	
	for idx := 0; idx <= (letterCount * 3); idx += 3 {
		letterTop := input[idx : idx + 2]
		letterMid := input[idx + secondLineOffset: idx + secondLineOffset + 2]
		letterBottom := input[idx + thirdLineOffset: idx + thirdLineOffset + 2]
		
		letter, error := EnglishBraille26Letter(letterTop + letterMid + letterBottom)
		if error == nil {
			word += letter
		}
	}
	
	return word
}