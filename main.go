package main

import (
	"fmt"
	"strings"
)

// var means variable (mutable), const means constant (unmutable), unit does not allow negative integers.
const conferenceTickets = 50

var conferenceName = "GO conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {
	// call user greeting func
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickAmt := validateUserInput(firstName, lastName, email, userTickets)
		//if user tickets is greater than the remaining tickets continue.
		if isValidName && isValidEmail && isValidTickAmt {
			bookTicket(userTickets, firstName, lastName, email)

			firstName := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstName)

			// check if remaining tickets is 0
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email you entered needs to include the @ symbol")
			}
			if !isValidTickAmt {
				fmt.Println("your ticket amount is invalid")
			}
		}
	}
}
func greetUser() {
	// Println = print with a new line, printf works like a f print in python.
	fmt.Printf("Book your tickets for %v now!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	// loop through bookings list and slice and add first names to a list.
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	// when you dont assign a value you have to specify the data type being passed.
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)

	// add user to bookings.
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
