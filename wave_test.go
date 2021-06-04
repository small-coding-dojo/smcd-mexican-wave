package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWaveEmpty(t *testing.T) {
	assert.Equal(t, []string{}, wave(""))
}

func TestSingleLetter(t *testing.T) {
	assert.Equal(t, []string{"A"}, wave("a"))
}

func TestTwoCharacters(t *testing.T) {
	assert.Equal(t, []string{"Ab", "aB"}, wave("ab"))
}

func TestThreeCharacters(t *testing.T) {
	assert.Equal(t, []string{"Abc", "aBc", "abC"}, wave("abc"))
}

func TestFiveCharacters(t *testing.T) {
	assert.Equal(t, []string{"Hello", "hEllo", "heLlo", "helLo", "hellO"}, wave("hello"))
}

func TestSixCharactersIncludingSpace(t *testing.T) {
	assert.Equal(t, []string{"Hell o", "hEll o", "heLl o", "helL o", "hell O"}, wave("hell o"))
}

func BenchmarkWave(b *testing.B) {
	for i:=0;i<b.N;i++{
		wave("lorem ipsum")
	}
}

func wave(s string) []string {
	wave := []string{}

	for i:=0; i<len(s);i++ {
		if s[i] == ' ' {
			continue
		}
		
		sb := []byte(s)
		sb[i] = sb[i] + 'A' - 'a'
		wave = append(wave, string(sb))

		//too slow
		//wave = append(wave, s[:i] + strings.ToUpper(string(s[i])) + s[i+1:] )
	}
	return wave
}