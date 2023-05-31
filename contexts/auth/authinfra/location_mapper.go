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

func FromDepartmentModelToDepartmentEntity(model dbshared.Department) authdomain.Department {
	return authdomain.Department{
		Id:   model.ID,
		Name: model.Name,
	}
}

func FromMunicipalityModelToMunicipalityEntity(model dbshared.Municipality) authdomain.Municipality {
	return authdomain.Municipality{
		Id:   model.ID,
		Name: model.Name,
	}
}
