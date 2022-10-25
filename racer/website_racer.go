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
