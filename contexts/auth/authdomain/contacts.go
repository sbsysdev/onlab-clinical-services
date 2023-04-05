package authdomain

import (
	"errors"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
	"github.com/go-playground/validator/v10"
)

// Contact Email Value Object
type ContactEmail string

const (
	ERRORS_CONTACT_EMAIL_EMPTY shareddomain.DomainError = "ERRORS_CONTACT_EMAIL_EMPTY"
	ERRORS_CONTACT_EMAIL_MIN   shareddomain.DomainError = "ERRORS_CONTACT_EMAIL_MIN"
)

func CreateEmail(email string) (ContactEmail, error) {
	if len(email) == 0 {
		return ContactEmail(""), errors.New(string(ERRORS_CONTACT_EMAIL_EMPTY))
	}

	return ContactEmail(email), nil
}

func CreateEmailList(min uint8, emails ...string) ([]ContactEmail, error) {
	if len(emails) < int(min) {
		return []ContactEmail{}, errors.New(string(ERRORS_CONTACT_EMAIL_MIN))
	}

	emailList := make([]ContactEmail, len(emails))

	for _, v := range emails {
		email, err := CreateEmail(v)

		if err != nil {
			return []ContactEmail{}, err
		}

		emailList = append(emailList, email)
	}

	return emailList, nil
}

// Contact Phone Value Object
type ContactPhone struct {
	Country uint8
	Phone   string
}

const (
	ERRORS_CONTACT_PHONE_COUNTRY_NOT_FOUND shareddomain.DomainError = "ERRORS_CONTACT_PHONE_COUNTRY_NOT_FOUND"
	ERRORS_CONTACT_PHONE_FORMAT            shareddomain.DomainError = "ERRORS_CONTACT_PHONE_FORMAT"
	ERRORS_CONTACT_PHONE_MIN               shareddomain.DomainError = "ERRORS_CONTACT_PHONE_MIN"
)

func CreatePhone(country uint8, phone string, locationRepo LocationRepository) (ContactPhone, error) {
	if !locationRepo.IsExistingCountry(country) {
		return ContactPhone{}, errors.New(string(ERRORS_CONTACT_PHONE_COUNTRY_NOT_FOUND))
	}

	validate := validator.New()

	if err := validate.Var(phone, "min=7,max=10,numeric,excludes=.,excludes=0x2C"); err != nil {
		return ContactPhone{}, errors.New(string(ERRORS_CONTACT_PHONE_FORMAT))
	}

	return ContactPhone{
		Country: country,
		Phone:   phone,
	}, nil
}

func CreatePhoneList(min uint8, locationRepo LocationRepository, phones ...ContactPhone) ([]ContactPhone, error) {
	if len(phones) < int(min) {
		return []ContactPhone{}, errors.New(string(ERRORS_CONTACT_PHONE_MIN))
	}

	phoneList := make([]ContactPhone, len(phones))

	for _, v := range phones {
		phone, err := CreatePhone(v.Country, v.Phone, locationRepo)

		if err != nil {
			return []ContactPhone{}, err
		}

		phoneList = append(phoneList, phone)
	}

	return phoneList, nil
}

// Contact Address Value Object
type ContactAddress struct {
	Municipality uint16
	Address      string
}

const (
	ERRORS_CONTACT_ADDRESS_MUNICIPALITY_NOT_FOUND shareddomain.DomainError = "ERRORS_CONTACT_ADDRESS_MUNICIPALITY_NOT_FOUND"
	ERRORS_CONTACT_ADDRESS_EMPTY                  shareddomain.DomainError = "ERRORS_CONTACT_ADDRESS_EMPTY"
	ERRORS_CONTACT_ADDRESS_MIN                    shareddomain.DomainError = "ERRORS_CONTACT_ADDRESS_MIN"
)

func CreateAddress(municipality uint16, address string, locationRepo LocationRepository) (ContactAddress, error) {
	if !locationRepo.IsExistingMunicipality(municipality) {
		return ContactAddress{}, errors.New(string(ERRORS_CONTACT_ADDRESS_MUNICIPALITY_NOT_FOUND))
	}

	if len(address) == 0 {
		return ContactAddress{}, errors.New(string(ERRORS_CONTACT_ADDRESS_EMPTY))
	}

	return ContactAddress{
		Municipality: municipality,
		Address:      address,
	}, nil
}

func CreateAddressList(min uint8, locationRepo LocationRepository, addresses ...ContactAddress) ([]ContactAddress, error) {
	if len(addresses) < int(min) {
		return []ContactAddress{}, errors.New(string(ERRORS_CONTACT_ADDRESS_MIN))
	}

	addressList := make([]ContactAddress, len(addresses))

	for _, v := range addresses {
		address, err := CreateAddress(v.Municipality, v.Address, locationRepo)

		if err != nil {
			return []ContactAddress{}, err
		}

		addressList = append(addressList, address)
	}

	return addressList, nil
}

// Contacts Value Object
type Contacts struct {
	Emails    []ContactEmail
	Phones    []ContactPhone
	Addresses []ContactAddress
}

// Contacts Value Object Factory
func CreateContacts(emails []ContactEmail, phones []ContactPhone, addresses []ContactAddress) Contacts {
	return Contacts{
		Emails:    emails,
		Phones:    phones,
		Addresses: addresses,
	}
}
