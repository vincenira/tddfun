package fun11

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)

	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL

	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestConcRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)

		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL

		got, _ := ConcRacer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond with 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Second)

		defer server.Close()

		_, err := ConfigurableConcRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}
