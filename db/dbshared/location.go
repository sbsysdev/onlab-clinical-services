package dbshared

import (
	"time"

	"gorm.io/gorm"
)

// Country data
type Country struct {
	ID          string       `gorm:"column:country_id;type:uuid;not null;unique;primaryKey;default:gen_random_uuid()"`
	Name        string       `gorm:"type:VARCHAR(64);not null;unique"`
	Calling     string       `gorm:"type:VARCHAR(4);not null;unique"`
	Departments []Department `gorm:"foreignKey:CountryID;references:ID"`
	Time        TimeAt       `gorm:"embedded"`
}

// Department data
type Department struct {
	ID             string         `gorm:"column:department_id;type:uuid;not null;unique;primaryKey;default:gen_random_uuid()"`
	CountryID      string         `gorm:"column:country_id;type:uuid"`
	Country        Country        `gorm:"foreignKey:CountryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name           string         `gorm:"type:VARCHAR(64);not null;unique"`
	Municipalities []Municipality `gorm:"foreignKey:DepartmentID;references:ID"`
	Time           TimeAt         `gorm:"embedded"`
}

// Municipality data
type Municipality struct {
	ID           string     `gorm:"column:municipality_id;type:uuid;not null;unique;primaryKey;default:gen_random_uuid()"`
	DepartmentID string     `gorm:"column:department_id;type:uuid"`
	Department   Department `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name         string     `gorm:"type:VARCHAR(64);not null;unique"`
	Time         TimeAt     `gorm:"embedded"`
}

// Migrate shared location
func MigrateSharedLocation(db *gorm.DB) error {
	countries := []Country{
		{
			Name:    "Nicaragua",
			Calling: "+505",
			Departments: []Department{
				{
					Name: "Boaco",
					Municipalities: []Municipality{
						{
							Name: "Boaco",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Camoapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San José de los Remates",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Lorenzo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Santa Lucía",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Carazo",
					Municipalities: []Municipality{
						{
							Name: "Diriamba",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Dolores",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Rosario",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Jinotepe",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "La Conquista",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "La Paz de Oriente",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Marcos",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Chinandega",
					Municipalities: []Municipality{
						{
							Name: "Chichigalpa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Chinandega",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Cinco Pinos",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Corinto",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Realejo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Viejo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Posoltega",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Puerto Morazán",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Francisco del Norte",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Pedro del Norte",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Santo Tomás del Norte",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Somotillo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Chontales",
					Municipalities: []Municipality{
						{
							Name: "Acoyapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Comalapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Cuapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Coral",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Juigalpa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "La Libertad",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Pedro de Lóvago",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Santo Domingo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Santo Tomás",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Costa Caribe Norte",
					Municipalities: []Municipality{
						{
							Name: "Bonanza",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Mulukukú",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Prinzapolka",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Puerto Cabezas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Rosita",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Siuna",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Waslala",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Costa Caribe Sur",
					Municipalities: []Municipality{
						{
							Name: "Bluefields",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Corn Island",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Desembocadura de Río Grande",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Ayote",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Rama",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Tortuguero",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Kukra Hill",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "La Cruz de Río Grande",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Laguna de Perlas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Muelle de los Bueyes",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Nueva Guinea",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Estelí",
					Municipalities: []Municipality{
						{
							Name: "Condega",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Estelí",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "La Trinidad",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Pueblo Nuevo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Juan de Limay",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Granada",
					Municipalities: []Municipality{
						{
							Name: "Diriá",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Diriomo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Granada",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Jinotega",
					Municipalities: []Municipality{
						{
							Name: "El Cuá",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Jinotega",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "La Concordia",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San José de Bocay",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Rafael del Norte",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Sebastián de Yalí",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Santa María de Pantasma",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "León",
					Municipalities: []Municipality{
						{
							Name: "Achuapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Jicaral",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Sauce",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "La Paz Centro",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Larreynaga",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "León",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Nagarote",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Quezalguaque",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Santa Rosa del Peñón",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Madriz",
					Municipalities: []Municipality{
						{
							Name: "Las Sabanas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Palacagüina",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San José de Cusmapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Juan de Río Coco",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Lucas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Somoto",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Telpaneca",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Totogalpa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Managua",
					Municipalities: []Municipality{
						{
							Name: "Ciudad Sandino",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Crucero",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Managua",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Mateare",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Francisco Libre",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Rafael del Sur",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Ticuantepe",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Tipitapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Masaya",
					Municipalities: []Municipality{
						{
							Name: "Catarina",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "La Concepción",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Masatepe",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Masaya",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Nandasmo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Nindirí",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Niquinohomo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Juan de Oriente",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Matagalpa",
					Municipalities: []Municipality{
						{
							Name: "Ciudad Darío",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Tuma - La Dalia",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Esquipulas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Matagalpa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Matiguás",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Muy Muy",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Rancho Grande",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Río Blanco",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Dionisio",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Isidro",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Ramón",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Sébaco",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
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

					Name: "Nueva Segovia",
					Municipalities: []Municipality{
						{
							Name: "Ciudad Antigua",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Dipilto",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Jícaro",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Jalapa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Macuelizo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Mozonte",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Murra",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Ocotal",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Quilalí",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Fernando",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Santa María",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Wiwilí",
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

					Name: "Río San Juan",
					Municipalities: []Municipality{
						{
							Name: "El Almendro",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "El Castillo",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Morrito",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Carlos",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Juan del Norte",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Miguelito",
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

					Name: "Rivas",
					Municipalities: []Municipality{
						{
							Name: "Altagracia",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Belén",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Buenos Aires",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Cárdenas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Moyogalpa",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Potosí",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Rivas",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Jorge",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "San Juan del Sur",
							Time: TimeAt{
								CreatedAt: time.Date(2023, time.March, 27, 12, 18, 0, 0, time.UTC),
							},
						},
						{
							Name: "Tola",
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
