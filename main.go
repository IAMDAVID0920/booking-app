package main

import "fmt"

func main() {
	// syntac sugar := declaration + assignment
	conferenceName := "Go Conference"
	const conferenceTicket int = 50
	// uint => unsigned int
	var remainingTickets uint = 50
	// var bookings = [50]string{"Nana", "Nicole", "Peter"}
	var bookings []string

	// %T a placeholder for getting type of variables
	// %v a placeholder for variable value
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTicket, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("We have total", conferenceTicket, "tickets and", remainingTickets, "tickets left")

	// You need to tell go compiler
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// Get user input
	// Will print the memory address of this userName variable
	// fmt.Println(&userName)
	fmt.Print("What is your first name: ")
	// ptr in golang is special variable
	// print out the memory location
	fmt.Scan(&firstName)

	fmt.Print("What is your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets you want to buy: ")
	fmt.Scan(&userTickets)

	// convert userTickets also unsigned int type
	remainingTickets -= uint(userTickets)
	// bookings[0] = firstName + " " + lastName
	// dynamic array list -> Called Slice
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings: %v\n", bookings)
}
