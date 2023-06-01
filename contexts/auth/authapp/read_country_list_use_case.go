package authapp

import "github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"

// Use Case
type ReadCountryListUseCase struct {
}

func (uc ReadCountryListUseCase) Query() ([]authdomain.Country, error) {
	return []authdomain.Country{}, nil
}
