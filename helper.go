package main

import "strings"

func validateUserInput(firstName string, email string, lastName string, userTickets uint) (bool, bool, bool) {
	// input validation
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickAmt := userTickets > 0 && userTickets <= uint(remainingTickets)
	return isValidName, isValidEmail, isValidTickAmt
}
