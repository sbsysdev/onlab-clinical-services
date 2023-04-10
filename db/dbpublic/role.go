package dbpublic

import "github.com/OnLab-Clinical/onlab-clinical-services/db/dbshared"

// Public Role data
type Role struct {
	dbshared.Role
	Alias string `gorm:"column:alias;type:VARCHAR(64) NOT NULL UNIQUE;"`

	SystemUsers []*User `gorm:"many2many:user_role;"`

	UserUsers []*User `gorm:"many2many:user_role_user;"`

	Organizations     []*Organization `gorm:"many2many:user_role_organization;"`
	OrganizationUsers []*User         `gorm:"many2many:user_role_organization;"`
}
