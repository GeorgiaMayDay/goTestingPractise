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

func Racer(linkA string, linkB string) (winner string) {
	aDuration := measureResponseTime(linkA)

	bDuration := measureResponseTime(linkB)

	if bDuration < aDuration {
		return linkB
	} else {
		return linkA
	}
}

func betterRacer(linkA string, linkB string) (winner string) {
	select {
	case <-ping(linkA):
		return linkA
	case <-ping(linkB):
		return linkB
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
