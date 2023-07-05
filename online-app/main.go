package main

import "fmt"

func main() {
	fmt.Println("Welcome to the application!")
	var userName string
	userName = "harsh"
	bookedTicket := 2
	totalTickets := 50
	userCount := 12
	remainingTickets := 50 - userCount
	fmt.Printf("there are %v tickets for the conference and now remaining is %v\n", totalTickets, remainingTickets)
	if bookedTicket > 1 {
		fmt.Printf("User %v booked  %v tickets for conference", userName, bookedTicket)

	}

}
