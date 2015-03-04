package gads

import (
//	"encoding/xml"
//	"fmt"
)

type AdGroupFeedService struct {
	Auth
}

func NewAdGroupFeedService(auth *Auth) *AdGroupFeedService {
	return &AdGroupFeedService{Auth: *auth}
}

type AdGroupFeedOperations struct {
}

type AdGroupFeed struct {
}

// Get is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupFeedService#get
//
func (s AdGroupFeedService) Get(selector Selector) (adGroupFeeds []AdGroupFeed, err error) {
	return adGroupFeeds, ERROR_NOT_YET_IMPLEMENTED
}

// Mutate is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupFeedService#mutate
//
func (s *AdGroupFeedService) Mutate(adGroupFeedOperations AdGroupFeedOperations) (adGroupFeeds []AdGroupFeed, err error) {
	return adGroupFeeds, ERROR_NOT_YET_IMPLEMENTED
}

// Query is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupFeedService#query
//
func (s *AdGroupFeedService) Query(query string) (adGroupFeeds []AdGroupFeed, err error) {
	return adGroupFeeds, ERROR_NOT_YET_IMPLEMENTED
}
