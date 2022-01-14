package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceName string = "Wach Conference"
const conferenceTickets uint = 100

var remainingTickets uint = 100
var bookings = make([]UserData, 0)

type UserData struct {
	firstName     string
	lastName      string
	email         string
	ticketsBooked uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	// for {
	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(firstName, lastName, email, userTickets)

		wg.Add(1)
		go sendTicket(firstName, lastName, email, userTickets)

		// display user first name
		firstNames := getFirstNames()
		fmt.Printf("These are the first name of the booked tickets %v\n", firstNames)

	} else if remainingTickets == 0 {
		fmt.Printf("Sorry, %v tickets have been booked out for this year\n", conferenceName)
		// break
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

	wg.Wait()
	// }

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("------------------------------------------------------")
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

// # Get user first name logic
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		ticketsBooked: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation sent to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are left for %v\n", remainingTickets, conferenceName)
	fmt.Println("------------------------------------------------------")
}

func sendTicket(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("Dear %v, you booked %v tickets for %v.", firstName, userTickets, conferenceName)
	fmt.Println("##########################################")
	fmt.Printf("Sending ticket:\n %v The tickets will be sent to %v\n", ticket, email)
	fmt.Println("##########################################")
	wg.Done()
}
