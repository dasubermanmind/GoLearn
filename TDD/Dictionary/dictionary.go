package main

import "os/exec"

type Dictionary map[string]string

func (d Dictionary)Search(word string) (string, error){
	definition , ok := d[word]
	if !ok{
		return "" , exec.ErrNotFound
	}
	return definition, nil
}


func (d Dictionary)Add(word, definition string) error{
	d[word] = definition
	return nil
}

func (d Dictionary)Delete(word string){
	delete(d,word)
}

