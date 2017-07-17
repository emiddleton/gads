package gads

import "encoding/xml"

type FeedService struct {
	Auth
}

func NewFeedService(auth *Auth) *FeedService {
	return &FeedService{Auth: *auth}
}

// https://developers.google.com/adwords/api/docs/reference/v201609/FeedService
func (s *FeedService) Query(query string) (page []Feed, totalCount int64, err error) {
	respBody, err := s.Auth.request(
		feedServiceUrl,
		"query",
		AWQLQuery{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "query",
			},
			Query: query,
		},
		nil,
	)
	if err != nil {
		return
	}

	getResp := struct {
		Size  int64  `xml:"rval>totalNumEntries"`
		Feeds []Feed `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return
	}
	return getResp.Feeds, getResp.Size, err
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.Feed.Type
// Feed hard type. Values coincide with placeholder type id.
// Enum: NONE, SITELINK, CALL, APP, REVIEW, AD_CUSTOMIZER, CALLOUT, STRUCTURED_SNIPPET, PRICE
type FeedType string

// https://developers.google.com/adwords/api/docs/reference/v201609/FeedService.Feed.Status
// Status of the Feed.
// ENABLED, REMOVED, UNKNOWN
type FeedStatus string

// Used to Specify who manages the FeedAttributes for the Feed.
// https://developers.google.com/adwords/api/docs/reference/v201609/FeedService.Feed.Origin
// USER, ADWORDS, UNKNOWN
type FeedOrigin string

// https://developers.google.com/adwords/api/docs/reference/v201609/FeedService.FeedAttribute.Type
// Possible data types.
type FeedAttributeType string

// Configuration data allowing feed items to be populated for a system feed.
// https://developers.google.com/adwords/api/docs/reference/v201609/FeedService.SystemFeedGenerationData
type SystemFeedGenerationData struct {
	SystemFeedGenerationDataType string `xml:"https://adwords.google.com/api/adwords/cm/v201609 SystemFeedGenerationData.Type,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201609/FeedService.FeedAttribute
// FeedAttributes define the types of data expected to be present in a Feed.
// A single FeedAttribute specifies the expected type of the FeedItemAttributes with the same FeedAttributeId.
// Optionally, a FeedAttribute can be marked as being part of a FeedItem's unique key.
type FeedAttribute struct {
	Id          int64             `xml:"https://adwords.google.com/api/adwords/cm/v201609 id,omitempty"`
	Name        string            `xml:"https://adwords.google.com/api/adwords/cm/v201609 name,omitempty"`
	Type        FeedAttributeType `xml:"https://adwords.google.com/api/adwords/cm/v201609 type,omitempty"`
	IsPartOfKey bool              `xml:"https://adwords.google.com/api/adwords/cm/v201609 isPartOfKey,omitempty"`
}

// A Feed identifies a source of data and its schema.
// The data for the Feed can either be user-entered via the FeedItemService or system-generated, in which case the data is provided automatically.
// https://developers.google.com/adwords/api/docs/reference/v201609/FeedService.Feed
type Feed struct {
	Id                       int64                      `xml:"https://adwords.google.com/api/adwords/cm/v201609 id,omitempty"`
	Name                     string                     `xml:"https://adwords.google.com/api/adwords/cm/v201609 name,omitempty"`
	Attributes               []FeedAttribute            `xml:"https://adwords.google.com/api/adwords/cm/v201609 attributes,omitempty"`
	Status                   FeedStatus                 `xml:"https://adwords.google.com/api/adwords/cm/v201609 status,omitempty"`
	Origin                   FeedOrigin                 `xml:"https://adwords.google.com/api/adwords/cm/v201609 origin,omitempty"`
	SystemFeedGenerationData []SystemFeedGenerationData `xml:"https://adwords.google.com/api/adwords/cm/v201609 systemFeedGenerationData,omitempty"`
}
