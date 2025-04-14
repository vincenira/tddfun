/*
make a program that races two websites by hitting with HTTP GET
and returning the URL which return first. if none of them return
with 10 seconds then it should return an error
*/
package fun11

import (
	"fmt"
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

/*
You'll recall from the concurrency chapter that you can wait for values to be sent
to a channel with myVar := <-ch. This is a blocking call, as you're waiting for a value.

select allows you to wait on multiple channels. The first one to send a value "wins"
and the code underneath the case is execute
*/
func ConcRacer(a, b string) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("time out waiting for %s and %s", a, b)
	}
}
