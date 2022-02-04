package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		ServerA := makeDelayedServer(20 * time.Millisecond)
		ServerB := makeDelayedServer(10 * time.Millisecond)

		defer ServerA.Close()
		defer ServerB.Close()

		_, err := Racer(ServerA.URL, ServerB.URL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
}
