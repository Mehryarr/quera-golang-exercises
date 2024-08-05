package main

type Qutex struct {
	locker chan int
}

func NewQutex() *Qutex {
	return &Qutex{
		locker: make(chan int, 1),
	}
}

func (q *Qutex) Lock() {
	q.locker <- 1
}

func (q *Qutex) Unlock() {
	select {
	case <-q.locker:
	default:
		panic("channel is already Unlocked")
	}
}
