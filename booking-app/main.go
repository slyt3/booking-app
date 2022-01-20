package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// book tickets

			remainingTickets, bookings := bookTickets(remainingTickets, bookings, firstName, lastName, userTickets, email, conferenceName)

			var printsFirstNames []string = FirstNames(bookings)

			fmt.Printf("The first names of bookings are: %v\n", printsFirstNames)

			var noTicketsRemaining bool = remainingTickets == 0

			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("your email address you entered doesnt contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}

	}
}

func greetUsers(conferenceName string, conferenceTickets uint, remainingTickets uint) {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func FirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// asks user for first name

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	// asks user for they last name

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	// asks user for they email

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	//asks user how much tickets do they want

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets uint, bookings []string, firstName string, lastName string, userTickets uint, email string, conferenceName string) (uint, []string) {
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return remainingTickets, bookings

}
