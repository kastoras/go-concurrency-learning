package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var (
	seatingCapacity = 10
	arrivalRate     = 100
	cutDuration     = 1000 * time.Millisecond
	timeOpen        = 10 * time.Second
)

func main() {

	rand.Seed(time.Now().UnixNano())

	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarberDoneChan:  doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")
	shop.AddBarber("Frank")
	shop.AddBarber("Petros")
	shop.AddBarber("Nick")
	shop.AddBarber("Katia")
	shop.AddBarber("Maria")

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.CloseShopForTheDay()
		closed <- true
	}()

	i := 1

	go func() {
		for {
			randomMiliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMiliseconds)):
				shop.AddClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	<-closed
}
