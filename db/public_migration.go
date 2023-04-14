package db

import (
	"gorm.io/gorm"

	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbpublic"
	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbshared"
)

// Public migration
func PublicMigration(db *gorm.DB) {
	if err := dbpublic.MigratePublicPrerequisites(db); err != nil {
		panic(err)
	}

	if err := dbshared.MigrateSharedPrerequisites(db); err != nil {
		panic(err)
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
		panic(err)
	}
}
