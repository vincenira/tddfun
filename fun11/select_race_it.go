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

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)

	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}
