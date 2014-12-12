package main

import (
	"io"
	"bufio"
)

type SpellCheck struct {
	matrix [][]float32
}

func NewSpellCheck(reader io.Reader) (*SpellCheck) {
	matrix, totals := createMarkovMatrix(reader)	
	
	check := new(SpellCheck)
	check.matrix = calculatePropabilities(matrix, totals)

	return check
}

func newMarkovMatrix() [][]float32 {
	matrix := make([][]float32, 26)
	for i := 0; i < 26; i++ {
		matrix[i] = make([]float32, 26)
	}

	return matrix
}

func createMarkovMatrix(reader io.Reader) ([][]float32, []int) {
	totals := make([]int, 27)
	matrix := newMarkovMatrix()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		word := scanner.Text()
		if (len(word) > 2) {
			for idx, currentRune := range word[:len(word) - 1] {
				nextRune := word[idx + 1]
				matrix[currentRune - 'a'][nextRune - 'a']++
				totals[currentRune - 'a']++
			}
		}
	}

	return matrix, totals
}

func calculatePropabilities(matrix [][]float32, totals []int) [][]float32 {
	for idx, nextRuneList := range matrix {
		for nextRuneIdx, _ := range nextRuneList {
			if totals[idx] > 0 {
				nextRuneList[nextRuneIdx] = nextRuneList[nextRuneIdx] / float32(totals[idx])
			}
		}
	}

	return matrix
}

func (self *SpellCheck) CalculateCorrectness(word string) float32 {
	correctness := float32(1.0)
	foundCombinations := 0
	for idx, currentRune := range word[:len(word) - 1] {
		nextRune := word[idx + 1]
		probability := self.matrix[currentRune - 'a'][nextRune - 'a']
		if (probability > 0.0) {
			correctness *= probability
			foundCombinations++
		}
	}

	if (foundCombinations == 0) {
		return 0.0
	}

	correctness *= (float32(foundCombinations) / float32(len(word) - 1))
	return correctness * 100.0
}
