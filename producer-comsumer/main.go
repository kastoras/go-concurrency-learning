package main

import (
	"math/rand"
	"time"

	"github.com/labstack/gommon/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of whick pizza we are making

	// run forever or until we reveive a quit notification
	// try to make pizza

	for {
		//try to make a pizza
		// decision
	}

}

func main() {
	// seed the ranom number generator
	rand.Seed(time.Now().UnixNano())

	//print out a message
	color.Cyan("The Pizzaria is open fot business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer

	//printout the ending message
}
