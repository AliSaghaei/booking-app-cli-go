package main

import (
	"fmt"
	"sync"
	"time"
)

// package level variables
const conferenceTickets int= 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	

	for {
	
	// Getting User Input
	firstName,lastName,email, userTickets := getUserInput()

	// User input validation function
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName,lastName, email, userTickets, remainingTickets)

	// Condition for program logic (booking) execution
		if isValidName && isValidEmail && isValidTicketNumber {
		
			bookTicket(userTickets, firstName, lastName,email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName,email)
		
			// Call function GetFirstNames
			firstNames := getFirstNames()
			fmt.Printf("The fist names of bookings are: %v\n", firstNames)

			// ending the program when we run out of tickets
			if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
		}else {
			if !isValidName{
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail{
				fmt.Println("Email address you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber{
				fmt.Println("Number of tickets you entered is invalid.")
			}
		}
	}
	wg.Wait()
}

//function for greating users
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// function for extracting the first names from the users slice and storing them in a slice
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

// fucntion for getting user input (first name, last name ,email and number of tickets)
func getUserInput() (string,string,string,uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	
	fmt.Println("Enter Your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName,lastName,email,userTickets
}

// booking function
func bookTicket(userTickets uint, firstName string,lastName string, email string){
	remainingTickets = remainingTickets - userTickets


	var userData = userData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData )

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("########################")
	wg.Done()
}