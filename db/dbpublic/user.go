package dbpublic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbshared"
)

// Person data
type Person struct {
	Name    string                    `json:"name"`
	Surname string                    `json:"surname"`
	Birth   time.Time                 `json:"birth"`
	Sex     string                    `json:"sex"`
	Ind     dbshared.IdentityDocument `json:"ind"`
	Hpc     dbshared.IdentityDocument `json:"hpc"`
}

func (person *Person) Scan(v interface{}) error {
	bytes, ok := v.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", v))
	}

	return json.Unmarshal(bytes, &person)
}
func (Person) GormDataType() string {
	return "jsonb NOT NULL UNIQUE"
}
func (person Person) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	jsonValue, _ := json.Marshal(person)

	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{string(jsonValue)},
	}
}

// User data
type User struct {
	ID       string `gorm:"column:user_id;type:uuid NOT NULL UNIQUE PRIMARY KEY;default:gen_random_uuid();"`
	Name     string `gorm:"type:VARCHAR(64) NOT NULL UNIQUE"`
	Password string `gorm:"type:TEXT NOT NULL"`
	Person   Person
	Contacts dbshared.Contacts
	State    string          `gorm:"type:public.USER_STATE_ENUM NOT NULL;default:'unverified'"`
	Time     dbshared.TimeAt `gorm:"embedded"`

	SystemRoles []*Role `gorm:"many2many:user_role;"`

	UserRoles []*Role `gorm:"many2many:user_role_user;"`

	Organizations     []*Organization `gorm:"many2many:user_role_organization;"`
	OrganizationRoles []*Role         `gorm:"many2many:user_role_organization;"`
}
