package authdomain

import "github.com/google/uuid"

// Patient Entity
type PatientEntity struct {
	ID       string
	Person   Person
	NID      NationalIdentityDocumentNumber
	User     User
	Contacts Contacts
}

// Patient Entity Factory
func CreatePatient(person Person, nid NationalIdentityDocumentNumber, user User, contacts Contacts) PatientEntity {
	return PatientEntity{
		ID:       uuid.New().String(),
		Person:   person,
		NID:      nid,
		User:     user,
		Contacts: contacts,
	}
}
