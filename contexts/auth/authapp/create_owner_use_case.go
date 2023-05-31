package authapp

// Request
type CreateOwnerRequest struct{}

// Use Case
type CreateOwnerUseCase struct{}

func (uc CreateOwnerUseCase) Command(request CreateOwnerRequest) error {
	return nil
}
