package authapp

import "time"

// Person request
type PersonRequest struct {
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Birth   time.Time `json:"birth"`
	Sex     string    `json:"sex"`
}

// Document request
type PersonalDocumentRequest struct {
	Number string `json:"number"`
	Front  string `json:"front"`
	Back   string `json:"back"`
}

// Contacts request
type PhoneRequest struct {
	Country string `json:"country"`
	Phone   string `json:"phone"`
}
type AddressRequest struct {
	Municipality string  `json:"municipality"`
	Address      string  `json:"address"`
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
}
type ContactsRequest struct {
	Emails    []string         `json:"emails"`
	Phones    []PhoneRequest   `json:"phones"`
	Addresses []AddressRequest `json:"addresses"`
}

type SingleContactsRequest struct {
	Email   string         `json:"email"`
	Phone   PhoneRequest   `json:"phone"`
	Address AddressRequest `json:"address"`
}

// User request
type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
