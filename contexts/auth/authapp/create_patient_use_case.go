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
	PatientRepository  authdomain.PatientRepository
	LocationRepository shareddomain.LocationRepository
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

	emails, emailErr := authdomain.CreateEmailList(0, request.Contacts.Email)

	if emailErr != nil {
		return emailErr
	}

	phones, phoneErr := authdomain.CreatePhoneList(1, uc.LocationRepository, authdomain.ContactPhoneRequest(request.Contacts.Phone))

	if phoneErr != nil {
		return phoneErr
	}

	addresss, addressErr := authdomain.CreateAddressList(1, uc.LocationRepository, authdomain.ContactAddressRequest(request.Contacts.Address))

	if addressErr != nil {
		return addressErr
	}

	contacts := authdomain.CreateContacts(emails, phones, addresss)

	// Patient entity

	patient := authdomain.CreatePatient(person, nid, user, contacts)

	//Store patient

	if err := uc.PatientRepository.CreatePatient(patient); err != nil {
		return err
	}

	// TODO: Dispath PATIENT_CREATED_EVENT
	return nil
}
