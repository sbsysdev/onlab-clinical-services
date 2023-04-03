package authdomain

import (
	"errors"
	"time"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
)

// Person Name Value Object
type PersonName string

const (
	ERRORS_PERSON_NAME_EMPTY shareddomain.DomainError = "ERRORS_PERSON_NAME_EMPTY"
)

func CreatePersonName(name string) (PersonName, error) {
	if len(name) == 0 {
		return PersonName(""), errors.New(string(ERRORS_PERSON_NAME_EMPTY))
	}

	return PersonName(name), nil
}

// Person Surname Value Object
type PersonSurname string

const (
	ERRORS_PERSON_SURNAME_EMPTY shareddomain.DomainError = "ERRORS_PERSON_SURNAME_EMPTY"
)

func CreatePersonSurname(surname string) (PersonSurname, error) {
	if len(surname) == 0 {
		return PersonSurname(""), errors.New(string(ERRORS_PERSON_SURNAME_EMPTY))
	}

	return PersonSurname(surname), nil
}

// Person Birth Value Object
type PersonBirth time.Time

const (
	ERRORS_PERSON_BIRTH_NOT_VALID shareddomain.DomainError = "ERRORS_PERSON_BIRTH_NOT_VALID"
	ERRORS_PERSON_BIRTH_UNDER_AGE shareddomain.DomainError = "ERRORS_PERSON_BIRTH_UNDER_AGE"
)

func CreatePersonBirth(birth time.Time) (PersonBirth, error) {
	if birth.After(time.Now()) {
		return PersonBirth(birth), errors.New(string(ERRORS_PERSON_BIRTH_NOT_VALID))
	}

	return PersonBirth(birth), nil
}

func CreatePersonBirthAdult(birth time.Time, minAdult uint8) (PersonBirth, error) {
	if birth.AddDate(int(minAdult), 0, 0).After(time.Now()) {
		return PersonBirth(birth), errors.New(string(ERRORS_PERSON_BIRTH_UNDER_AGE))
	}

	return PersonBirth(birth), nil
}

// Person Sex Value Object
type PersonSex string

const (
	ERRORS_PERSON_SEX_NOT_VALID shareddomain.DomainError = "ERRORS_PERSON_SEX_NOT_VALID"
)

const (
	SEX_MALE   PersonSex = "male"
	SEX_FEMALE PersonSex = "female"
)

func CreatePersonSex(sex string) (PersonSex, error) {
	if sex != string(SEX_MALE) && sex != string(SEX_FEMALE) {
		return PersonSex(""), errors.New(string(ERRORS_PERSON_SEX_NOT_VALID))
	}

	return PersonSex(sex), nil
}

// Person Value Object
type Person struct {
	Name    PersonName
	Surname PersonSurname
	Birth   PersonBirth
	Sex     PersonSex
}

// Person Entity Factory
func CreatePerson(name PersonName, surname PersonSurname, birth time.Time, sex PersonSex) (Person, error) {
	return Person{}, nil
}
