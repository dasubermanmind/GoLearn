package main

import "fmt"

const engPrefix = "Hello"
const spanishPrefix= "Hola"
const frenchPrefix = "Bonjour"

const spanish= "Spanish"
const french = "French"

func Hello(name string, language string) string{
	if name == ""{
		name = "World"
	}

	prefix := engPrefix

	switch language {

	case french:
		prefix = frenchPrefix
		break
	case spanish:
		prefix = spanishPrefix

	default:
		prefix = engPrefix
	}
	return prefix + name
}

func main(){
	fmt.Println(Hello("Eliot", "french"))
}
