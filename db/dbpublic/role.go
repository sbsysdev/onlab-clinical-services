package dbpublic

import (
	"time"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbshared"
	"gorm.io/gorm"
)

// Public Role data
type Role struct {
	Alias         string `gorm:"column:alias;type:VARCHAR(64);not null;unique;index"`
	dbshared.Role `gorm:"embedded"`

	SystemUsers []*User `gorm:"many2many:user_role"`

	UserUsers []*User `gorm:"many2many:user_role_user"`

	Organizations     []*Organization `gorm:"many2many:user_role_organization"`
	OrganizationUsers []*User         `gorm:"many2many:user_role_organization"`
}

// Migrate public system roles
func MigratePublicSystemRoles(db *gorm.DB) error {
	roles := []Role{
		{
			Alias: string(authdomain.ALIAS_PATIENT),
			Role: dbshared.Role{
				Name: map[string]string{
					"en": "Patient",
					"es": "Paciente",
				},
				Scope: string(authdomain.SCOPE_SYSTEM),
				State: string(authdomain.ROLE_STATE_ACTIVE),
				Time: dbshared.TimeAt{
					CreatedAt: time.Date(2023, time.March, 24, 20, 44, 0, 0, time.UTC),
				},
			},
		},
		{
			Alias: string(authdomain.ALIAS_OWNER),
			Role: dbshared.Role{
				Name: map[string]string{
					"en": "Owner",
					"es": "Propietario",
				},
				Scope: string(authdomain.SCOPE_SYSTEM),
				State: string(authdomain.ROLE_STATE_ACTIVE),
				Time: dbshared.TimeAt{
					CreatedAt: time.Date(2023, time.March, 24, 20, 44, 0, 0, time.UTC),
				},
			},
		},
		{
			Alias: string(authdomain.ALIAS_COLLABORATOR),
			Role: dbshared.Role{
				Name: map[string]string{
					"en": "Collaborator",
					"es": "Colaborador",
				},
				Scope: string(authdomain.SCOPE_SYSTEM),
				State: string(authdomain.ROLE_STATE_ACTIVE),
				Time: dbshared.TimeAt{
					CreatedAt: time.Date(2023, time.March, 24, 20, 44, 0, 0, time.UTC),
				},
			},
		},
	}

	for _, v := range roles {
		if err := db.Save(&v).Error; err != nil {
			return err
		}
	}

	return nil
}

// Migrate public user roles
func MigratePublicUserRoles(db *gorm.DB) error {
	roles := []Role{
		{
			Alias: string(authdomain.ALIAS_PARENT),
			Role: dbshared.Role{
				Name: map[string]string{
					"en": "Parent",
					"es": "Padre",
				},
				Scope: string(authdomain.SCOPE_USER),
				State: string(authdomain.ROLE_STATE_ACTIVE),
				Time: dbshared.TimeAt{
					CreatedAt: time.Date(2023, time.March, 24, 20, 44, 0, 0, time.UTC),
				},
			},
		},
	}

	for _, v := range roles {
		if err := db.Save(&v).Error; err != nil {
			return err
		}
	}

	return nil
}
