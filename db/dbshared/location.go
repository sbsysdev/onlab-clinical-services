package dbshared

import (
	"time"

	"gorm.io/gorm"
)

// Country data
type Country struct {
	ID          uint         `gorm:"column:country_id;type:SMALLSERIAL PRIMARY KEY;"`
	Name        string       `gorm:"type:VARCHAR(64) NOT NULL UNIQUE"`
	Calling     string       `gorm:"type:VARCHAR(4) NOT NULL UNIQUE"`
	Departments []Department `gorm:"foreignKey:CountryID;references:ID"`
	Time        TimeAt       `gorm:"embedded"`
}

// Department data
type Department struct {
	ID             uint           `gorm:"column:department_id;type:SMALLSERIAL PRIMARY KEY;"`
	CountryID      uint           `gorm:"column:country_id;type:SMALLINT;"`
	Country        Country        `gorm:"foreignKey:CountryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name           string         `gorm:"type:VARCHAR(64) NOT NULL UNIQUE"`
	Municipalities []Municipality `gorm:"foreignKey:DepartmentID;references:ID"`
	Time           TimeAt         `gorm:"embedded"`
}

// Municipality data
type Municipality struct {
	ID           uint       `gorm:"column:municipality_id;type:SMALLSERIAL PRIMARY KEY;"`
	DepartmentID uint       `gorm:"column:department_id;type:SMALLINT;"`
	Department   Department `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name         string     `gorm:"type:VARCHAR(64) NOT NULL UNIQUE"`
	Time         TimeAt     `gorm:"embedded"`
}

// Migrate shared location
func MigrateSharedLocation(db *gorm.DB) error {
	countries := []Country{
		{
			ID:      1,
			Name:    "Nicaragua",
			Calling: "+505",
			Departments: []Department{
				{
					ID:   1,
					Name: "Boaco",
					Municipalities: []Municipality{
						{
							ID:   1,
							Name: "Boaco",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   2,
							Name: "Camoapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   3,
							Name: "San José de los Remates",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   4,
							Name: "San Lorenzo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   5,
							Name: "Santa Lucía",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   6,
							Name: "Teustepe",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   2,
					Name: "Carazo",
					Municipalities: []Municipality{
						{
							ID:   7,
							Name: "Diriamba",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   8,
							Name: "Dolores",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   9,
							Name: "El Rosario",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   10,
							Name: "Jinotepe",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   11,
							Name: "La Conquista",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   12,
							Name: "La Paz de Oriente",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   13,
							Name: "San Marcos",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   14,
							Name: "Santa Teresa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   3,
					Name: "Chinandega",
					Municipalities: []Municipality{
						{
							ID:   15,
							Name: "Chichigalpa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   16,
							Name: "Chinandega",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   17,
							Name: "Cinco Pinos",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   18,
							Name: "Corinto",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   19,
							Name: "El Realejo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   20,
							Name: "El Viejo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   21,
							Name: "Posoltega",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   22,
							Name: "Puerto Morazán",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   23,
							Name: "San Francisco del Norte",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   24,
							Name: "San Pedro del Norte",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   25,
							Name: "Santo Tomás del Norte",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   26,
							Name: "Somotillo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   27,
							Name: "Villanueva",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   4,
					Name: "Chontales",
					Municipalities: []Municipality{
						{
							ID:   28,
							Name: "Acoyapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   29,
							Name: "Comalapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   30,
							Name: "Cuapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   31,
							Name: "El Coral",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   32,
							Name: "Juigalpa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   33,
							Name: "La Libertad",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   34,
							Name: "San Pedro de Lóvago",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   35,
							Name: "Santo Domingo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   36,
							Name: "Santo Tomás",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   37,
							Name: "Villa Sandino",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   5,
					Name: "Costa Caribe Norte",
					Municipalities: []Municipality{
						{
							ID:   38,
							Name: "Bonanza",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   39,
							Name: "Mulukukú",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   40,
							Name: "Prinzapolka",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   41,
							Name: "Puerto Cabezas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   42,
							Name: "Rosita",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   43,
							Name: "Siuna",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   44,
							Name: "Waslala",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   45,
							Name: "Waspán",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   6,
					Name: "Costa Caribe Sur",
					Municipalities: []Municipality{
						{
							ID:   46,
							Name: "Bluefields",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   47,
							Name: "Corn Island",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   48,
							Name: "Desembocadura de Río Grande",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   49,
							Name: "El Ayote",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   50,
							Name: "El Rama",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   51,
							Name: "El Tortuguero",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   52,
							Name: "Kukra Hill",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   53,
							Name: "La Cruz de Río Grande",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   54,
							Name: "Laguna de Perlas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   55,
							Name: "Muelle de los Bueyes",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   56,
							Name: "Nueva Guinea",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   57,
							Name: "Paiwas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   7,
					Name: "Estelí",
					Municipalities: []Municipality{
						{
							ID:   58,
							Name: "Condega",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   59,
							Name: "Estelí",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   60,
							Name: "La Trinidad",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   61,
							Name: "Pueblo Nuevo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   62,
							Name: "San Juan de Limay",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   63,
							Name: "San Nicolás",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   8,
					Name: "Granada",
					Municipalities: []Municipality{
						{
							ID:   64,
							Name: "Diriá",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   65,
							Name: "Diriomo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   66,
							Name: "Granada",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   67,
							Name: "Nandaime",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   9,
					Name: "Jinotega",
					Municipalities: []Municipality{
						{
							ID:   68,
							Name: "El Cuá",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   69,
							Name: "Jinotega",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   70,
							Name: "La Concordia",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   71,
							Name: "San José de Bocay",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   72,
							Name: "San Rafael del Norte",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   73,
							Name: "San Sebastián de Yalí",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   74,
							Name: "Santa María de Pantasma",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   75,
							Name: "Wiwilí de Jinotega",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   10,
					Name: "León",
					Municipalities: []Municipality{
						{
							ID:   76,
							Name: "Achuapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   77,
							Name: "El Jicaral",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   78,
							Name: "El Sauce",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   79,
							Name: "La Paz Centro",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   80,
							Name: "Larreynaga",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   81,
							Name: "León",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   82,
							Name: "Nagarote",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   83,
							Name: "Quezalguaque",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   84,
							Name: "Santa Rosa del Peñón",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   85,
							Name: "Telica",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   11,
					Name: "Madriz",
					Municipalities: []Municipality{
						{
							ID:   86,
							Name: "Las Sabanas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   87,
							Name: "Palacagüina",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   88,
							Name: "San José de Cusmapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   89,
							Name: "San Juan de Río Coco",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   90,
							Name: "San Lucas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   91,
							Name: "Somoto",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   92,
							Name: "Telpaneca",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   93,
							Name: "Totogalpa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   94,
							Name: "Yalagüina",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   12,
					Name: "Managua",
					Municipalities: []Municipality{
						{
							ID:   95,
							Name: "Ciudad Sandino",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   96,
							Name: "El Crucero",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   97,
							Name: "Managua",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   98,
							Name: "Mateare",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   99,
							Name: "San Francisco Libre",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   100,
							Name: "San Rafael del Sur",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   101,
							Name: "Ticuantepe",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   102,
							Name: "Tipitapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   103,
							Name: "Villa El Carmen",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   13,
					Name: "Masaya",
					Municipalities: []Municipality{
						{
							ID:   104,
							Name: "Catarina",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   105,
							Name: "La Concepción",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   106,
							Name: "Masatepe",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   107,
							Name: "Masaya",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   108,
							Name: "Nandasmo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   109,
							Name: "Nindirí",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   110,
							Name: "Niquinohomo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   111,
							Name: "San Juan de Oriente",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   112,
							Name: "Tisma",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   14,
					Name: "Matagalpa",
					Municipalities: []Municipality{
						{
							ID:   113,
							Name: "Ciudad Darío",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   114,
							Name: "El Tuma - La Dalia",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   115,
							Name: "Esquipulas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   116,
							Name: "Matagalpa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   117,
							Name: "Matiguás",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   118,
							Name: "Muy Muy",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   119,
							Name: "Rancho Grande",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   120,
							Name: "Río Blanco",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   121,
							Name: "San Dionisio",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   122,
							Name: "San Isidro",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   123,
							Name: "San Ramón",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   124,
							Name: "Sébaco",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   125,
							Name: "Terrabona",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   15,
					Name: "Nueva Segovia",
					Municipalities: []Municipality{
						{
							ID:   126,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   127,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   128,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   129,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   130,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   131,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   132,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   133,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   134,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   135,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   136,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   137,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   16,
					Name: "Río San Juan",
					Municipalities: []Municipality{
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
				{
					ID:   17,
					Name: "Rivas",
					Municipalities: []Municipality{
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							ID:   0,
							Name: "",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
					},
					Time: TimeAt{
						CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
					},
				},
			},
			Time: TimeAt{
				CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
			},
		},
	}

	for _, v := range countries {
		if err := db.Save(&v).Error; err != nil {
			return err
		}
	}

	return nil
}
