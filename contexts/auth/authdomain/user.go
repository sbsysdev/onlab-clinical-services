package authdomain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
)

// User Name Value Object
type UserName string

const (
	ERRORS_USER_NAME_EMPTY         shareddomain.DomainError = "ERRORS_USER_NAME_EMPTY"
	ERRORS_USER_NAME_NOT_AVAILABLE shareddomain.DomainError = "ERRORS_USER_NAME_NOT_AVAILABLE"
)

func CreateUserName(name string) (UserName, error) {
	if len(name) == 0 {
		return UserName(""), errors.New(string(ERRORS_USER_NAME_EMPTY))
	}

	return UserName(name), nil
}

// User Password Value Object
type UserPassword string

const (
	ERRORS_USER_PASSWORD_EMPTY shareddomain.DomainError = "ERRORS_USER_PASSWORD_EMPTY"
)

func CreateUserPassword(password string) (UserPassword, error) {
	if len(password) == 0 {
		return UserPassword(""), errors.New(string(ERRORS_USER_PASSWORD_EMPTY))
	}

	// TODO: Validate min security format

	hashed, err := bcrypt.GenerateFromPassword([]byte(string(password)), 16)

	if err != nil {
		return UserPassword(""), err
	}

	return UserPassword(hashed), nil
}

// User State Value Object
type UserState string

const (
	STATES_USER_STATE_UNVERIFIED UserState = "unverified"
	STATES_USER_STATE_BLOCKED    UserState = "blocked"
	STATES_USER_STATE_VERIFIED   UserState = "verified"
	STATES_USER_STATE_SUSPENDED  UserState = "suspended"
)

const (
	ERRORS_USER_STATE_NOT_VALID shareddomain.DomainError = "ERRORS_USER_STATE_NOT_VALID"
)

func CreateUserState(state string) (UserState, error) {
	if state != string(STATES_USER_STATE_UNVERIFIED) && state != string(STATES_USER_STATE_BLOCKED) && state != string(STATES_USER_STATE_VERIFIED) && state != string(STATES_USER_STATE_SUSPENDED) {
		return UserState(""), errors.New(string(ERRORS_USER_STATE_NOT_VALID))
	}

	return UserState(state), nil
}

// User Value Object
type User struct {
	Name     UserName
	Password UserPassword
	State    UserState
}

// User Value Object Factory
func CreateUser(name UserName, password UserPassword) User {
	return User{
		Name:     name,
		Password: password,
		State:    STATES_USER_STATE_UNVERIFIED,
	}
}
