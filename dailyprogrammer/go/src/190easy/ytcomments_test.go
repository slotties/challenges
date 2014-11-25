package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadFromExistingUrl(t *testing.T) {
	comments := LoadComments("https://www.youtube.com/watch?v=BOByH_iOn88")

	assert.NotNil(t, comments)
	assert.True(t, len(comments) > 0)
}

func TestLoadFromNotExistingUrl(t *testing.T) {
	comments := LoadComments("https://www.youtube.com/watch?v=BOByH_iOn88xyz")

	assert.NotNil(t, comments)
	assert.True(t, len(comments) == 0)
}