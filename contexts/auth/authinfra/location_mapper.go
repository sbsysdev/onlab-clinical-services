package authinfra

import (
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbshared"
)

func FromCountryModelToCountryEntity(model dbshared.Country) authdomain.Country {
	return authdomain.Country{
		Id:      model.ID,
		Name:    model.Name,
		Calling: model.Calling,
	}
}

func FromCountryModelToCountryEntityFilled(model dbshared.Country) authdomain.Country {
	departments := make([]authdomain.Department, len(model.Departments))

	for i, v := range model.Departments {
		departments[i] = FromDepartmentModelToDepartmentEntityFilled(v)
	}

	return authdomain.Country{
		Id:          model.ID,
		Name:        model.Name,
		Calling:     model.Calling,
		Departments: departments,
	}
}

func FromDepartmentModelToDepartmentEntity(model dbshared.Department) authdomain.Department {
	return authdomain.Department{
		Id:   model.ID,
		Name: model.Name,
	}
}

func FromDepartmentModelToDepartmentEntityFilled(model dbshared.Department) authdomain.Department {
	municipalities := make([]authdomain.Municipality, len(model.Municipalities))

	for i, v := range model.Municipalities {
		municipalities[i] = FromMunicipalityModelToMunicipalityEntity(v)
	}

	return authdomain.Department{
		Id:             model.ID,
		Name:           model.Name,
		Municipalities: municipalities,
	}
}

func FromMunicipalityModelToMunicipalityEntity(model dbshared.Municipality) authdomain.Municipality {
	return authdomain.Municipality{
		Id:   model.ID,
		Name: model.Name,
	}
}
