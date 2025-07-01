package main

import "fmt"

import "rsc.io/quote" // module that

import "math/rand"

func main() {
	quotes := []string{quote.Go(), quote.Hello(), quote.Glass(), quote.Opt()}
	fmt.Println(quotes[rand.Intn(len(quotes))])
	// fmt.Println(quote.Go())
}
