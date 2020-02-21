package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := startServer(20 * time.Millisecond)
	fastServer := startServer(0)

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}

func startServer(responseDelay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(responseDelay)
		w.WriteHeader(http.StatusOK)
	}))
}
