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
// make the timeout Configurable to facilitate test since 10 seconds can be longer.
// Since we don't care about the timeout in happy test, we can be strategic in our implementation
var tenSecondTimeout = 10 * time.Second

func ConcRacer(a, b string) (winner string, error error) {
	return ConfigurableConcRacer(a, b, tenSecondTimeout)
}

func ConfigurableConcRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out waiting for %s and %s", a, b)
	}
}

/*
important:
select:
- Helps you wait on multiple channels
- Sometimes you'll want to include time.After in one of your cases to prevent your system
  blocking forever
httptest:
- Convenient way of creating test servers so you can have reliable and controllable tests.
- uses the same interfaces as the "real" net/http servers which is consistent and less for you
to learn
*/
