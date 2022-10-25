package racer

import (
	"net/http"
	"time"
)

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

func betterRacer(linkA string, linkB string) (winner string, err error) {
	select {
	case <-ping(linkA):
		return linkA, nil
	case <-ping(linkB):
		return linkB, nil
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
