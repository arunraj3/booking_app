package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {
	//Greet the User
	greetUsers()

	//Taking input from the user and processing the inputs
	for {
		//User Inputs
		fName, lName, email, userTickets := getUserInputs()
		//Validating the UserInputs
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(fName, lName, userTickets, email)

		//if all the validations are true book the tickets requested by the user
		if isValidEmail && isValidName && isValidTicketNumber {

			bookTickets(userTickets, fName, lName, email)
			// to create an empty array slices of strings basically the dynamic Array
			// to extract the firstnames of the user registered for the conference for privacy purpose
			var firstNames []string = getFirstNames()
			fmt.Printf("The FirstNames of the peoples who have registered : %v\n", firstNames)

			//All the tickets are booked
			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out,Please come next Year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or Last name you have entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ symbol")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is Invalid")
			}
			fmt.Printf("Please try again\n")
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome To %v booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, names := range bookings {
		firstNames = append(firstNames, (strings.Fields(names))[0])
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var fName string
	var lName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your firstName : ")
	fmt.Scan(&fName)
	fmt.Print("Enter your lastName : ")
	fmt.Scan(&lName)
	fmt.Printf("Enter your email : ")
	fmt.Scan(&email)
	fmt.Printf("Enter number of tickets : ")
	fmt.Scan(&userTickets)

	return fName, lName, email, userTickets
}

func bookTickets(userTickets uint, fName string, lName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, fName+" "+lName)
	fmt.Printf("%v", bookings)
	fmt.Printf("Thank you %s for booking %d tickets. You will receive a confirmation email at %s\n", fName+" "+lName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
