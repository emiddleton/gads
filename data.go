package gads

import (
	"encoding/xml"
)

type DataService struct {
	Auth
}

func NewDataService(auth *Auth) *DataService {
	return &DataService{Auth: *auth}
}

type LandscapePoint struct {
	Bid                 *int64 `xml:"bid>microAmount,omitempty"`
	Clicks              *int64 `xml:"clicks,omitempty"`
	Cost                *int64 `xml:"cost>microAmount,omitempty"`
	Impressions         *int64 `xml:"impressions,omitempty"`
	PromotedImpressions *int64 `xml:"promotedImpressions,omitempty"`
}

type BidLandscape struct {
	CampaignId      *int64           `xml:"campaignId,omitempty"`
	AdGroupId       *int64           `xml:"adGroupId,omitempty"`
	StartDate       *string          `xml:"startDate,omitempty"`
	EndDate         *string          `xml:"endDate,omitempty"`
	LandscapePoints []LandscapePoint `xml:"landscapePoints"`
}

type AdGroupBidLandscape struct {
	BidLandscape
	Type             string  `xml:"type"`
	LandscapeCurrent bool    `xml:"landscapeCurrent"`
	Errors           []error `xml:"-"`
}

type CriterionBidLandscape struct {
	BidLandscape
	CriterionId *int64 `xml:"criterionId,omitempty"`
}

// GetAdGroupBidLandscape returns an array of AdGroupBidLandscape objects
// the selector.
//
// Example
//
//   adGroupBidLandscape, totalCount, err := dataService.GetAdGroupBidLandscape(
//     gads.Selector{
//       Fields: []string{
//         "AdGroupId",
//         "Bid",
//         "CampaignId",
//         "LocalClicks",
//         "LocalCost",
//		   "LocalImpressions",
//       },
//       Predicates: []gads.Predicate{
//         {"AdGroupId", "EQUALS", []string{adGroupId}},
//       },
//     },
//   )
//
// Selectable fields are
//   "AdGroupId", "Bid", "CampaignId", "EndDate", "LandscapeCurrent", "LandscapeType", "LocalClicks",
//   "LocalCost", "LocalImpressions", "PromotedImpressions", "StartDate"
//
// filterable fields are
//   "AdGroupId", "Bid", "CampaignId", "LandscapeCurrent", "LandscapeType", "LocalClicks",
//   "LocalCost", "LocalImpressions", "PromotedImpressions"
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201609/DataService#getadgroupbidlandscape
//	   https://developers.google.com/adwords/api/docs/appendix/selectorfields#v201609-DataService
//
func (s *DataService) GetAdGroupBidLandscape(selector Selector) (adGroupBidLandscapes []AdGroupBidLandscape, totalCount int64, err error) {
	// The default namespace, "", will break in 1.5 with the addition of
	// custom namespace support.  Hence, we have to ensure that the baseUrl is
	// set again as the proper namespace for the service/serviceSelector element
	selector.XMLName = xml.Name{baseUrl, "serviceSelector"}

	respBody, err := s.Auth.request(
		dataServiceUrl,
		"getAdGroupBidLandscape",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "getAdGroupBidLandscape",
			},
			Sel: selector,
		},
		nil,
	)
	if err != nil {
		return adGroupBidLandscapes, totalCount, err
	}
	getResp := struct {
		Size                 int64                 `xml:"rval>totalNumEntries"`
		AdGroupBidLandscapes []AdGroupBidLandscape `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroupBidLandscapes, totalCount, err
	}
	return getResp.AdGroupBidLandscapes, getResp.Size, err
}

// GetCriterionBidLandscape returns an array of CriterionBidLandscape objects
// the selector.
//
// Example
//
//   criterionBidLandscape, totalCount, err := dataService.GetCriterionBidLandscape(
//     gads.Selector{
//       Fields: []string{
//         "AdGroupId",
//		   "CriterionId",
//         "Bid",
//         "CampaignId",
//         "LocalClicks",
//         "LocalCost",
//		   "LocalImpressions",
//       },
//       Predicates: []gads.Predicate{
//         {"CriterionId", "EQUALS", []string{criterionId}},
//       },
//     },
//   )
//
// Selectable fields are
//   "AdGroupId", "Bid", "CampaignId", "CriterionId", "EndDate", "LandscapeCurrent", "LandscapeType", "LocalClicks",
//   "LocalCost", "LocalImpressions", "PromotedImpressions", "StartDate"
//
// filterable fields are
//   "AdGroupId", "Bid", "CampaignId", "CriterionId", "LandscapeCurrent", "LandscapeType", "LocalClicks",
//   "LocalCost", "LocalImpressions", "PromotedImpressions"
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201609/DataService#getcriterionbidlandscape
//	   https://developers.google.com/adwords/api/docs/appendix/selectorfields#v201609-DataService
//
func (s *DataService) GetCriterionBidLandscape(selector Selector) (criterionBidLandscapes []CriterionBidLandscape, totalCount int64, err error) {
	// The default namespace, "", will break in 1.5 with the addition of
	// custom namespace support.  Hence, we have to ensure that the baseUrl is
	// set again as the proper namespace for the service/serviceSelector element
	selector.XMLName = xml.Name{baseUrl, "serviceSelector"}

	respBody, err := s.Auth.request(
		dataServiceUrl,
		"getCriterionBidLandscape",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "getCriterionBidLandscape",
			},
			Sel: selector,
		},
		nil,
	)
	if err != nil {
		return criterionBidLandscapes, totalCount, err
	}
	getResp := struct {
		Size                   int64                   `xml:"rval>totalNumEntries"`
		CriterionBidLandscapes []CriterionBidLandscape `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return criterionBidLandscapes, totalCount, err
	}
	return getResp.CriterionBidLandscapes, getResp.Size, err
}

// QueryAdGroupBidLandscape returns a slice of AdGroupBidLandscapes based on passed in AWQL query
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201609/DataService#queryadgroupbidlandscape
//
func (s *DataService) QueryAdGroupBidLandscape(query string) (adGroupBidLandscapes []AdGroupBidLandscape, totalCount int64, err error) {

	respBody, err := s.Auth.request(
		dataServiceUrl,
		"queryAdGroupBidLandscape",
		AWQLQuery{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "queryAdGroupBidLandscape",
			},
			Query: query,
		},
		nil,
	)

	if err != nil {
		return adGroupBidLandscapes, totalCount, err
	}

	getResp := struct {
		Size                 int64                 `xml:"rval>totalNumEntries"`
		AdGroupBidLandscapes []AdGroupBidLandscape `xml:"rval>entries"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroupBidLandscapes, totalCount, err
	}
	return getResp.AdGroupBidLandscapes, getResp.Size, err
}

// QueryCriterionBidLandscape returns a slice of CriterionBidLandscapes based on passed in AWQL query
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201609/DataService#querycriterionbidlandscape
//
func (s *DataService) QueryCriterionBidLandscape(query string) (criterionBidLandscapes []CriterionBidLandscape, totalCount int64, err error) {

	respBody, err := s.Auth.request(
		dataServiceUrl,
		"queryCriterionBidLandscape",
		AWQLQuery{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "queryCriterionBidLandscape",
			},
			Query: query,
		},
		nil,
	)

	if err != nil {
		return criterionBidLandscapes, totalCount, err
	}

	getResp := struct {
		Size                   int64                   `xml:"rval>totalNumEntries"`
		CriterionBidLandscapes []CriterionBidLandscape `xml:"rval>entries"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return criterionBidLandscapes, totalCount, err
	}
	return getResp.CriterionBidLandscapes, getResp.Size, err
}
