package authapp

import "github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"

// Request
type CreatePatientRequest struct{}

// Use Case
type CreatePatientUseCase struct {
	LocationRepository shareddomain.LocationRepository
}

func (uc CreatePatientUseCase) Command(request CreatePatientRequest) error {
	return nil
}
