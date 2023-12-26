package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchInt(t *testing.T) {
	testdata := []int{1, 3, 7, 9, 10, 13, 107}
	searchedElem := 7
	expected := 2
	res, isFound := BinarySearch(testdata, searchedElem)

	assert.True(t, isFound)
	assert.Equal(t, expected, res)
}

func TestBinarySearchFloat(t *testing.T) {
	testdata := []float64{1, 3, 7.2, 9, 10, 13, 107.5}
	searchedElem := 107.5
	expected := 6
	res, isFound := BinarySearch(testdata, searchedElem)

	assert.True(t, isFound)
	assert.Equal(t, expected, res)
}

func TestBinarySearchNotFound(t *testing.T) {
	testdata := []int{1, 3, 7, 9, 10, 13, 107}
	searchedElem := 0
	expected := 0
	res, isFound := BinarySearch(testdata, searchedElem)

	assert.False(t, isFound)
	assert.Equal(t, expected, res)
}
