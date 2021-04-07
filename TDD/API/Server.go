package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

const jsonContentType = "application/json"


type InMemoryPlayerStore struct {
	store map[string]int
}

//type Animal interface{
//	Eater
//	Sleeper
//}

type Player struct{
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() [] Player
}

type PlayerServer struct {
	store PlayerStore
	router *http.ServeMux
}

func newPlayerServer(store PlayerStore) *PlayerServer {

	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	//p.Handler = router
	//p := &PlayerServer{
	//	store,
	//	http.NewServeMux(),
	//}
	//
	//p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	//p.router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	return p
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
	//router := http.NewServeMux()
	//
	//router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	//router.Handle("/players/", http.HandlerFunc(p.playerHandler))


	//router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	//	w.WriteHeader(http.StatusOK)
	//}))
	//
	//router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	//	player := strings.TrimPrefix(r.URL.Path, "/players/")
	//	switch r.Method {
	//	case http.MethodPost:
	//		p.processWin(w, player)
	//	case http.MethodGet:
	//		p.showScore(w, player)
	//	}
	//}))
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request){

	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}

}

func(i *InMemoryPlayerStore)GetLeague() []Player{
	return nil
}

func (p *PlayerServer)leagueHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}












































