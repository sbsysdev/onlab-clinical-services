package authdomain

// Patient Entity
type PatientEntity struct {
	Name     string
	Password string
	Person   interface{}
	Contacts interface{}
	State    string
}

// Patient Entity Factory
func CreatePatient() (PatientEntity, error) {
	return PatientEntity{}, nil
}
