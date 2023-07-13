package main 
import "fmt"

func main(){
	var conferenceName  = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is of Type : %T\n",conferenceName)
	fmt.Println("Welcome to our",conferenceName,"Booking Application")
	fmt.Printf("We have total of %d tickets and %d are still available\n",conferenceTickets,remainingTickets)

	var userName string
	var userTickets int 


	
	userName = "arun"
	userTickets = 2
	fmt.Printf("User %s booked %d tickets",userName,userTickets)
}