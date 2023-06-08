package authapp

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/utils"
)

type RefreshPatientTokenRequest struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh"`
}

type RefreshPatientTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh"`
}

type RefreshPatientTokenUseCase struct {
	// Repositories
	PatientRepository authdomain.PatientRepository
}

func (uc RefreshPatientTokenUseCase) Query(request RefreshPatientTokenRequest) (RefreshPatientTokenResponse, error) {
	currentToken := request.Token
	currentRefreshToken := request.RefreshToken
	jwtKey := utils.GetEnv("JWT_KEY", "qwerty")

	// Current Token
	token, _ := jwt.Parse(currentToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(string(shareddomain.ERRORS_UNAUTHORIZED))
		}

		return []byte(jwtKey), nil
	})

	claims, claimsOk := token.Claims.(jwt.MapClaims)

	if !claimsOk {
		return RefreshPatientTokenResponse{}, errors.New(string(shareddomain.ERRORS_UNAUTHORIZED))
	}

	tokenExp, _ := claims.GetExpirationTime()

	if tokenExp.Time.UTC().After(time.Now().UTC()) {
		return RefreshPatientTokenResponse{}, errors.New(string(authdomain.ERRORS_TOKEN_ALREADY_WORKING))
	}

	// Current Refresh Token
	refreshToken, _ := jwt.Parse(currentRefreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(string(shareddomain.ERRORS_UNAUTHORIZED))
		}

		return []byte(jwtKey), nil
	})

	refreshClaims, refreshClaimsOk := refreshToken.Claims.(jwt.MapClaims)

	if !refreshClaimsOk {
		return RefreshPatientTokenResponse{}, errors.New(string(shareddomain.ERRORS_UNAUTHORIZED))
	}

	refreshTokenExp, _ := refreshClaims.GetExpirationTime()

	if refreshTokenExp.Time.UTC().Before(time.Now().UTC()) {
		return RefreshPatientTokenResponse{}, errors.New(string(authdomain.ERRORS_REFRESH_TOKEN_EXPIRED))
	}

	// Compare Subject

	tokenSub, tsErr := claims.GetSubject()

	if tsErr != nil {
		return RefreshPatientTokenResponse{}, tsErr
	}

	refreshTokenSub, rtsErr := refreshClaims.GetSubject()

	if rtsErr != nil {
		return RefreshPatientTokenResponse{}, rtsErr
	}

	if tokenSub != refreshTokenSub {
		return RefreshPatientTokenResponse{}, errors.New(string(authdomain.ERRORS_TOKEN_SUBJECT_MISMATCH))
	}

	// Get Current Patient
	patient, patientErr := uc.PatientRepository.ReadPatientById(tokenSub)

	if patientErr != nil {
		return RefreshPatientTokenResponse{}, patientErr
	}

	if patient.User.State != authdomain.USER_STATE_VERIFIED {
		return RefreshPatientTokenResponse{}, errors.New(string(authdomain.ERRORS_USER_STATE_UNVERIFIED))
	}

	// Get Token & Refresh Token
	signed, signedRefresh, signedErr := authdomain.CreatePatientTokenAndRefreshToken(patient.ID)

	if signedErr != nil {
		return RefreshPatientTokenResponse{}, signedErr
	}

	return RefreshPatientTokenResponse{
		Token:        signed,
		RefreshToken: signedRefresh,
	}, nil
}
