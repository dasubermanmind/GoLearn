package main

import "time"

type WebsiteChecker func(string)bool

type result struct{
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string )map[string]bool{

	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls{
		go func(u string) {
			//this sends the statement vai channel
			resultsChannel <- result{u, wc(u)}
		}(url)
	}

	// receives the expression
	for i:= 0; i<len(urls); i++{
		r := <-resultsChannel
		results[r.string] = r.bool
	}

	time.Sleep(2*time.Second)

	return results
}