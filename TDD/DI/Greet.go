package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"net/http"
)

type Writer interface{
	Write(p []byte)( n int, err error)
}

func Fprintf(w io.Writer, format string, a ...interface{}) ( n int, err error){
	p := newPrinter()
	p.doPrintF(format,a)
	n, err = w.Write(p.buf)
	p.free()
	return
}

func Greet(writer io.Writer, name string) {
	fmt.Printf("String passed in as %s", name)
}


func PrintF(format string , a ...interface{}) (n int, err error){
	return Fprintf(os.Stdout, format, a...)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request){
	Greet(w,"hi")
}


func main(){
	
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
	
}
