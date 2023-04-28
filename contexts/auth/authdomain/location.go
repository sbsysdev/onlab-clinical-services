package authdomain

// Municipality Value Object
type Municipality struct {
	Id   uint16 `json:"id"`
	Name string `json:"name"`
}

// Department Value Object
type Department struct {
	Id             uint16         `json:"id"`
	Name           string         `json:"name"`
	Municipalities []Municipality `json:"municipality"`
}

// Country Value Object
type Country struct {
	Id          uint8        `json:"id"`
	Name        string       `json:"name"`
	CallingCode string       `json:"callingCode"`
	Departments []Department `json:"department"`
}
