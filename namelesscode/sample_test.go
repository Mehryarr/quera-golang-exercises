package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	testSlice := []int{1, 2, 3}
	AddElement(&testSlice, 4)
	assert.EqualValues(t, []int{1, 2, 3, 4}, testSlice)
	min := FindMin(&testSlice)
	assert.Equal(t, min, 1)
	ReverseSlice(&testSlice)
	assert.EqualValues(t, []int{4, 3, 2, 1}, testSlice)
}
