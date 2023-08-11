package main

import "strings"

// function for user input validation
func validateUserInput(firstName string,lastName string, email string, userTickets uint) (bool, bool, bool){
	isValidName :=  len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName ,isValidEmail, isValidTicketNumber
}