package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("quicker server wins", func(t *testing.T) {
		slowServer := startServer(20 * time.Millisecond)
		fastServer := startServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := startServer(11 * time.Second)
		serverB := startServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func startServer(responseDelay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(responseDelay)
		w.WriteHeader(http.StatusOK)
	}))
}
