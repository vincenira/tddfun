package fun13

import (
	"sync"
	"testing"
)

// API to give us a method to increment the counter and then retrieve its value
func TestCounter(t *testing.T) {
	t.Run("Incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
	/*
	   A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls
	   Add to set the number of goroutines to wait for. Then each of the goroutines runs and
	   calls Done when finished. At the same time, Wait can be used to block until all goroutines
	   have finished.
	*/
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()
		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, wantedCount)
	})
}

func NewCounter() *Counter {
	return &Counter{}
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
