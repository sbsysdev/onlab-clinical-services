package authinfra

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
)

type PatientRepository struct {
	DB *gorm.DB
}

func (repo PatientRepository) CreatePatient(patient authdomain.PatientEntity) error {
	// Get db models
	user, userRoles, err := FromPatientEntityToModels(patient)

	if err != nil {
		return err
	}

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
		tx.Rollback()

		return err.Error
	}
	if usernameCoincidences > 0 {
		return errors.New(string(authdomain.ERRORS_USER_NAME_NOT_AVAILABLE))
	}

	// Validate NID
	var nidCoincidences int64 = 0
	if err := tx.Table("users").Where(fmt.Sprintf(`person @> '{"nid":{"number":"%s"}}'`, user.Person.Nid.Number)).Count(&nidCoincidences); err.Error != nil {
		tx.Rollback()

		return err.Error
	}
	if nidCoincidences > 0 {
		return errors.New(string(authdomain.ERRORS_NID_NOT_AVAILABLE))
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
