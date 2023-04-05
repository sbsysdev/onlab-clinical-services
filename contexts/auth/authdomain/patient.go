package authdomain

// Patient Entity
type PatientEntity struct {
	Person   Person
	User     User
	Contacts Contacts
	State    string
}

// Patient Entity Factory
func CreatePatient(person Person, user User, contacts Contacts) PatientEntity {
	return PatientEntity{
		Person:   person,
		User:     user,
		Contacts: contacts,
	}
}
