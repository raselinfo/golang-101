package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTicket uint = 50
	var remainingTickets uint = conferenceTicket

	fmt.Printf("Welcome to our %v booking application", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTicket, remainingTickets)

	fmt.Println("Get your tickets here to attend the conference")

	var firstName string
	var lastName string
	var userTicket int
	var email string

	fmt.Println("Enter Your Name: ")
	fmt.Scan(&firstName, &lastName)

	fmt.Println("Enter Number of Tickets:")
	fmt.Scan(&userTicket)

	fmt.Println("Enter Your Email: ")
	fmt.Scan(&email)

	fmt.Printf("Thank you %v %v for booking %v tickets for %v . The confirmation mail you will recive at %v", firstName, lastName, userTicket, conferenceName, email)

}
