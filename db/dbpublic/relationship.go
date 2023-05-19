package dbpublic

import (
	"gorm.io/gorm"
)

// Public prerequisites
func MigratePublicPrerequisites(db *gorm.DB) error {
	if err := db.Exec("CREATE SCHEMA IF NOT EXISTS public;"); err.Error != nil {
		return err.Error
	}

	if err := db.Exec("SET search_path=public;"); err.Error != nil {
		return err.Error
	}

	if err := db.Exec(`DO $$ BEGIN
	IF NOT EXISTS (SELECT FROM pg_type WHERE typname ILIKE 'USER_STATE_ENUM') THEN
	CREATE TYPE USER_STATE_ENUM AS ENUM('unverified', 'blocked', 'verified', 'suspended');
	END IF;
	END$$;`).Error; err != nil {
		return err
	}

	if err := db.Exec(`DO $$ BEGIN
	IF NOT EXISTS (SELECT FROM pg_type WHERE typname ILIKE 'ORG_STATE_ENUM') THEN
	CREATE TYPE ORG_STATE_ENUM AS ENUM('unverified', 'refused', 'verified', 'suspended');
	END IF;
	END$$;`).Error; err != nil {
		return err
	}

	return nil
}

// User-Role data
type UserRole struct {
	UserID string `gorm:"column:user_id;type:uuid NOT NULL;primaryKey;"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID string `gorm:"column:role_id;type:uuid NOT NULL;primaryKey;"`
	Role   Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (UserRole) TableName() string {
	return "user_role"
}

// User-Role-User data
type UserRoleUser struct {
	UserID          string `gorm:"column:user_id;type:uuid NOT NULL;primaryKey;"`
	User            User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID          string `gorm:"column:role_id;type:uuid NOT NULL;primaryKey;"`
	Role            Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DependentUserID string `gorm:"column:dependent_user_id;type:uuid NOT NULL;primaryKey;"`
	DependentUser   User   `gorm:"foreignKey:DependentUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (UserRoleUser) TableName() string {
	return "user_role_user"
}

// User-Role-User data
type UserRoleOrganization struct {
	UserID         string       `gorm:"column:user_id;type:uuid NOT NULL;primaryKey;"`
	User           User         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleID         string       `gorm:"column:role_id;type:uuid NOT NULL;primaryKey;"`
	Role           Role         `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrganizationID string       `gorm:"column:organization_id;type:uuid NOT NULL;primaryKey;"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (UserRoleOrganization) TableName() string {
	return "user_role_organization"
}
