package dbpublic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/OnLab-Clinical/onlab-clinical-services/db/dbshared"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Organization info data
type OrganizationInfo struct {
	Type    string                    `json:"type"`
	RUC     dbshared.IdentityDocument `json:"ruc"`
	Logo    string                    `json:"logo"`
	ISOType string                    `json:"isotype"`
	Slogan  string                    `json:"slogan"`
	Mission string                    `json:"mission"`
	Vision  string                    `json:"vision"`
}

func (info *OrganizationInfo) Scan(v interface{}) error {
	bytes, ok := v.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", v))
	}

	return json.Unmarshal(bytes, &info)
}
func (OrganizationInfo) GormDataType() string {
	return "jsonb NOT NULL UNIQUE"
}
func (info OrganizationInfo) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	jsonValue, _ := json.Marshal(info)

	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{string(jsonValue)},
	}
}

// Organization data
type Organization struct {
	ID       string `gorm:"column:organization_id;type:uuid NOT NULL UNIQUE PRIMARY KEY;default:gen_random_uuid();"`
	Name     string `gorm:"type:VARCHAR(64) NOT NULL;"`
	Domain   string `gorm:"type:VARCHAR(64) NOT NULL UNIQUE;"`
	Schema   string `gorm:"type:VARCHAR(64) NOT NULL UNIQUE;"`
	Info     OrganizationInfo
	Contacts dbshared.Contacts
	State    string          `gorm:"type:public.ORG_STATE_ENUM NOT NULL;default:'unverified'"`
	Time     dbshared.TimeAt `gorm:"embedded"`

	Users []*User `gorm:"many2many:user_role_organization;"`
	Roles []*Role `gorm:"many2many:user_role_organization;"`
}
