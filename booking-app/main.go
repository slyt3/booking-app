package main

// imported packages
import (
	"fmt"
	"sync"
	"time"
)

// global variables starts with FIRST letter of variable name that indicates that its global
//package variables

// const - can't change value
const conferenceTickets = 50

var conferenceName = "Go Conference"

// uint - can't be negative
var remainingTickets uint = 50

// creating slice with type of structure
var bookings = make([]UserData, 0)

// creating slice with type of maps.
// var bookings = make([]map[string]string, 0)

//creating custom data structure
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// creating empty thread
var wg = sync.WaitGroup{}

func main() {

	//printing information which welcomes user
	greetUsers()

	for {

		// gets user input
		firstName, lastName, email, userTickets := getUserInput()

		//validate user input
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// book tickets
			bookTickets(firstName, lastName, userTickets, email)

			//adding thread
			wg.Add(1)
			//sends ticket to user information
			go sendTicket(userTickets, firstName, lastName, email)

			//prints all first names of who booked
			var printsFirstNames []string = getFirstNames()

			fmt.Printf("The first names of bookings are: %v\n", printsFirstNames)

			var noTicketsRemaining bool = remainingTickets == 0

			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				// break
			}
		} else {

			// checks if its invalid input and prints what is wrong

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
		//if program is complete it will wait until all thread will be executed
		wg.Wait()
	}
}

func greetUsers() {

	// printing information
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {

	// creating empty slice
	firstNames := []string{}

	//loops through slice to get names
	for _, booking := range bookings {
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {

	//local variables
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

func bookTickets(firstName string, lastName string, userTickets uint, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	// var userData = make(map[string]string)

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// creating structures with own data type
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// adds created structure to slice
	bookings = append(bookings, userData)

	//prints information
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	// 10 second delay
	time.Sleep(10 * time.Second)

	// prints information
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n################")
	fmt.Printf("Sending ticket:\n%v\nto email adress %v\n", ticket, email)
	fmt.Println("################")

	//thread get closed
	wg.Done()
}
