package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int= 50
	var remainingTickets uint = 50
	bookings := []string{}
	// var bookings = [50]string{"Nana","Nicole","Ali"}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	

	for {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name, email and number of tickets
	fmt.Println("Enter Your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	// User input validation function
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName,lastName, email, userTickets, remainingTickets)

	// Condition for program logic (booking) execution
	if isValidName && isValidEmail && isValidTicketNumber {
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName )

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

	
	// Call function GetFirstNames
	firstNames := getFirstNames(bookings)
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
}


func greetUsers(confName string, confTickets int, remaining uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, remaining)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	// extracting the first names from the users slice and storing them in a slice
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func validateUserInput(firstName string,lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool){
	isValidName :=  len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName ,isValidEmail, isValidTicketNumber
}