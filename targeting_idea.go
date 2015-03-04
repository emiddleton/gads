package gads

type TargetIdeaService struct {
	Auth
}

func NewTargetIdeaService(auth *Auth) *TargetIdeaService {
	return &TargetIdeaService{Auth: *auth}
}
