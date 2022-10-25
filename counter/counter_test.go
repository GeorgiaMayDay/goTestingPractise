package counter

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()
	if counter.Current() != want {
		t.Errorf("got %d want %d", counter.Current(), want)
	}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
	})

	t.Run("test concurrency", func(t *testing.T) {
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
		//Means all of the concurrent functions will have
		// fired off by the time we get here
		wg.Wait()

		assertCounter(t, counter, wantedCount)

	})
}
