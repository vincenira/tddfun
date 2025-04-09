package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

/*
We know we want our Countdown function to write data somewhere and io.Writer is the de-facto
way of capturing that as an interface in Go.
- In main we will send to os.Stdout so our users see the countdown printed to the terminal.
- In test we will send to bytes.Buffer so our tests can capture what data is being generated.
*/
func TestCountdown(t *testing.T) {
	t.Run("Prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

/*
Our tests take 3 seconds to run.
- Every forward-thinking post about software development emphasises the importance of
  quick feedback loops.
- Slow tests ruin developer productivity.
- Imagine if the requirements get more sophisticated warranting more tests.
 Are we happy with 3s added to the test run for every new test of Countdown?

We have not tested an important property of our function
We have a dependency on Sleeping which we need to extract so we can then control it
in our tests.
If we can mock time.Sleep we can use dependency injection to use it instead of a
"real" time.Sleep and then we can spy on the calls to make assertions on them.
*/
