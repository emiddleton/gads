package gads

type ManagedCustomerService struct {
	Auth
}

func NewManagedCustomerService(auth *Auth) *ManagedCustomerService {
	return &ManagedCustomerService{Auth: *auth}
}
