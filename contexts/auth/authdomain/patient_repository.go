package authdomain

type PatientRepository interface {
	CreatePatient(PatientEntity) error

	ReadPatientByName(string) (PatientEntity, error)

	ReadPatientById(string) (PatientEntity, error)

	ReadPatientByEmail(string) (PatientEntity, error)
}
