package tests

import (
	c "github.com/diego-lutke/golang-playground/challenges"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoveFunc(t *testing.T) {
	assert.Equal(t, true, c.LoveFunc(1, 4))
	assert.Equal(t, false, c.LoveFunc(2, 2))
	assert.Equal(t, true, c.LoveFunc(0, 1))
	assert.Equal(t, false, c.LoveFunc(0, 0))
}
