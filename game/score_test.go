package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewScoreSetsScoreToZero(t *testing.T) {
	assert.Equal(t, int64(0), NewScore().Get())
}

func TestAddAddsTheCorrectAmount(t *testing.T) {
	s := NewScore()
	s.Add(2)
	assert.Equal(t, int64(2), s.Get())
	s.Add(32)
	assert.Equal(t, int64(34), s.Get())
}
