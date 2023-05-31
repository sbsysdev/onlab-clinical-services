package authdomain

import "github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"

// Municipality Value Object
type Municipality struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Department Value Object
type Department struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Municipalities []Municipality `json:"municipality"`
}

// Country Value Object
type Country struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Calling     string       `json:"calling"`
	Departments []Department `json:"department"`
}

const (
	ERRORS_COUNTRY_NOT_FOUND      shareddomain.DomainError = "ERRORS_COUNTRY_NOT_FOUND"
	ERRORS_DEPARTMENT_NOT_FOUND   shareddomain.DomainError = "ERRORS_DEPARTMENT_NOT_FOUND"
	ERRORS_MUNICIPALITY_NOT_FOUND shareddomain.DomainError = "ERRORS_MUNICIPALITY_NOT_FOUND"
)
