package gads

type FeedService struct {
	Auth
}

func NewFeedService(auth *Auth) *FeedService {
	return &FeedService{Auth: *auth}
}
