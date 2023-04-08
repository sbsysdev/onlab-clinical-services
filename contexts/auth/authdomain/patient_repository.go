package authdomain

type PatientRepository interface {
	CreatePatient(PatientEntity) error

	// ReadPatientById(string) (PatientEntity, error)
}
