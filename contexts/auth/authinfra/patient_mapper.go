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

func FromPatientModelToEntityFilled(userModel dbpublic.User, countryModel dbshared.Country, municipalityModel dbshared.Municipality, roleModels []dbpublic.Role) authdomain.PatientEntity {
	roles := make([]authdomain.AliasRoleEntity, len(roleModels))

	for i, roleEntity := range roleModels {
		roles[i] = authdomain.AliasRoleEntity{
			Alias: authdomain.RoleAlias(roleEntity.Alias),
			RoleEntity: authdomain.RoleEntity{
				ID:    roleEntity.ID,
				Name:  authdomain.RoleName(roleEntity.Name),
				Scope: authdomain.RoleScope(roleEntity.Scope),
				State: authdomain.RoleState(roleEntity.State),
			},
		}
	}

	return authdomain.PatientEntity{
		ID: userModel.ID,
		Person: authdomain.Person{
			Name:    authdomain.PersonName(userModel.Person.Name),
			Surname: authdomain.PersonSurname(userModel.Person.Surname),
			Birth:   authdomain.PersonBirth(userModel.Person.Birth),
			Sex:     authdomain.PersonSex(userModel.Person.Sex),
		},
		NID: authdomain.NationalIdentityDocumentNumber(userModel.Person.Nid.Number),
		User: authdomain.User{
			Name:     authdomain.UserName(userModel.Name),
			Password: authdomain.UserPassword(userModel.Password),
			State:    authdomain.UserState(userModel.State),
		},
		Contacts: authdomain.SingleContacts{
			Email: authdomain.ContactEmail(userModel.Contacts.Email),
			Phone: authdomain.ContactPhone{
				Country: authdomain.Country{
					Id:          countryModel.ID,
					Name:        countryModel.Name,
					Calling:     countryModel.Calling,
					Departments: []authdomain.Department{},
				},
				Phone: userModel.Contacts.Phone.Phone,
			},
			Address: authdomain.ContactAddress{
				Municipality: authdomain.Municipality{
					Id:   municipalityModel.ID,
					Name: municipalityModel.Name,
				},
				Address: userModel.Contacts.Address.Address,
			},
		},
		Roles: roles,
	}
}

func FromPatientModelToEntity(model dbpublic.User) authdomain.PatientEntity {
	return authdomain.PatientEntity{
		ID: model.ID,
		Person: authdomain.Person{
			Name:    authdomain.PersonName(model.Person.Name),
			Surname: authdomain.PersonSurname(model.Person.Surname),
			Birth:   authdomain.PersonBirth(model.Person.Birth),
			Sex:     authdomain.PersonSex(model.Person.Sex),
		},
		NID: authdomain.NationalIdentityDocumentNumber(model.Person.Nid.Number),
		User: authdomain.User{
			Name:     authdomain.UserName(model.Name),
			Password: authdomain.UserPassword(model.Password),
			State:    authdomain.UserState(model.State),
		},
		Contacts: authdomain.SingleContacts{
			Email: authdomain.ContactEmail(model.Contacts.Email),
			Phone: authdomain.ContactPhone{
				Country: authdomain.Country{
					Id: model.Contacts.Phone.Country,
				},
				Phone: model.Contacts.Phone.Phone,
			},
			Address: authdomain.ContactAddress{
				Municipality: authdomain.Municipality{
					Id: model.Contacts.Address.Municipality,
				},
				Address: model.Contacts.Address.Address,
			},
		},
	}
}
