package gads

type CustomerService struct {
	Auth
}

func NewCustomerService(auth *Auth) *CustomerService {
	return &CustomerService{Auth: *auth}
}
