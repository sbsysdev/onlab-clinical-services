package authapp

import "github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"

// Use Case
type ReadCountryListUseCase struct {
	LocationRepository authdomain.LocationRepository
}

func (uc ReadCountryListUseCase) Query() ([]authdomain.Country, error) {
	return uc.LocationRepository.GetCountryList()
}
