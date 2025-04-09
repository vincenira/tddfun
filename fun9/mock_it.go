package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countDownStart = 3
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const (
	write = "write"
	sleep = "sleep"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

/*
We can make the code more readable with an iterator custom
Function signature of an custom iterator
To write an iterator like countDownFrom, you need to write a function in a particular way. From the docs:

The “range” clause in a “for-range” loop now accepts iterator functions of
the following types

	func(func() bool)
	func(func(K) bool)
	func(func(K, V) bool)
*/
func CountdownRange(out io.Writer, sleeper Sleeper) {
	for i := range countDownFrom(countDownStart) {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	/*
	   let's use the configurableSleeper
	   	sleeper := &DefaultSleeper{}
	   	Countdown(os.Stdout, sleeper)
	*/
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
	fmt.Println("Second countDown")
	CountdownRange(os.Stdout, sleeper)
}
