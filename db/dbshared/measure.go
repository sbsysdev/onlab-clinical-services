package dbshared

// Measure data
type Measure struct {
	ID    string        `gorm:"column:measure_id;type:uuid;not null;unique;primaryKey;default:gen_random_uuid()"`
	Name  MultiLanguage `gorm:"not null;unique"`
	Short MultiLanguage `gorm:"not null;unique"`
}
