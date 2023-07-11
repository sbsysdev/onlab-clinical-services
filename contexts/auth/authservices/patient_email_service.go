package authservices

import "github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"

type PatientEmailService interface {
	SendWelcomeEmailToPatient(authdomain.PatientEntity, string) error

	SendRecoveryEmailToPatient(authdomain.PatientEntity, string) error
}
