package authdomain

// Municipality Value Object
type Municipality struct {
	Id   uint16
	Name string
}

// Department Value Object
type Department struct {
	Id             uint16
	Name           string
	Municipalities []Municipality
}

// Country Value Object
type Country struct {
	Id          uint8
	Name        string
	CallingCode string
	Departments []Department
}
