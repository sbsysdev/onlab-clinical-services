package authinfra

import (
	"time"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbpublic"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbshared"
)

func FromPatientEntityToModels(patient authdomain.PatientEntity) (dbpublic.User, error) {
	// Map contacts

	// Contact emails
	emails := make([]string, len(patient.Contacts.Emails))
	for i, v := range patient.Contacts.Emails {
		emails[i] = string(v)
	}

	// Contact phones
	phones := make([]dbshared.Phone, len(patient.Contacts.Phones))
	for i, v := range patient.Contacts.Phones {
		phones[i] = dbshared.Phone{
			Country: v.Country.Id,
			Phone:   v.Phone,
		}
	}

	// Contact addresses
	addresses := make([]dbshared.Address, len(patient.Contacts.Addresses))
	for i, v := range patient.Contacts.Addresses {
		addresses[i] = dbshared.Address{
			Municipality: v.Municipality.Id,
			Address:      v.Address,
		}
	}

	return dbpublic.User{
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
		Contacts: dbshared.Contacts{
			Emails:    emails,
			Phones:    phones,
			Addresses: addresses,
		},
		State: string(patient.User.State),
		Time: dbshared.TimeAt{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}
