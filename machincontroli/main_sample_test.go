package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	got := NewCar(20, 80)
	assert.Equal(t, GetSpeed(got), 20)
	assert.Equal(t, GetBattery(got), 80)
	ChargeCar(got, 4)
	assert.Equal(t, GetBattery(got), 82)
	s := TryFinish(got, 4)
	assert.Equal(t, GetBattery(got), 80)
	assert.Equal(t, s, "0.20")
}
