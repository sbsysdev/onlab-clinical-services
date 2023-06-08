package authapp

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/utils"
)

// Request
type SignInPatientRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Response
type SignInPatientResponse struct {
	Token        string                   `json:"token"`
	RefreshToken string                   `json:"refresh"`
	Patient      authdomain.PatientEntity `json:"patient"`
}

// Use Case
type SignInPatientUseCase struct {
	// Repositories
	PatientRepository authdomain.PatientRepository
}

func (uc SignInPatientUseCase) Query(request SignInPatientRequest) (SignInPatientResponse, error) {
	patient, patientErr := uc.PatientRepository.ReadPatientByName(request.Name)

	if patientErr != nil {
		return SignInPatientResponse{}, patientErr
	}

	if err := authdomain.ComparePasswordAndHash(request.Password, string(patient.User.Password)); err != nil {
		return SignInPatientResponse{}, err
	}

	if patient.User.State == authdomain.USER_STATE_SUSPENDED {
		return SignInPatientResponse{}, errors.New(string(authdomain.ERRORS_USER_STATE_SUSPENDED))
	}

	if patient.User.State == authdomain.USER_STATE_BLOCKED {
		return SignInPatientResponse{}, errors.New(string(authdomain.ERRORS_USER_STATE_BLOCKED))
	}

	patient.User.Password = ""

	jwtKey := utils.GetEnv("JWT_KEY", "qwerty")

	// Token
	exp := time.Now().UTC().Add(time.Hour * 2)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "OnLab-Clinical",
		"sub": patient.ID,
		"exp": exp,
	})
	signed, signedErr := token.SignedString([]byte(jwtKey))

	if signedErr != nil {
		return SignInPatientResponse{}, signedErr
	}

	// Refresh Token
	expRefresh := exp.Add(time.Minute * 5)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "OnLab-Clinical",
		"sub": patient.ID,
		"exp": expRefresh,
	})
	signedRefresh, signedRefreshErr := refreshToken.SignedString([]byte(jwtKey))

	if signedRefreshErr != nil {
		return SignInPatientResponse{}, signedErr
	}

	return SignInPatientResponse{
		Token:        signed,
		RefreshToken: signedRefresh,
		Patient:      patient,
	}, nil
}
