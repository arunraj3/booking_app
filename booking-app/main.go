package main

import (
	"fmt"
	"strings"
)

func main(){
	conferenceName  := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets  uint = 50

	fmt.Printf("conferenceName is of Type : %T\n",conferenceName)
	fmt.Println("Welcome to our",conferenceName,"Booking Application")
	fmt.Printf("We have total of %d tickets and %d are still available\n",conferenceTickets,remainingTickets)

	bookings := []string{}

	

	var fName string
	var lName string
	var email string
	var userTickets uint


	for{
		fmt.Printf("Enter your firstName : ")
		fmt.Scan(&fName)
		fmt.Print("Enter your lastName : ")
		fmt.Scan(&lName)
		fmt.Printf("Enter your email : ")
		fmt.Scan(&email)
		fmt.Printf("Enter number of tickets : ")
		fmt.Scan(&userTickets)
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings,fName+" "+lName)

		fmt.Printf("%v",bookings)
		fmt.Printf("Thank you %s for booking %d tickets. You will receive a confirmation email at %s\n",fName+" "+lName,userTickets,email)
		fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)

		// to create an empty array slices of strings basically the dynamic Array
		firstNames := []string{}

		for _ , name := range bookings{
			firstNames = append(firstNames,(strings.Fields(name))[0])
		}
		fmt.Printf("%v",firstNames)
		
	}


}