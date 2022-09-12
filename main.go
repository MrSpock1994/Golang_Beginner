package main

import "fmt"

func main() {
	var conference_name = "Go Conference"
	const conference_tickets = 50
	var remaining_tickets = 50

	//fmt.Println("Welcome to our", conference_name, "booking application!")
	fmt.Printf("Welcome to our %v booking application.\n", conference_name)
	//fmt.Println("We have a total of", conference_tickets, "tickets and", remaining_tickets, "are still avaiable.")
	fmt.Printf("We have a total of %v tickets and %v are still avaiable.\n", conference_tickets, remaining_tickets)
	fmt.Println("Get Your Tickets here!")

	var first_name string
	var user_tickets int

	// Get user input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&first_name)

	user_tickets = 2
	fmt.Printf("User %v has %v tickets\n", first_name, user_tickets)

}
