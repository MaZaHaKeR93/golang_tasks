package main

import (
	"fmt"
	"math/rand"
)

func showPriceList() {
	const secondsPerDay = 86400
	spacelines := [3]string{"Space Adventures", "SpaceX", "Virgin Galactic"}
	ticketsCount := 10
	distanceToMars := 62100000
	fmt.Printf("%-16v %4v %-10v %4v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("======================================")
	for i := 1; i <= ticketsCount; i++ {
		spaceline := spacelines[rand.Intn(3)]
		sheepSpeed := rand.Intn(15) + 16
		ticketPrice := sheepSpeed + 20
		tripType := "One-way"
		daysToMars := distanceToMars / sheepSpeed / secondsPerDay
		if rand.Intn(2) == 1 {
			tripType = "Round-way"
			ticketPrice *= 2
		}
		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceline, daysToMars, tripType, ticketPrice)
	}
	fmt.Println("======================================")
}
