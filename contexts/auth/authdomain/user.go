package authdomain

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/argon2"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
)

// User Name Value Object
type UserName string

const (
	ERRORS_USER_NAME_EMPTY         shareddomain.DomainError = "ERRORS_USER_NAME_EMPTY"
	ERRORS_USER_NAME_NOT_AVAILABLE shareddomain.DomainError = "ERRORS_USER_NAME_NOT_AVAILABLE"
	ERRORS_USER_NAME_NOT_FOUND     shareddomain.DomainError = "ERRORS_USER_NAME_NOT_FOUND"
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
	ERRORS_USER_PASSWORD_EMPTY    shareddomain.DomainError = "ERRORS_USER_PASSWORD_EMPTY"
	ERRORS_USER_PASSWORD_FORMAT   shareddomain.DomainError = "ERRORS_USER_PASSWORD_FORMAT"
	ERRORS_USER_PASSWORD_MISMATCH shareddomain.DomainError = "ERRORS_USER_PASSWORD_MISMATCH"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var p = &params{
	memory:      128 * 1024,
	iterations:  4,
	parallelism: 4,
	saltLength:  16,
	keyLength:   32,
}

func CreateUserPassword(password string) (UserPassword, error) {
	if len(password) == 0 {
		return UserPassword(""), errors.New(string(ERRORS_USER_PASSWORD_EMPTY))
	}

	// Validate min security format
	if matched, err := regexp.MatchString("^.*[a-z]+.*$", password); !matched || err != nil {
		return UserPassword(""), errors.New(string(ERRORS_USER_PASSWORD_FORMAT))
	}
	if matched, err := regexp.MatchString("^.*[A-Z]+.*$", password); !matched || err != nil {
		return UserPassword(""), errors.New(string(ERRORS_USER_PASSWORD_FORMAT))
	}
	if matched, err := regexp.MatchString("^.*\\d+.*$", password); !matched || err != nil {
		return UserPassword(""), errors.New(string(ERRORS_USER_PASSWORD_FORMAT))
	}
	if matched, err := regexp.MatchString("^.*\\W+.*$", password); !matched || err != nil {
		return UserPassword(""), errors.New(string(ERRORS_USER_PASSWORD_FORMAT))
	}
	if matched, err := regexp.MatchString("^.{8,32}$", password); !matched || err != nil {
		return UserPassword(""), errors.New(string(ERRORS_USER_PASSWORD_FORMAT))
	}

	salt := make([]byte, p.saltLength)
	if _, err := rand.Read(salt); err != nil {
		return UserPassword(""), err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return UserPassword(encodedHash), nil
}

const (
	ERRORS_INVALID_HASH         = "ERRORS_INVALID_HASH"
	ERRORS_INCOMPATIBLE_VERSION = "ERRORS_INCOMPATIBLE_VERSION"
)

func decodeHash(encodedHash string) (p *params, salt []byte, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.New(ERRORS_INVALID_HASH)
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New(ERRORS_INCOMPATIBLE_VERSION)
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
func ComparePasswordAndHash(password, encodedHash string) error {
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return err
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return nil
	}
	return errors.New(string(ERRORS_USER_PASSWORD_MISMATCH))
}

// User State Value Object
type UserState string

const (
	USER_STATE_UNVERIFIED UserState = "unverified"
	USER_STATE_BLOCKED    UserState = "blocked"
	USER_STATE_VERIFIED   UserState = "verified"
	USER_STATE_SUSPENDED  UserState = "suspended"
)

const (
	ERRORS_USER_STATE_NOT_VALID shareddomain.DomainError = "ERRORS_USER_STATE_NOT_VALID"
	ERRORS_USER_STATE_SUSPENDED shareddomain.DomainError = "ERRORS_USER_STATE_SUSPENDED"
	ERRORS_USER_STATE_BLOCKED   shareddomain.DomainError = "ERRORS_USER_STATE_BLOCKED"
)

func CreateUserState(state string) (UserState, error) {
	if state != string(USER_STATE_UNVERIFIED) && state != string(USER_STATE_BLOCKED) && state != string(USER_STATE_VERIFIED) && state != string(USER_STATE_SUSPENDED) {
		return UserState(""), errors.New(string(ERRORS_USER_STATE_NOT_VALID))
	}

	return UserState(state), nil
}

// User Value Object
type User struct {
	Name     UserName     `json:"name"`
	Password UserPassword `json:"password"`
	State    UserState    `json:"state"`
}

// User Value Object Factory
func CreateUser(name UserName, password UserPassword) User {
	return User{
		Name:     name,
		Password: password,
		State:    USER_STATE_VERIFIED,
	}
}
