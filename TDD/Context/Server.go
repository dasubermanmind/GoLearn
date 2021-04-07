package main

import (
	"context"
	"fmt"
	"net/http"
)

//type Store interface{
//	Fetch() string
//	Cancel()
//}

type Store interface {
	Fetch(ctx context.Context)(string error)
}

type StubStore struct{
	reponse string
}

func (s *StubStore) Fetch() string{
	return s.reponse
}

func Server(store Store) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		//context has a method Done()
		//which returns a channel which gets sent a signal when the context
		//is "done" or "cancelled". We want to listen to that signal and
		//call store.Cancel if we get it but we want to ignore it if our
		//Store manages to Fetch before it.
		ctx := r.Context()

		data := make(chan string, 1)

		go func(){
			data <- store.Fetch()
		}()

		select{
		case d:= <- data:
			fmt.Println(w,d)
		case <- ctx.Done():
			store.Cancel()
		}
	}
}