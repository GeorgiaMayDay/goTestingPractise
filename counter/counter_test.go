package counter

import "testing"

func assertCounter(t testing.TB, counter Counter, want int) {
	t.Helper()
	if counter.Current() != want {
		t.Errorf("got %d want %d", counter.Current(), want)
	}
}

func TestCounter(t *testing.T) {
	counter := Counter{}
	counter.Inc()
	counter.Inc()
	counter.Inc()
}
