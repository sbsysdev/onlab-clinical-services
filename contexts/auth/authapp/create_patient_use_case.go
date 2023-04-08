package authapp

import (
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
)

// Request
type CreatePatientRequest struct {
	// TODO: Define create patient request
}

// Use Case
type CreatePatientUseCase struct {
	PatientRepository  authdomain.PatientRepository
	LocationRepository shareddomain.LocationRepository
}

func (uc CreatePatientUseCase) Command(request CreatePatientRequest) error {
	// TODO: Validate person value object
	// TODO: Validate User value object
	// TODO: Validate contacts value object
	// TODO: Validate patient entity

	// TODO: Store patient in repository

	// TODO: Dispath PATIENT_CREATED_EVENT
	return nil
}
