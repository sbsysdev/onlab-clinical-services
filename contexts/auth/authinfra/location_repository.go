package authinfra

import (
	"gorm.io/gorm"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
)

type LocationRepository struct {
	DB *gorm.DB
}

func (repo LocationRepository) GetMunicipalityById(municipalityId uint16) (authdomain.Municipality, error) {
	return authdomain.Municipality{}, nil
}

func (repo LocationRepository) GetDepartmentById(departmentId uint16, fillMunicipalities bool) (authdomain.Department, error) {
	return authdomain.Department{}, nil
}

func (repo LocationRepository) GetCountryById(countryId uint8, fillDepartments bool, fillMunicipalities bool) (authdomain.Country, error) {
	return authdomain.Country{}, nil
}

func (repo LocationRepository) GetCountryList(fillDepartments bool, fillMunicipalities bool) ([]authdomain.Country, error) {
	return []authdomain.Country{}, nil
}
