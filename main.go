package main

import "fmt"

func main() {
	// var means variable (mutable), const means constant (unmutable), unit does not allow negative integers.
	var conferenceName = "X-Games Aspen"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// create an array to hold bookings. (Arrays are a fixed size) similar to a list in python.
	bookings := []string{}

	// Println = print with a new line, printf works like a f print in python.
	fmt.Printf("Book your tickets for %v now!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// when you dont assign a value you have to specify the data type being passed.
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		// ask user for input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// take user tickets - remaining tickets
		remainingTickets = remainingTickets - uint(userTickets)

		// add user to bookings.
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("These are all our bookings: %v\n", bookings)
	}
}
