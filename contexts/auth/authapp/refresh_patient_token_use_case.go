package authapp

import (
	"errors"
	"time"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
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
	// Current Token
	currentIssuer, currentSubject, currentExpiration, currentErr := authdomain.DecodeToken(request.Token)

	if currentErr != nil {
		return RefreshPatientTokenResponse{}, errors.New(string(shareddomain.ERRORS_UNAUTHORIZED))
	}

	if currentIssuer != "OnLab-Clinical" {
		return RefreshPatientTokenResponse{}, errors.New(string(authdomain.ERRORS_TOKEN_UNKNOWN))
	}

	if currentExpiration.UTC().After(time.Now().UTC()) {
		return RefreshPatientTokenResponse{}, errors.New(string(authdomain.ERRORS_TOKEN_ALREADY_WORKING))
	}

	// Current Refresh Token
	refreshIssuer, refreshSubject, refreshExpiration, refreshErr := authdomain.DecodeToken(request.RefreshToken)

	if refreshErr != nil {
		return RefreshPatientTokenResponse{}, errors.New(string(shareddomain.ERRORS_UNAUTHORIZED))
	}

	if refreshIssuer != "OnLab-Clinical" {
		return RefreshPatientTokenResponse{}, errors.New(string(authdomain.ERRORS_REFRESH_TOKEN_UNKNOWN))
	}

	if refreshExpiration.UTC().Before(time.Now().UTC()) {
		return RefreshPatientTokenResponse{}, errors.New(string(authdomain.ERRORS_REFRESH_TOKEN_EXPIRED))
	}

	// Compare Subject
	if currentSubject != refreshSubject {
		return RefreshPatientTokenResponse{}, errors.New(string(authdomain.ERRORS_TOKEN_SUBJECT_MISMATCH))
	}

	// Get Current Patient
	patient, patientErr := uc.PatientRepository.ReadPatientById(currentSubject)

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
