package main

import "fmt"

func main() {
	// var means variable (mutable), const means constant (unmutable)
	var conferenceName = "Go Conference"
	const conferenceTickets = 50

	// Println = print with a new line
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("Get your tickets here to attend")

}
