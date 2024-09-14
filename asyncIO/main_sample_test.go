package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func simpleTask() string {
	time.Sleep(1 * time.Second)
	return "result"
}

func TestSimple(t *testing.T) {
	fResult := Async(simpleTask)
	assert.False(t, fResult.Done.Load())
	result := fResult.Await()
	assert.Equal(t, "result", result)
	assert.True(t, fResult.Done.Load())
}

func handleTimeout(d time.Duration, t *testing.T) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		if elapsed > d {
			t.FailNow()
		}
	}
}

func TestMultiple(t *testing.T) {
	defer handleTimeout(1100*time.Millisecond, t)()

	fResult1 := Async(simpleTask)
	fResult2 := Async(simpleTask)

	res1 := fResult1.Await()
	res2 := fResult2.Await()

	assert.Equal(t, "result", res1)
	assert.Equal(t, "result", res2)

	// they should run concurrently and done in about a second
}

func TestCombine(t *testing.T) {
	fResult1 := Async(simpleTask)
	fResult2 := Async(simpleTask)

	combinedFResult := CombineFutureResults(fResult1, fResult2)

	// first item
	select {
	case <-time.After(1100 * time.Millisecond):
		t.FailNow()

	case res := <-combinedFResult.ResultChan:
		assert.Equal(t, "result", res)
	}

	// second item should be availble fast
	select {
	case <-time.After(100 * time.Millisecond):
		t.FailNow()

	case res := <-combinedFResult.ResultChan:
		assert.Equal(t, "result", res)
	}
}

func TestTimeout(t *testing.T) {
	fResult := AsyncWithTimeout(simpleTask, 700*time.Millisecond)

	select {
	case <-time.After(800 * time.Millisecond):
		t.FailNow()

	case res := <-fResult.ResultChan:
		assert.Equal(t, "timeout", res) // timeout is reached before 800ms
	}
}
