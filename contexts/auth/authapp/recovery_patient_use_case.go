package authapp

import (
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
)

type RecoveryPatientRequest struct {
	Email string `json:"email"`
}

type RecoveryPatientUseCase struct {
	// Repositories
	PatientRepository authdomain.PatientRepository
}

func (uc RecoveryPatientUseCase) Command(request RecoveryPatientRequest) error {
	return nil
}
