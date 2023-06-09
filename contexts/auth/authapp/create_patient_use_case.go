package authapp

import (
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
)

// Request
type CreatePatientRequest struct {
	Person   PersonRequest         `json:"person"`
	NID      string                `json:"nid"`
	User     UserRequest           `json:"user"`
	Contacts SingleContactsRequest `json:"contacts"`
}

// Use Case
type CreatePatientUseCase struct {
	// Repositories
	PatientRepository  authdomain.PatientRepository
	RoleRepository     authdomain.RoleRepository
	LocationRepository authdomain.LocationRepository
	// Publish Event
	PublishEvent shareddomain.PublishDomainEvent
}

func (uc CreatePatientUseCase) Command(request CreatePatientRequest) error {
	// Person value object

	personName, nameErr := authdomain.CreatePersonName(request.Person.Name)

	if nameErr != nil {
		return nameErr
	}

	personSurname, surnameErr := authdomain.CreatePersonSurname(request.Person.Surname)

	if surnameErr != nil {
		return surnameErr
	}

	personBirth, birthErr := authdomain.CreatePersonBirthAdult(request.Person.Birth, 18 /* Implement min age config by organization or country */)

	if birthErr != nil {
		return birthErr
	}

	personSex, sexErr := authdomain.CreatePersonSex(request.Person.Sex)

	if sexErr != nil {
		return sexErr
	}

	person := authdomain.CreatePerson(personName, personSurname, personBirth, personSex)

	// NID value object

	nid, nidErr := authdomain.CreateNIDNumber(request.NID)

	if nidErr != nil {
		return nidErr
	}

	// User value object

	userName, usernameErr := authdomain.CreateUserName(request.User.Name)

	if usernameErr != nil {
		return usernameErr
	}

	userPassword, passwordErr := authdomain.CreateUserPassword(request.User.Password)

	if passwordErr != nil {
		return passwordErr
	}

	user := authdomain.CreateUser(userName, userPassword)

	// Contacts value object

	email, emailErr := authdomain.CreateEmail(request.Contacts.Email)

	if emailErr != nil {
		return emailErr
	}

	country, countryErr := uc.LocationRepository.GetCountryById(request.Contacts.Phone.Country)

	if countryErr != nil {
		return countryErr
	}

	phone, phoneErr := authdomain.CreatePhone(country, request.Contacts.Phone.Phone)

	if phoneErr != nil {
		return phoneErr
	}

	municipality, municipalityErr := uc.LocationRepository.GetMunicipalityById(request.Contacts.Address.Municipality)

	if municipalityErr != nil {
		return municipalityErr
	}

	address, addressErr := authdomain.CreateAddress(municipality, request.Contacts.Address.Address, request.Contacts.Address.Latitude, request.Contacts.Address.Longitude)

	if addressErr != nil {
		return addressErr
	}

	contacts := authdomain.CreateSingleContacts(email, phone, address)

	// Role entity
	patientRoles, patientRolesErr := uc.RoleRepository.GetAliasRolesByAlias([]authdomain.RoleAlias{authdomain.ALIAS_PATIENT})

	if patientRolesErr != nil {
		return patientRolesErr
	}

	// Patient entity

	patient := authdomain.CreatePatient(person, nid, user, contacts, patientRoles)

	//Store patient

	if err := uc.PatientRepository.CreatePatient(patient); err != nil {
		return err
	}

	// Dispath events

	if err := uc.PublishEvent(authdomain.CreatePatientCreatedEvent(patient)); err != nil {
		return err
	}

	return nil
}
