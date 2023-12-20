package wordcounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCounterSimple(t *testing.T) {
	testdata := "i love my mom very very much"
	expected := map[string]int{
		"love": 1, "i": 1, "my": 1, "very": 2, "mom": 1, "much": 1,
	}
	res := CountWords(testdata)

	assert.Equal(t, expected, res)
}

func TestWordCounterDifferentRegisters(t *testing.T) {
	testdata := "I lOve lovE my mom VeRy very much"
	expected := map[string]int{
		"love": 2, "i": 1, "my": 1, "very": 2, "mom": 1, "much": 1,
	}
	res := CountWords(testdata)

	assert.Equal(t, expected, res)
}

func TestWordCounterSingleWord(t *testing.T) {
	testdata := " hello   \n\t"
	expected := map[string]int{
		"hello": 1,
	}
	res := CountWords(testdata)

	assert.Equal(t, expected, res)
}

func TestWordCounterMultiLine(t *testing.T) {
	testdata := " hello   \n\t \t\t hello"
	expected := map[string]int{
		"hello": 2,
	}
	res := CountWords(testdata)

	assert.Equal(t, expected, res)
}

func TestWordCounterEmptyString(t *testing.T) {
	testdata := ""
	expected := map[string]int{}
	res := CountWords(testdata)

	assert.Equal(t, expected, res)

	testdata = "\t\t\t\n   \t,,,, "
	expected = map[string]int{}
	res = CountWords(testdata)

	assert.Equal(t, expected, res)
}
