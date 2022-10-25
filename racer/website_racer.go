package racer

import (
	"errors"
	"net/http"
	"time"
)

var ErrTakesTooLong = errors.New("websites took too long to respond")

func measureResponseTime(link string) time.Duration {
	start := time.Now()
	http.Get(link)
	duration := time.Since(start)
	return duration
}

func Racer(linkA string, linkB string) (winner string, err error) {
	aDuration := measureResponseTime(linkA)

	bDuration := measureResponseTime(linkB)

	if bDuration < aDuration {
		return linkB, nil
	} else {
		return linkA, nil
	}
}

func BetterRacer(linkA string, linkB string) (winner string, err error) {
	return TestableRacer(linkA, linkB, time.Duration(10*time.Second))
}

func TestableRacer(linkA string, linkB string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(linkA):
		return linkA, nil
	case <-ping(linkB):
		return linkB, nil
	case <-time.After(timeout):
		return "Error", ErrTakesTooLong
	}
}

// This function is used to signal when the request has processed
func ping(url string) chan struct{} {
	done := make(chan struct{})
	go func() {
		http.Get(url)
		close(done)
	}()
	return done
}
