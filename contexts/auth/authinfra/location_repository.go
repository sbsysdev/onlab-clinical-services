package authinfra

import (
	"errors"

	"gorm.io/gorm"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbshared"
)

type LocationRepository struct {
	DB *gorm.DB
}

func (repo LocationRepository) GetMunicipalityById(municipalityId string) (authdomain.Municipality, error) {
	var founded dbshared.Municipality

	if err := repo.DB.Table("municipalities").First(&founded, "municipality_id = ?", municipalityId).Error; err != nil {
		return authdomain.Municipality{}, errors.New(string(authdomain.ERRORS_MUNICIPALITY_NOT_FOUND))
	}

	return FromMunicipalityModelToMunicipalityEntity(founded), nil
}

func (repo LocationRepository) GetDepartmentById(departmentId string) (authdomain.Department, error) {
	var founded dbshared.Department

	if err := repo.DB.Table("departments").First(&founded, "department_id = ?", departmentId).Error; err != nil {
		return authdomain.Department{}, errors.New(string(authdomain.ERRORS_DEPARTMENT_NOT_FOUND))
	}

	return FromDepartmentModelToDepartmentEntity(founded), nil
}

func (repo LocationRepository) GetCountryById(countryId string) (authdomain.Country, error) {
	var founded dbshared.Country

	if err := repo.DB.Table("countries").First(&founded, "country_id = ?", countryId).Error; err != nil {
		return authdomain.Country{}, errors.New(string(authdomain.ERRORS_COUNTRY_NOT_FOUND))
	}

	return FromCountryModelToCountryEntity(founded), nil
}

func (repo LocationRepository) GetCountryList() ([]authdomain.Country, error) {
	var founded []dbshared.Country

	if err := repo.DB.Table("countries").Preload("Departments").Preload("Departments.Municipalities").Find(&founded).Error; err != nil {
		return []authdomain.Country{}, errors.New(string(authdomain.ERRORS_COUNTRY_NOT_FOUND))
	}

	countries := make([]authdomain.Country, len(founded))

	for i, v := range founded {
		countries[i] = FromCountryModelToCountryEntityFilled(v)
	}

	return countries, nil
}
