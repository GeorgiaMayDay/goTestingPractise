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

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
