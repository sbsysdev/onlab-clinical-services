package authdomain

import (
	"errors"
	"mime/multipart"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
)

type NationalIdentityDocumentNumber string

type NationalIdentityDocument struct {
	Number       NationalIdentityDocumentNumber
	FrontPicture *multipart.FileHeader
	BackPicture  *multipart.FileHeader
}

const (
	ERRORS_NID_NUMBER_EMPTY shareddomain.DomainError = "ERRORS_NID_NUMBER_EMPTY"
	ERRORS_NID_FRONT_EMPTY  shareddomain.DomainError = "ERRORS_NID_FRONT_EMPTY"
	ERRORS_NID_BACK_EMPTY   shareddomain.DomainError = "ERRORS_NID_BACK_EMPTY"
)

func CreateNID(number string, front, back *multipart.FileHeader) (NationalIdentityDocument, error) {
	if len(number) == 0 {
		return NationalIdentityDocument{}, errors.New(string(ERRORS_NID_NUMBER_EMPTY))
	}

	if front.Size == 0 || len(front.Filename) == 0 {
		return NationalIdentityDocument{}, errors.New(string(ERRORS_NID_FRONT_EMPTY))
	}

	if back.Size == 0 || len(back.Filename) == 0 {
		return NationalIdentityDocument{}, errors.New(string(ERRORS_NID_BACK_EMPTY))
	}

	return NationalIdentityDocument{
		Number:       NationalIdentityDocumentNumber(number),
		FrontPicture: front,
		BackPicture:  back,
	}, nil
}

func CreateNIDNumber(number string) (NationalIdentityDocumentNumber, error) {
	if len(number) == 0 {
		return "", errors.New(string(ERRORS_NID_NUMBER_EMPTY))
	}

	return NationalIdentityDocumentNumber(number), nil
}
