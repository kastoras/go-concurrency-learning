package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping
		if !ok {
			//do something
		}

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	//create channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press Enter (enter Q to quit)")

	for {
		// print promt
		fmt.Print("-> ")

		var userInput string
		_, _ = fmt.Scan(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput

		response := <-pong

		fmt.Println("Response: ", response)
	}

	fmt.Println("All done, closing channels")
	close(ping)
	close(pong)
}
