package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type Seeker interface{
	Seek(offset int64, whence int) (int64, error)
}

type Reader interface{
	Reader
	Seeker
}

func NewLeague(rdr io.ReadCloser)([]Player,error){
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil{
		err = fmt.Errorf("Got this error, %v", err)
	}

	return league, err
}

