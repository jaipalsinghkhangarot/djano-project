package main

import (
	"fmt"
	"sync"
	"time"
)

// package level variables
const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	fName        string
	lName        string
	emailadd     string
	numOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// var bookings = []string{} or bookings := []string{} .//this is slice. [50]string is Array
	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T.\n", conferenceTickets, remainingTickets, conferenceName)

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInputs()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			userFirstNames := getFirstNames()
			fmt.Printf("The First names of bookings are %v\n", userFirstNames)
			//var noTicketsRemaining bool = remainingTickets == 0
			if remainingTickets == 0 {
				//end the application
				fmt.Println("Our Conference is booked out. Come back next Year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Entered First name or last name is too short, Try Again!!")
			}
			if !isValidEmail {
				fmt.Println("Entered email doesn't contains @ sign, Try Again!!")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets entered is invalid, Try Again!!")
			}

		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Println("Welcome to", conferenceName, "booking application.")
	fmt.Printf("We have Total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask information from user
	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName) //pointer to memory address of firstName

	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email Address:")
	fmt.Scan(&email)

	fmt.Println("Enter Number of Tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getFirstNames() []string {
	userFirstNames := []string{}
	for _, booking := range bookings {
		userFirstNames = append(userFirstNames, booking.fName)
	}
	return userFirstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//var userData = make(map[string]string) //create map userData
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	//create struct userData
	var userData = UserData{
		fName:        firstName,
		lName:        lastName,
		emailadd:     email,
		numOfTickets: userTickets,
	}

	//bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)
	fmt.Printf("Booking data: %v\n", bookings)

	fmt.Printf("Thanks you %v %v for booking %v tickets. You will receive confirmation mail at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are now remaining for Conference.\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending Ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#####################")
	wg.Done()
}

// //switch case example
// city := "London"

// switch city {
// 	case "London":
// 		//executes code for London
// 	case "Delhi":
// 		//executes code for delhi
// 	case "Berlin","Paris":
// 		//executes code for berlin or paris
// 	case "New York":
// 		//executes code for new york
// 	default:
// 		fmt.Println("Wrong city name")
// }
