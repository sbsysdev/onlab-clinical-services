package authinfra

import (
	"time"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbpublic"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbshared"
)

func FromPatientEntityToModels(patient authdomain.PatientEntity) (dbpublic.User, []dbpublic.UserRole) {
	// User
	user := dbpublic.User{
		ID:       patient.ID,
		Name:     string(patient.User.Name),
		Password: string(patient.User.Password),
		Person: dbpublic.Person{
			Name:    string(patient.Person.Name),
			Surname: string(patient.Person.Surname),
			Birth:   time.Time(patient.Person.Birth),
			Sex:     string(patient.Person.Sex),
			Nid: dbshared.IdentityDocument{
				Number: string(patient.NID),
			},
		},
		Contacts: dbshared.SingleContacts{
			Email: string(patient.Contacts.Email),
			Phone: dbshared.Phone{
				Country: patient.Contacts.Phone.Country.Id,
				Phone:   patient.Contacts.Phone.Phone,
			},
			Address: dbshared.Address{
				Municipality: patient.Contacts.Address.Municipality.Id,
				Address:      patient.Contacts.Address.Address,
			},
		},
		State: string(patient.User.State),
		Time: dbshared.TimeAt{
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}

	// User roles
	userRoles := make([]dbpublic.UserRole, len(patient.Roles))
	for i, role := range patient.Roles {
		userRoles[i] = dbpublic.UserRole{
			UserID: patient.ID,
			RoleID: role.ID,
		}
	}

	return user, userRoles
}
