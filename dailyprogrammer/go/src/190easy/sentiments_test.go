package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPureHappy(t *testing.T) {
	s := New()
	s.Add("happy", 1)
	s.Add("love", 1)
	s.Add("amazing", 1)
	score := s.Score("I am so happy. I really love to see you!\n That's so amazing!")

	assert.Equal(t, 3, score)
}

func TestPureSad(t *testing.T) {
	s := New()
	s.Add("hate", -1)
	s.Add("awful", -1)
	s.Add("worst", -1)
	score := s.Score("Gosh, I hate you!\n You're so awful!\n Can't you just die? You're the worst!")

	assert.Equal(t, -3, score)
}

func TestNeutral(t *testing.T) {
	s := New()
	score := s.Score("I don't trigger anything. LOL!")

	assert.Equal(t, 0, score)
}

func TestMixed(t *testing.T) {
	s := New()
	s.Add("happy", 1)
	s.Add("love", 1)
	s.Add("worst", -1)
	score := s.Score("Gosh, I'm so happy you're alive. You know how much I love you. You're the worst.")
	assert.Equal(t, 1, score)
}

func TestOverlappingWords(t *testing.T) {
	s := New()
	s.Add("love", 1)
	s.Add("loved", 1)
	score := s.Score("I loved you!")
	assert.Equal(t, 1, score)
}
