package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeMeAServer(delay time.Duration) (server *httptest.Server) {
	newServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return newServer
}
func TestRacer(t *testing.T) {
	slowServer := makeMeAServer(20 * time.Millisecond)
	fastServer := makeMeAServer(2 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	t.Run("Racer no concurrency", func(t *testing.T) {

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if err != nil {
			t.Errorf("got error, want no error")
		}
	})

	t.Run("Racer with concurrency", func(t *testing.T) {
		want := fastURL
		got, err := BetterRacer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if err != nil {
			t.Errorf("got error, want no error")
		}
	})

	t.Run("Racer returns error if both servers take 10 sec", func(t *testing.T) {
		tooSlowServer := makeMeAServer(24 * time.Millisecond)

		slowestURL := tooSlowServer.URL

		_, err := TestableRacer(slowestURL, slowestURL, time.Duration(2*time.Millisecond))

		if err == nil {
			t.Errorf("got no error, want error")
		}

		tooSlowServer.Close()
	})

}
