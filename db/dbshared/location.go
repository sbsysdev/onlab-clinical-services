package dbshared

// Country data
type Country struct {
	ID   uint   `gorm:"column:country_id;type:SMALLSERIAL PRIMARY KEY;"`
	Name string `gorm:"type:VARCHAR(64) NOT NULL UNIQUE"`
	Time TimeAt `gorm:"embedded"`
}

// Department data
type Department struct {
	ID        uint    `gorm:"column:department_id;type:SMALLSERIAL PRIMARY KEY;"`
	CountryID uint    `gorm:"column:country_id;type:SMALLINT;"`
	Country   Country `gorm:"foreignKey:CountryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name      string  `gorm:"type:VARCHAR(64) NOT NULL UNIQUE"`
	Time      TimeAt  `gorm:"embedded"`
}

// Municipality data
type Municipality struct {
	ID           uint       `gorm:"column:municipality_id;type:SMALLSERIAL PRIMARY KEY;"`
	DepartmentID uint       `gorm:"column:department_id;type:SMALLINT;"`
	Department   Department `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name         string     `gorm:"type:VARCHAR(64) NOT NULL UNIQUE"`
	Time         TimeAt     `gorm:"embedded"`
}
