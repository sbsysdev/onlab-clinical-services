package authinfra

import "github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"

type PatientEmailService struct{}

func (service PatientEmailService) SendWelcomeEmailToPatient(patient authdomain.PatientEntity) error {
	return nil
}
func (service PatientEmailService) SendRecoveryEmailToPatient(patient authdomain.PatientEntity, token string) error {
	return nil
}
