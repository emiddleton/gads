package gads

type CampaignSharedSetService struct {
	Auth
}

func NewCampaignSharedSetService(auth *Auth) *CampaignSharedSetService {
	return &CampaignSharedSetService{Auth: *auth}
}
