package gads

type MutateJobService struct {
	Auth
}

func NewMutateJobService(auth *Auth) *MutateJobService {
	return &MutateJobService{Auth: *auth}
}
