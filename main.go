package main

import (
	"fmt"
	"strings"
)

func main() {
	var conference_name = "Go Conference"
	const conference_tickets = 50
	var remaining_tickets uint = 50
	bookings := []string{}

	//fmt.Println("Welcome to our", conference_name, "booking application!")
	fmt.Printf("Welcome to our %v booking application.\n", conference_name)
	//fmt.Println("We have a total of", conference_tickets, "tickets and", remaining_tickets, "are still avaiable.")
	fmt.Printf("We have a total of %v tickets and %v are still avaiable.\n", conference_tickets, remaining_tickets)
	fmt.Println("Get Your Tickets here!")

	for {
		var first_name string
		var last_name string
		var email string
		var user_tickets uint

		// Get user input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&first_name)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&last_name)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&user_tickets)

		remaining_tickets -= user_tickets
		bookings = append(bookings, first_name+" "+last_name)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", first_name, last_name, user_tickets, email)
		fmt.Printf("%v tickets remaining.\n", remaining_tickets)

		firstnames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstnames = append(firstnames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstnames)

	}

}
