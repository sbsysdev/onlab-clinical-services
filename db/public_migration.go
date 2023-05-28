package db

import (
	"fmt"

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

	// Role echemas
	if err := db.AutoMigrate(
		&dbshared.Operation{},
	); err != nil {
		fmt.Println(err.Error())
	}
	if err := db.AutoMigrate(
		&dbshared.Permission{},
	); err != nil {
		fmt.Println(err.Error())
	}
	if err := db.AutoMigrate(
		&dbshared.PermissionOperation{},
	); err != nil {
		fmt.Println(err.Error())
	}

	if err := db.AutoMigrate(
		&dbpublic.Role{},
	); err != nil {
		fmt.Println(err.Error())
	}
	if err := db.AutoMigrate(
		&dbshared.RolePermission{},
	); err != nil {
		fmt.Println(err.Error())
	}

	// User schemas
	if err := db.AutoMigrate(
		&dbpublic.User{},
	); err != nil {
		fmt.Println(err.Error())
	}
	if err := db.AutoMigrate(
		&dbpublic.UserRole{},
	); err != nil {
		fmt.Println(err.Error())
	}
	if err := db.AutoMigrate(
		&dbpublic.UserRoleUser{},
	); err != nil {
		fmt.Println(err.Error())
	}

	// Organization schemas
	if err := db.AutoMigrate(
		&dbpublic.Organization{},
	); err != nil {
		fmt.Println(err.Error())
	}
	if err := db.AutoMigrate(
		&dbpublic.UserRoleOrganization{},
	); err != nil {
		fmt.Println(err.Error())
	}

	// Migrate public seeds
	if err := dbpublic.MigratePublicSystemRoles(db); err != nil {
		panic(err)
	}
	if err := dbpublic.MigratePublicUserRoles(db); err != nil {
		panic(err)
	}

	// Migrate location seed
	if err := db.AutoMigrate(
		&dbshared.Country{},
	); err != nil {
		fmt.Println(err.Error())
	}

	if err := db.AutoMigrate(
		&dbshared.Department{},
	); err != nil {
		fmt.Println(err.Error())
	}

	if err := db.AutoMigrate(
		&dbshared.Municipality{},
	); err != nil {
		fmt.Println(err.Error())
	}

	if err := dbshared.MigrateSharedLocation(db); err != nil {
		panic(err)
	}
}
