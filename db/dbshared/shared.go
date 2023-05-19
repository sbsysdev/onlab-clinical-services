package dbshared

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Shared prerequisites
func MigrateSharedPrerequisites(db *gorm.DB) error {
	if err := db.Exec(`DO $$ BEGIN
	IF NOT EXISTS (SELECT FROM pg_type WHERE typname ILIKE 'COMMON_SCOPE_ENUM') THEN
	CREATE TYPE COMMON_SCOPE_ENUM AS ENUM('system', 'user', 'org', 'branch');
	END IF;
	END$$;`).Error; err != nil {
		return err
	}

	if err := db.Exec(`DO $$ BEGIN
	IF NOT EXISTS (SELECT FROM pg_type WHERE typname ILIKE 'COMMON_STATE_ENUM') THEN
	CREATE TYPE COMMON_STATE_ENUM AS ENUM('active', 'inactive');
	END IF;
	END$$;`).Error; err != nil {
		return err
	}

	return nil
}

// Multilanguage data
type MultiLanguage map[string]string

func (lang *MultiLanguage) Scan(v interface{}) error {
	bytes, ok := v.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", v))
	}

	return json.Unmarshal(bytes, &lang)
}
func (MultiLanguage) GormDataType() string {
	return "jsonb NOT NULL UNIQUE"
}
func (lang MultiLanguage) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	jsonValue, _ := json.Marshal(lang)

	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{string(jsonValue)},
	}
}

// Time data
type TimeAt struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Identity document data
type IdentityDocument struct {
	Number       string `json:"number"`
	FrontPicture string `json:"front"`
	BackPicture  string `json:"back"`
}

// Contacts data
type Phone struct {
	Country uint8  `json:"country"`
	Phone   string `json:"phone"`
}
type Address struct {
	Municipality uint16 `json:"municipality"`
	Address      string `json:"address"`
}
type Contacts struct {
	Emails    []string  `json:"emails"`
	Phones    []Phone   `json:"phones"`
	Addresses []Address `json:"addresses"`
}

func (contacts *Contacts) Scan(v interface{}) error {
	bytes, ok := v.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", v))
	}

	return json.Unmarshal(bytes, &contacts)
}
func (Contacts) GormDataType() string {
	return "jsonb"
}
func (contacts Contacts) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	jsonValue, _ := json.Marshal(contacts)

	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{string(jsonValue)},
	}
}
