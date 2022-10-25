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
		got := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Racer with concurrency", func(t *testing.T) {
		want := fastURL
		got := betterRacer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
