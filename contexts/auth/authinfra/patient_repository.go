package authinfra

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
)

type PatientRepository struct {
	DB *gorm.DB
}

func (repo *PatientRepository) CreatePatient(patient authdomain.PatientEntity) error {
	// Get db models
	user, err := FromPatientEntityToModels(patient)

	if err != nil {
		return err
	}

	// Store patient
	transactionErr := repo.DB.Transaction(func(tx *gorm.DB) error {
		tx.Save(&user)

		return nil
	}, &sql.TxOptions{})

	if transactionErr != nil {
		return transactionErr
	}

	return nil
}
