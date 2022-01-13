package main

import (
	"fmt"
	"strings"
)

const conferenceName string = "Wach Conference"
const conferenceTickets uint = 100

var remainingTickets uint = 100
var bookings []string

func main() {

	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookedUser(firstName, lastName, email, userTickets)

			firstNames := getFirstNames()
			fmt.Printf("These are the first name of the booked tickets %v\n", firstNames)

		} else if remainingTickets == 0 {
			fmt.Printf("Sorry, %v tickets have been booked out for this year\n", conferenceName)
			break
		} else {
			if !isValidName {
				fmt.Println("first or last name entered is invalid")
			}
			if !isValidEmail {
				fmt.Println("email entered is missing @")
			}
			if !isValidTicketNumber {
				fmt.Println("ticket number entered is invalid")
			}
			fmt.Println("Try Again...")
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Print("Enter your First Name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your Email address: ")
	fmt.Scan(&email)
	fmt.Print("How many tickets would you like to purchase? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func bookedUser(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation sent to %v\n", firstName, lastName, userTickets, email)
	fmt.Println("We'll be glad to see you there!")
	fmt.Printf("%v tickets are left for %v\n", remainingTickets, conferenceName)

}
