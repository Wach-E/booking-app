package main

import "fmt"

func main() {
	const conferenceName string = "Wach Conference"
	const conferenceTickets uint = 100
	var remainingTickets uint = 100
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var bookings []string

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Print("Enter your First Name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your Email address: ")
	fmt.Scan(&email)
	fmt.Print("How many tickets would you like to purchase? ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation sent to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are left for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are the bookings\n%v", bookings)
}
