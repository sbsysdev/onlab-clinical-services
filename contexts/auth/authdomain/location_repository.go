package authdomain

type LocationRepository interface {
	IsExistingCountry(uint8) bool
	IsExistingMunicipality(uint16) bool
}
