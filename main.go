package main

import (
	"booking-app/helper"
	"fmt"
)

func main() {
	var conferenceName = "Go conference"
	const conferenceTicket uint = 50
	var remainingTickets uint = conferenceTicket
	// var booking []string

	fmt.Printf("Welcome to our %v booking application", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTicket, remainingTickets)

	fmt.Println("Get your tickets here to attend the conference")

	name := helper.PrintName("Rakibul")

	fmt.Println(name)

	// var city string

	// fmt.Scan(&city)

	// switch city {
	// case "New York":
	// 	fmt.Println("New York")
	// case "London":
	// 	fmt.Println("London")
	// default:
	// 	fmt.Println("No city")
	// }

	// for {
	// 	if remainingTickets == 0 {
	// 		fmt.Println("Sorry, we are sold out")
	// 		break
	// 	}

	// 	var firstName string
	// 	var lastName string
	// 	var userTicket int
	// 	var email string

	// 	fmt.Println("Enter Your Name: ")
	// 	fmt.Scan(&firstName, &lastName)

	// 	fmt.Println("Enter Number of Tickets:")
	// 	fmt.Scan(&userTicket)

	// 	fmt.Println("Enter Your Email: ")
	// 	fmt.Scan(&email)

	// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// 	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")

	// 	if userTicket < 0 {
	// 		println("Please enter valid number of tickets")
	// 		break
	// 	} else if uint(userTicket) > remainingTickets {
	// 		println("Sorry, we don't have enough tickets")
	// 	} else {
	// 		remainingTickets = remainingTickets - uint(userTicket)

	// 		booking = append(booking, firstName+" "+lastName)

	// 		fmt.Printf("\n \nThank you %v %v for booking %v tickets for %v . The confirmation mail you will recive at %v\n", firstName, lastName, userTicket, conferenceName, email)

	// 		fmt.Printf("\n Wee have %v tickets remaining\n", remainingTickets)

	// 		firstNames := []string{}
	// 		for _, booking := range booking {
	// 			var names = strings.Fields(booking)
	// 			var firstName = names[0]
	// 			firstNames = append(firstNames, firstName)
	// 		}

	// 		fmt.Printf("\n List of attendees: %v\n", firstNames)
	// 	}
	// }

}
