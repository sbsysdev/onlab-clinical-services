package authdomain

type LocationRepository interface {
	GetMunicipalityById(uint16) (Municipality, error)

	GetDepartmentById(uint16, bool) (Department, error)

	GetCountryById(uint8, bool, bool) (Country, error)
	GetCountryList(bool, bool) ([]Country, error)
}
