package shareddomain

type DomainError string

const (
	ERRORS_NOT_FOUND     DomainError = "ERRORS_NOT_FOUND"
	ERRORS_UNIMPLEMENTED DomainError = "ERRORS_UNIMPLEMENTED"
)
