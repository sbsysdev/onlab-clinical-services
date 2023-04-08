package authdomain

import "github.com/google/uuid"

// Patient Entity
type PatientEntity struct {
	ID       string
	Person   Person
	User     User
	Contacts Contacts
}

// Patient Entity Factory
func CreatePatient(person Person, user User, contacts Contacts) PatientEntity {
	return PatientEntity{
		ID:       uuid.New().String(),
		Person:   person,
		User:     user,
		Contacts: contacts,
	}
}
