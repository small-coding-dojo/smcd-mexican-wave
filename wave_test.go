package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestWaveEmpty(t *testing.T) {
	assert.Equal(t, []string{}, wave(""))
}

func TestSingleLetter(t *testing.T) {
	assert.Equal(t, []string{"A"}, wave("a"))
}

func wave(s string) []string {
	return []string {strings.ToUpper(s)}
}