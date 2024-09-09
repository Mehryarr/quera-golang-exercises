package main

import (
	"sync/atomic"
	"time"
)

type FutureResult struct {
	Done       atomic.Bool
	ResultChan chan string
	Result     string
	doneChan   chan struct{}
}

type Task func() string

func Async(t Task) *FutureResult {
	fResult := &FutureResult{
		ResultChan: make(chan string, 1),
		doneChan:   make(chan struct{}),
	}

	go func() {
		result := t()
		fResult.ResultChan <- result
		fResult.Result = result
		fResult.Done.Store(true)
		close(fResult.doneChan)
	}()

	return fResult
}
func AsyncWithTimeout(t Task, timeout time.Duration) *FutureResult {
	fResult := &FutureResult{
		ResultChan: make(chan string, 1),
		doneChan:   make(chan struct{}),
	}

	go func() {
		select {
		case result := <-asyncTask(t):
			fResult.ResultChan <- result
			fResult.Result = result
			fResult.Done.Store(true)
			close(fResult.doneChan)
		case <-time.After(timeout):
			fResult.ResultChan <- "timeout"
			fResult.Result = "timeout"
			fResult.Done.Store(true)
			close(fResult.doneChan)
		}
	}()

	return fResult
}

func asyncTask(t Task) <-chan string {
	resultChan := make(chan string, 1)
	go func() {
		resultChan <- t()
	}()
	return resultChan
}

func (fResult *FutureResult) Await() string {
	<-fResult.doneChan
	return fResult.Result
}

func CombineFutureResults(fResults ...*FutureResult) *FutureResult {
	combinedResult := &FutureResult{
		ResultChan: make(chan string, len(fResults)),
		doneChan:   make(chan struct{}),
	}

	go func() {
		for _, fResult := range fResults {
			// Send each result individually as they complete
			res := fResult.Await()
			combinedResult.ResultChan <- res
		}

		combinedResult.Done.Store(true)
		close(combinedResult.doneChan)
		close(combinedResult.ResultChan) // Close the channel once all results are sent
	}()

	return combinedResult
}
