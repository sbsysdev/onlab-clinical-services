package authdomain

type LocationRepository interface {
	GetMunicipalityById(string) (Municipality, error)

	GetDepartmentById(string) (Department, error)

	GetCountryById(string) (Country, error)
	GetCountryList() ([]Country, error)
}
