package authdomain

type PatientRepository interface {
	CreatePatient(PatientEntity) error

	ReadPatientByName(string) (PatientEntity, error)
}
