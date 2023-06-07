package authctrls

type PersonResponse struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	NID     string `json:"nid"`
	Birth   string `json:"birth"`
	Sex     string `json:"sex"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

// Location Response
type CountryResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Calling string `json:"calling"`
}
type MunicipalityResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PhoneResponse struct {
	Country CountryResponse `json:"country"`
	Phone   string          `json:"phone"`
}

type AddressResponse struct {
	Municipality MunicipalityResponse `json:"municipality"`
	Address      string               `json:"address"`
}

type SingleContactsResponse struct {
	Email   string          `json:"email"`
	Phone   PhoneResponse   `json:"phone"`
	Address AddressResponse `json:"address"`
}

type AliasRoleResponse struct {
	ID    string            `json:"id"`
	Alias string            `json:"alias"`
	Name  map[string]string `json:"name"`
}
