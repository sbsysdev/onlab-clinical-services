package dbpublic

import (
	"errors"
	"fmt"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbshared"
	"gorm.io/gorm"
)

// Public Role data
type Role struct {
	Alias string `gorm:"column:alias;type:VARCHAR(64) NOT NULL UNIQUE;index;"`
	dbshared.Role

	SystemUsers []*User `gorm:"many2many:user_role;"`

	UserUsers []*User `gorm:"many2many:user_role_user;"`

	Organizations     []*Organization `gorm:"many2many:user_role_organization;"`
	OrganizationUsers []*User         `gorm:"many2many:user_role_organization;"`
}

// Migrate public system roles
func MigratePublicSystemRoles(db *gorm.DB) error {
	roles := []Role{
		{
			Alias: string(authdomain.ALIAS_PATIENT),
			Role: dbshared.Role{
				ID: "1a631fc7-4fbc-4361-a9b0-7a5f15679203",
				Name: map[string]string{
					"en": "Patient",
					"es": "Paciente",
				},
				Scope: string(authdomain.SCOPE_SYSTEM),
				State: string(authdomain.ROLE_STATE_ACTIVE),
			},
		},
		{
			Alias: string(authdomain.ALIAS_OWNER),
			Role: dbshared.Role{
				ID: "82f0ead3-edd2-4437-86d6-9208ce3f85c4",
				Name: map[string]string{
					"en": "Owner",
					"es": "Propietario",
				},
				Scope: string(authdomain.SCOPE_SYSTEM),
				State: string(authdomain.ROLE_STATE_ACTIVE),
			},
		},
		{
			Alias: string(authdomain.ALIAS_COLLABORATOR),
			Role: dbshared.Role{
				ID: "1d64222c-fa18-4c27-8c7e-a0406684ad63",
				Name: map[string]string{
					"en": "Collaborator",
					"es": "Colaborador",
				},
				Scope: string(authdomain.SCOPE_SYSTEM),
				State: string(authdomain.ROLE_STATE_ACTIVE),
			},
		},
		{
			Alias: string(authdomain.ALIAS_PARENT),
			Role: dbshared.Role{
				ID: "982b4857-48a8-4e7b-9b4b-6527cf26c046",
				Name: map[string]string{
					"en": "Parent",
					"es": "Padre",
				},
				Scope: string(authdomain.SCOPE_USER),
				State: string(authdomain.ROLE_STATE_ACTIVE),
			},
		},
	}

	for _, v := range roles {
		if err := db.Save(&v).Error; err != nil {
			fmt.Println(errors.Is(err, gorm.ErrDuplicatedKey))
			return err
		}
	}

	return nil
}
