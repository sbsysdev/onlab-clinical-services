package authinfra

import (
	"gorm.io/gorm"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
)

type PatientRepository struct {
	DB *gorm.DB
}

func (repo PatientRepository) CreatePatient(patient authdomain.PatientEntity) error {
	// Get db models
	user, err := FromPatientEntityToModels(patient)

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

	if txErr := tx.Save(&user).Error; txErr != nil {
		tx.Rollback()

		return txErr
	}

	return tx.Commit().Error
}
