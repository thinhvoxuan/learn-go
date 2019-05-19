package main

import (
	"fmt"
	"io"
	"net/http"
)

const testVersion = 2
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// Helloworld function
func Helloworld(name string, lang string) string {
	if len(name) == 0 {
		name = "World"
	}
	return greetingPrefix(lang) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		return frenchHelloPrefix
	case spanish:
		return spanishHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
