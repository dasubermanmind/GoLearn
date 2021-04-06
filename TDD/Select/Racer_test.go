package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/*
func TestRacer(t *testing.T){
	slowUrl := "http://www.facebook.com"
	fastUrl := "http://www.quii.co.uk"

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}

 */

func TestRacer(t *testing.T){
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want{
		t.Errorf("Got %q want %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

t.Run("returns an error if a server doesnt response within 10s", func(t *testing.T){
	serverA := makeDelayedServer(11*time.Second)
	serverB := makeDelayedServer(12*time.Second)

	defer serverA.Close()
	defer serverB.Close()

	_, err := Racer(serverA.URL, serverB.URL)

	if err != nil {
		t.Errorf("expected and error but didnt get one")
	}
})