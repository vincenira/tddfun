/*
make a program that races two websites by hitting with HTTP GET
and returning the URL which return first. if none of them return
with 10 seconds then it should return an error
*/
package fun11

import (
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)

	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func ConcRacer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}
