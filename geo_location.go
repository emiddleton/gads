package gads

type GeoLocationService struct {
	Auth
}

func NewGeoLocationService(auth *Auth) *GeoLocationService {
	return &GeoLocationService{Auth: *auth}
}
