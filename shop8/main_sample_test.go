package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	s := NewStore()
	err := s.AddProduct("apple", 20000, 10)
	assert.Nil(t, err)
	c, err2 := s.GetProductCount("apple")
	assert.Nil(t, err2)
	assert.Equal(t, 10, c)
	p, err3 := s.GetProductPrice("apple")
	assert.Nil(t, err3)
	assert.Equal(t, 20000.0, p)
	err = s.Order("apple", 8)
	assert.Nil(t, err)
	c, _ = s.GetProductCount("apple")
	assert.Equal(t, 2, c)
}
