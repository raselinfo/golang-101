package test

import "fmt"

// UserData represents the details of a user.
type UserData struct {
	Name            string
	Email           string
	NumberOfTickets uint
}

// NewUserData creates a new UserData instance.
func NewUserData(name, email string, numberOfTickets uint) *UserData {
	return &UserData{
		Name:            name,
		Email:           email,
		NumberOfTickets: numberOfTickets,
	}
}

// TestFunction creates a booking for a user.
// It returns the booking details of the user.
func TestFunction(name, email string, numberOfTickets uint) []UserData {
	booking := make([]UserData, 0)

	booking = append(booking, *NewUserData(name, email, numberOfTickets))

	return booking
}

// String prints the UserData in a human-readable format.
func (u *UserData) String() string {
	return fmt.Sprintf("Name: %s, Email: %s, Number of Tickets: %d", u.Name, u.Email, u.NumberOfTickets)
}
