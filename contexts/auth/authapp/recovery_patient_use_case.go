package authapp

import (
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authservices"
)

// Request
type RecoveryPatientRequest struct {
	Email string `json:"email"`
}

// Use Case
type RecoveryPatientUseCase struct {
	// Repositories
	PatientRepository authdomain.PatientRepository
	// Services
	EmailService authservices.PatientEmailService
}

func (uc RecoveryPatientUseCase) Command(request RecoveryPatientRequest) error {
	// Verify Patient by Email
	patient, patientErr := uc.PatientRepository.ReadPatientByEmail(request.Email)

	if patientErr != nil {
		return patientErr
	}

	// Generate recovery token
	signed, signedErr := authdomain.CreatePatientRecoveryToken(patient.ID, string(patient.Contacts.Email))

	if signedErr != nil {
		return signedErr
	}

	// Send email to patient with recovery token
	return uc.EmailService.SendRecoveryEmailToPatient(patient, signed)
}
