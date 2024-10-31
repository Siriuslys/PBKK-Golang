package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	//for remainingTickets > 0 && len(bookings) < 50 {

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
	// isInValidCity := city == "singapore" || city == "london"

	if isValidEmail && isValidName && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("These are all our bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("Ticket is sold out.")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("name is to short")
		}

		if !isValidEmail {
			fmt.Println("email addres is not valid")
		}

		if !isValidTicketNumber {
			fmt.Println("number of tickets is invalid")
		}
	}

	// }

	// city := "London"

	// switch city{
	// 	case "New York":

	// 	case "Singapore":

	// 	case "London":

	// 	default:
	// 		fmt.Print("No valid city")
	// }

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v  tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
// 	return isValidName, isValidEmail, isValidTicketNumber
// }

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string

	var userTickets uint
	// ask user for their name
	fmt.Printf("Enter your first name:\n")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name:\n")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email:\n")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets:\n")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var UserData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, UserData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
