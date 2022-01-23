package main

import (
	"strings"
)

// same package share same package variables

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {

	// checks is firstname and lastname longer then 2 char
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2

	// checks if it contains in string given char
	var isValidEmail bool = strings.Contains(email, "@")

	//check is ticket more then left tickets 0 or equal or less to remaining tickets
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
