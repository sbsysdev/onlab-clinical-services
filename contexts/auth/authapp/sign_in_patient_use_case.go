package authapp

import (
	"errors"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
)

// Request
type SignInPatientRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Response
type SignInPatientResponse struct {
	Token   string                   `json:"token"`
	Patient authdomain.PatientEntity `json:"patient"`
}

// Use Case
type SignInPatientUseCase struct {
	// Repositories
	PatientRepository authdomain.PatientRepository
}

func (uc SignInPatientUseCase) Query(request SignInPatientRequest) (SignInPatientResponse, error) {
	patient, patientErr := uc.PatientRepository.ReadPatientByName(request.Name)

	if patientErr != nil {
		return SignInPatientResponse{}, errors.New(string(authdomain.ERRORS_USER_NAME_NOT_FOUND))
	}

	if err := authdomain.ComparePasswordAndHash(request.Password, string(patient.User.Password)); err != nil {
		return SignInPatientResponse{}, err
	}

	patient.User.Password = ""

	return SignInPatientResponse{
		Token:   "",
		Patient: patient,
	}, nil
}
