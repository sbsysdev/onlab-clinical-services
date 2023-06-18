package authctrls

import (
	"time"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authapp"
)

type SignInPatientResponse struct {
	Token    string                 `json:"token"`
	Refresh  string                 `json:"refresh"`
	Person   PersonResponse         `json:"person"`
	Contacts SingleContactsResponse `json:"contacts"`
	User     UserResponse           `json:"user"`
	Roles    []AliasRoleResponse    `json:"roles"`
}

func FromSignInUseCaseResponseToResponse(response authapp.SignInPatientResponse) SignInPatientResponse {
	roles := make([]AliasRoleResponse, len(response.Patient.Roles))

	for i, aliasRole := range response.Patient.Roles {
		roles[i] = AliasRoleResponse{
			ID:    aliasRole.ID,
			Alias: string(aliasRole.Alias),
			Name:  aliasRole.Name,
		}
	}

	return SignInPatientResponse{
		Token:   response.Token,
		Refresh: response.RefreshToken,
		Person: PersonResponse{
			Name:    string(response.Patient.Person.Name),
			Surname: string(response.Patient.Person.Surname),
			NID:     string(response.Patient.NID),
			Birth:   time.Time(response.Patient.Person.Birth).Format("2006-01-02"),
			Sex:     string(response.Patient.Person.Sex),
		},
		Contacts: SingleContactsResponse{
			Email: string(response.Patient.Contacts.Email),
			Phone: PhoneResponse{
				Country: CountryResponse{
					ID:      response.Patient.Contacts.Phone.Country.Id,
					Name:    response.Patient.Contacts.Phone.Country.Name,
					Calling: response.Patient.Contacts.Phone.Country.Calling,
				},
				Phone: response.Patient.Contacts.Phone.Phone,
			},
			Address: AddressResponse{
				Municipality: MunicipalityResponse{
					ID:   response.Patient.Contacts.Address.Municipality.Id,
					Name: response.Patient.Contacts.Address.Municipality.Name,
				},
				Address:   response.Patient.Contacts.Address.Address,
				Latitude:  response.Patient.Contacts.Address.Latitude,
				Longitude: response.Patient.Contacts.Address.Longitude,
			},
		},
		User: UserResponse{
			ID:    response.Patient.ID,
			Name:  string(response.Patient.User.Name),
			State: string(response.Patient.User.State),
		},
		Roles: roles,
	}
}
