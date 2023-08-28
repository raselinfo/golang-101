package test

import (
	"reflect"
	"testing"
)

// Test if the TestFunction creates a booking for a user with valid name, email, and number of tickets
func TestValidBooking(t *testing.T) {
	name := "John Doe"
	email := "johndoe@example.com"
	numberOfTickets := uint(2)

	booking := TestFunction(name, email, numberOfTickets)

	if len(booking) != 1 {
		t.Errorf("Expected booking length to be 1, got %d", len(booking))
	}

	expectedBooking := UserData{
		Name:            name,
		Email:           email,
		NumberOfTickets: numberOfTickets,
	}

	if !reflect.DeepEqual(booking[0], expectedBooking) {
		t.Errorf("Expected booking to be %v, got %v", expectedBooking, booking[0])
	}
}
