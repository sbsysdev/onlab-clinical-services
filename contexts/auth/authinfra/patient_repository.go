package authinfra

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbpublic"
)

type PatientRepository struct {
	// Postgresql connection
	DB *gorm.DB
	// Repositories
	LocationRepository LocationRepository
	RoleRepository     RoleRepository
}

func (repo PatientRepository) CreatePatient(patient authdomain.PatientEntity) error {
	// Get db models
	user, userRoles := FromPatientEntityToModels(patient)

	// Store patient
	tx := repo.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Exec("SET search_path=public;").Error; err != nil {
		return err
	}

	// Validate user name
	var usernameCoincidences int64 = 0
	if err := tx.Table("users").Where("name = ?", user.Name).Count(&usernameCoincidences); err.Error != nil {
		return err.Error
	}
	if usernameCoincidences > 0 {
		return errors.New(string(authdomain.ERRORS_USER_NAME_NOT_AVAILABLE))
	}

	// Validate NID
	var nidCoincidences int64 = 0
	if err := tx.Table("users").Where(fmt.Sprintf(`person @> '{"nid":{"number":"%s"}}'`, user.Person.Nid.Number)).Count(&nidCoincidences); err.Error != nil {
		return err.Error
	}
	if nidCoincidences > 0 {
		return errors.New(string(authdomain.ERRORS_NID_NOT_AVAILABLE))
	}

	// Validate Email
	var emailCoincidences int64 = 0
	if err := tx.Table("users").Where(fmt.Sprintf(`contacts @> '{"email":"%s"}'`, user.Contacts.Email)).Count(&emailCoincidences); err.Error != nil {
		return err.Error
	}
	if emailCoincidences > 0 {
		return errors.New(string(authdomain.ERRORS_CONTACT_EMAIL_NOT_AVAILABLE))
	}

	// Validate Phone
	var phoneCoincidences int64 = 0
	if err := tx.Table("users").Where(fmt.Sprintf(`contacts @> '{"phone":{"phone":"%s"}}'`, user.Contacts.Phone.Phone)).Count(&phoneCoincidences); err.Error != nil {
		return err.Error
	}
	if phoneCoincidences > 0 {
		return errors.New(string(authdomain.ERRORS_CONTACT_PHONE_NOT_AVAILABLE))
	}

	// Store user
	if txErr := tx.Save(&user).Error; txErr != nil {
		tx.Rollback()

		return txErr
	}

	// Add roles to stored user
	for _, userRole := range userRoles {
		if txErr := tx.Save(&userRole).Error; txErr != nil {
			tx.Rollback()

			return txErr
		}
	}

	// Try to commit all changes
	return tx.Commit().Error
}

func (repo PatientRepository) ReadPatientByName(name string) (authdomain.PatientEntity, error) {
	var user dbpublic.User

	if err := repo.DB.Table("users").Preload("UserRoles").Preload("SystemRoles").First(&user, "name = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return authdomain.PatientEntity{}, errors.New(string(authdomain.ERRORS_USER_NAME_NOT_FOUND))
		}

		return authdomain.PatientEntity{}, err
	}

	country, countryErr := repo.LocationRepository.GetCountryModelById(user.Contacts.Phone.Country)

	if countryErr != nil {
		return authdomain.PatientEntity{}, countryErr
	}

	municipality, municipalityErr := repo.LocationRepository.GetMunicipalityModelById(user.Contacts.Address.Municipality)

	if municipalityErr != nil {
		return authdomain.PatientEntity{}, municipalityErr
	}

	aliases := make([]authdomain.RoleAlias, len(user.SystemRoles)+len(user.UserRoles))

	for i, sysRole := range user.SystemRoles {
		aliases[i] = authdomain.RoleAlias(sysRole.Alias)
	}

	for i, userRole := range user.UserRoles {
		aliases[len(user.SystemRoles)+i] = authdomain.RoleAlias(userRole.Alias)
	}

	roles, roleErr := repo.RoleRepository.GetAliasRoleModelsByAlias(aliases)

	if roleErr != nil {
		return authdomain.PatientEntity{}, roleErr
	}

	return FromPatientModelToEntityFilled(user, country, municipality, roles), nil
}

func (repo PatientRepository) ReadPatientById(patientId string) (authdomain.PatientEntity, error) {
	var founded dbpublic.User

	if err := repo.DB.Table("users").First(&founded, "user_id = ?", patientId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return authdomain.PatientEntity{}, errors.New(string(authdomain.ERRORS_USER_NOT_FOUND))
		}

		return authdomain.PatientEntity{}, err
	}

	return FromPatientModelToEntity(founded), nil
}

func (repo PatientRepository) ReadPatientByEmail(email string) (authdomain.PatientEntity, error) {
	var founded dbpublic.User

	if err := repo.DB.Table("users").First(&founded, fmt.Sprintf(`contacts @> '{"email":"%s"}'`, email)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return authdomain.PatientEntity{}, errors.New(string(authdomain.ERRORS_USER_NOT_FOUND))
		}

		return authdomain.PatientEntity{}, err
	}

	return FromPatientModelToEntity(founded), nil
}
