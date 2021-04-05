package main

import "testing"

const repeatedCount = 5

func Repeat(character string ) string {
	var repeated string;

	for i := 0; i<repeatedCount; i++ {
		repeated += repeated+character
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B){
	for i:=0; i<b.N; i++{
		Repeat("a")
	}
}


