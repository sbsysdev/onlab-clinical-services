package dbshared

// Measure data
type Measure struct {
	ID    string `gorm:"column:measure_id;type:uuid NOT NULL UNIQUE PRIMARY KEY;default:gen_random_uuid();"`
	Name  MultiLanguage
	Short MultiLanguage
}
