package gads

type CustomerSyncService struct {
	Auth
}

func NewCustomerSyncService(auth *Auth) *CustomerSyncService {
	return &CustomerSyncService{Auth: *auth}
}
