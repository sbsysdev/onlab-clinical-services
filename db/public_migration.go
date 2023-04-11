package db

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbpublic"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbshared"
)

// Public migration
func PublicMigration(db *gorm.DB) {
	if err := dbshared.MigrateSharedPrerequisites(db); err != nil {
		fmt.Sprintln(err.Error())
	}

	if err := dbpublic.MigratePublicPrerequisites(db); err != nil {
		fmt.Sprintln(err.Error())
	}

	// Migrate schemas
	if err := db.AutoMigrate(
		// Role echemas
		&dbshared.Operation{},
		&dbshared.Permission{},
		&dbshared.PermissionOperation{},

		&dbpublic.Role{},
		&dbshared.RolePermission{},

		// User schemas
		&dbpublic.User{},
		&dbpublic.UserRole{},
		&dbpublic.UserRoleUser{},

		// Organization schemas
		&dbpublic.Organization{},
		&dbpublic.UserRoleOrganization{},
	); err != nil {
		fmt.Sprintln(err.Error())
	}
}
