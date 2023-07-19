package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	greetUsers(conferenceName, remainingTickets, conferenceTickets)

	bookings := []string{}
	var fName string
	var lName string
	var email string
	var userTickets uint

	for {
		fName, lName, email, userTickets = getUserInputs()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(fName, lName, userTickets, int(remainingTickets), email)

		if isValidEmail && isValidName && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, fName+" "+lName)

			fmt.Printf("%v", bookings)
			fmt.Printf("Thank you %s for booking %d tickets. You will receive a confirmation email at %s\n", fName+" "+lName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// to create an empty array slices of strings basically the dynamic Array
			var firstNames []string = getFirstNames(bookings)
			fmt.Printf("The FirstNames of the peoples who have registered : %v\n", firstNames)

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

func greetUsers(conferenceName string, remainingTickets uint, conferenceTickets int) {
	fmt.Printf("Welcome To %v booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, names := range bookings {
		firstNames = append(firstNames, (strings.Fields(names))[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, userTickets uint, remainingTickets uint, email string) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
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
