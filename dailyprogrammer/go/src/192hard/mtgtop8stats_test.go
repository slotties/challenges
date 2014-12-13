package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadLatestMajorEvents(t *testing.T) {
	events, err := LoadLatestMajorEvents("ST")

	assert.Nil(t, err)
	assert.True(t, len(events) > 0)
}

func TestLoadEventDecks(t *testing.T) {
	decks, err := LoadEventDecks("8689", "ST")

	assert.Nil(t, err)
	assert.Equal(t, 24, len(decks))
	assert.Equal(t, "Abzan Midrange", decks["249654"])
	assert.Equal(t, "Sultai Whip", decks["249642"])
}

func TestLoadCards(t *testing.T) {
	cards, err := LoadCards("8689", "ST", "249654")

	assert.Nil(t, err)
	assert.Equal(t, 21, len(cards))
	assert.Equal(t, 3, cards["Forest"])
	assert.Equal(t, 4, cards["Siege Rhino"])
}

func TestGetMajorEventCardStatistics(t *testing.T) {
	stats, err := getMajorEventCardStatistics("ST")

	assert.Nil(t, err)
	assert.True(t, stats["Forest"] > 0)
}