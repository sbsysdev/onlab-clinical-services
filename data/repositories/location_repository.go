package repositories

import (
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
	"gorm.io/gorm"
)

type LocationRepository struct {
	DB *gorm.DB
}

func (repository LocationRepository) GetMunicipalityById(municipalityId uint16) (shareddomain.Municipality, error) {
	return shareddomain.Municipality{}, nil
}

func (repository LocationRepository) GetDepartmentById(departmentId uint16, fillMunicipalities bool) (shareddomain.Department, error) {
	return shareddomain.Department{}, nil
}

func (repository LocationRepository) GetCountryById(countryId uint8, fillDepartments bool, fillMunicipalities bool) (shareddomain.Country, error) {
	return shareddomain.Country{}, nil
}

func (repository LocationRepository) GetCountryList(fillDepartments bool, fillMunicipalities bool) ([]shareddomain.Country, error) {
	return []shareddomain.Country{}, nil
}
