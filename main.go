package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
	// eb := englishBot{}
	// sb := spanishBot{}
	// printGreeting(eb)
	// printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hello there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
