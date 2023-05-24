package authdomain

import "github.com/google/uuid"

// Patient Entity
type PatientEntity struct {
	ID       string                         `json:"id"`
	Person   Person                         `json:"person"`
	NID      NationalIdentityDocumentNumber `json:"nid"`
	User     User                           `json:"user"`
	Contacts Contacts                       `json:"contacts"`
	Roles    []AliasRoleEntity              `json:"roles"`
}

// Patient Entity Factory
func CreatePatient(person Person, nid NationalIdentityDocumentNumber, user User, contacts Contacts, roles []AliasRoleEntity) PatientEntity {
	return PatientEntity{
		ID:       uuid.New().String(),
		Person:   person,
		NID:      nid,
		User:     user,
		Contacts: contacts,
		Roles:    roles,
	}
}
