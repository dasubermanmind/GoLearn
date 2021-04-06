package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string,error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	}
}
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bduration := time.Since(startB)

	if aDuration < bduration{
		return a
	}
	return b
}

func ping(url string)chan struct{}{
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}