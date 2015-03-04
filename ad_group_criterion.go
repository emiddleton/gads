package gads

import (
	"encoding/xml"
	"fmt"
)

type AdGroupCriterionService struct {
	Auth
}

func NewAdGroupCriterionService(auth *Auth) *AdGroupCriterionService {
	return &AdGroupCriterionService{Auth: *auth}
}

type QualityInfo struct {
	IsKeywordAdRelevanceAcceptable bool  `xml:"isKeywordAdRelevanceAcceptable,omitempty"`
	IsLandingPageQualityAcceptable bool  `xml:"isLandingPageQualityAcceptable,omitempty"`
	IsLandingPageLatencyAcceptable bool  `xml:"isLandingPageLatencyAcceptable,omitempty"`
	QualityScore                   int64 `xml:"QualityScore,omitempty"`
}

type Cpc struct {
	Amount int64 `xml:"amount"`
}

type AdGroupCriterions []interface{}

type AdGroupCriterionLabel struct {
	AdGroupCriterionId int64 `xml:"adGroupCriterionId"`
	LabelId            int64 `xml:"labelId"`
}

type AdGroupCriterionLabelOperations map[string][]AdGroupCriterionLabel

func (agcs *AdGroupCriterions) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	adGroupCriterionType, err := findAttr(start.Attr, xml.Name{
		Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"})
	if err != nil {
		return err
	}
	switch adGroupCriterionType {
	case "BiddableAdGroupCriterion":
		bagc := BiddableAdGroupCriterion{}
		err := dec.DecodeElement(&bagc, &start)
		if err != nil {
			return err
		}
		*agcs = append(*agcs, bagc)
	case "NegativeAdGroupCriterion":
		nagc := NegativeAdGroupCriterion{}
		err := dec.DecodeElement(&nagc, &start)
		if err != nil {
			return err
		}
		*agcs = append(*agcs, nagc)
	default:
		return fmt.Errorf("unknown AdGroupCriterion -> %#v", adGroupCriterionType)
	}
	return nil
}

type AdGroupCriterionOperations map[string]AdGroupCriterions

// Get returns an array of AdGroupCriterion's and the total number of AdGroupCriterion's matching
// the selector.
//
// Example
//
//   adGroupCriterions, totalCount, err := adGroupCriterionService.Get(
//     Selector{
//       Fields: []string{"Id","KeywordText","KeywordMatchType"},
//       Predicates: []Predicate{
//         {"AdGroupId", "EQUALS", []string{"432434"}},
//       },
//     },
//   )
//
// Selectable fields are
//   "AdGroupId", "CriterionUse", "Id", "CriteriaType", "Labels"
//
//   AgeRange
//     "AgeRangeType",
//
//   AppPaymentModel
//     "AppPaymentModelType",
//
//   CriterionUserInterest
//     "UserInterestId", "UserInterestName",
//
//   CriterionUserList
//     "UserListId", "UserListName", "UserListMembershipStatus"
//
//   Gender
//     "GenderType"
//
//   Keyword
//     "KeywordText", "KeywordMatchType"
//
//   MobileAppCategory
//     "MobileAppCategoryId"
//
//   MobileApplication
//     "DisplayName"
//
//   Placement
//     "PlacementUrl"
//
//   Product
//     "Text"
//
//   ProductPartition
//     "PartitionType", "ParentCriterionId", "CaseValue"
//
//   Vertical
//     "VerticalId", "VerticalParentId", "Path"
//
//   Webpage
//     "Parameter", "CriteriaCoverage", "CriteriaSamples"
//
// filterable fields are
//
//   "AdGroupId", "CriterionUse", "Id", "CriteriaType", "Labels"
//
//   CriterionUserList
//     "UserListMembershipStatus"
//
//   Keyword
//     "KeywordText", "KeywordMatchType"
//
//   MobileApplication
//     "DisplayName"
//
//   Placement
//     "PlacementUrl"
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupCriterionService#get
//
func (s AdGroupCriterionService) Get(selector Selector) (adGroupCriterions AdGroupCriterions, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		adGroupCriterionServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "get",
			},
			Sel: selector,
		},
	)
	if err != nil {
		return adGroupCriterions, totalCount, err
	}
	getResp := struct {
		Size              int64             `xml:"rval>totalNumEntries"`
		AdGroupCriterions AdGroupCriterions `xml:"rval>entries"`
	}{}
	fmt.Printf("%s\n", respBody)
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroupCriterions, totalCount, err
	}
	return getResp.AdGroupCriterions, getResp.Size, err
}

// Mutate allows you to add, modify and remove ad group criterion, returning the
// modified ad group criterion.
//
// Example
//
//  ads, err := adGroupService.Mutate(
//    gads.AdGroupCriterionOperations{
//      "ADD": {
//        BiddableAdGroupCriterion{
//          AdGroupId:  adGroupId,
//          Criterion:  gads.KeywordCriterion{Text: "test1", MatchType: "EXACT"},
//          UserStatus: "PAUSED",
//        },
//        NegativeAdGroupCriterion{
//          AdGroupId: adGroupId,
//          Criterion: gads.KeywordCriterion{Text: "test4", MatchType: "BROAD"},
//        },
//      },
//      "SET": {
//        modifiedAdGroupCriterion,
//      },
//      "REMOVE": {
//        adGroupCriterionNeedingRemoval,
//      },
//    },
//  )
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupCriterionService#mutate
//
func (s *AdGroupCriterionService) Mutate(adGroupCriterionOperations AdGroupCriterionOperations) (adGroupCriterions AdGroupCriterions, err error) {
	type adGroupCriterionOperation struct {
		Action           string      `xml:"operator"`
		AdGroupCriterion interface{} `xml:"operand"`
	}
	operations := []adGroupCriterionOperation{}
	for action, adGroupCriterions := range adGroupCriterionOperations {
		for _, adGroupCriterion := range adGroupCriterions {
			operations = append(operations,
				adGroupCriterionOperation{
					Action:           action,
					AdGroupCriterion: adGroupCriterion,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []adGroupCriterionOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: operations,
	}
	respBody, err := s.Auth.request(adGroupCriterionServiceUrl, "mutate", mutation)
	if err != nil {
		return adGroupCriterions, err
	}
	mutateResp := struct {
		AdGroupCriterions AdGroupCriterions `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return adGroupCriterions, err
	}

	return mutateResp.AdGroupCriterions, err
}

// MutateLabel allows you to add and removes labels from ad groups.
//
// Example
//
//  adGroupCriterions, err := adGroupCriterionService.MutateLabel(
//    gads.AdGroupCriterionLabelOperations{
//      "ADD": {
//        gads.AdGroupCriterionLabel{AdGroupCriterionId: 3200, LabelId: 5353},
//        gads.AdGroupCriterionLabel{AdGroupCriterionId: 4320, LabelId: 5643},
//      },
//      "REMOVE": {
//        gads.AdGroupCriterionLabel{AdGroupCriterionId: 3653, LabelId: 5653},
//      },
//    },
//  )
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupCriterionService#mutateLabel
//
func (s *AdGroupCriterionService) MutateLabel(adGroupCriterionLabelOperations AdGroupCriterionLabelOperations) (adGroupCriterionLabels []AdGroupCriterionLabel, err error) {
	type adGroupCriterionLabelOperation struct {
		Action                string                `xml:"operator"`
		AdGroupCriterionLabel AdGroupCriterionLabel `xml:"operand"`
	}
	operations := []adGroupCriterionLabelOperation{}
	for action, adGroupCriterionLabels := range adGroupCriterionLabelOperations {
		for _, adGroupCriterionLabel := range adGroupCriterionLabels {
			operations = append(operations,
				adGroupCriterionLabelOperation{
					Action:                action,
					AdGroupCriterionLabel: adGroupCriterionLabel,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []adGroupCriterionLabelOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutateLabel",
		},
		Ops: operations}
	respBody, err := s.Auth.request(adGroupCriterionServiceUrl, "mutateLabel", mutation)
	if err != nil {
		return adGroupCriterionLabels, err
	}
	mutateResp := struct {
		AdGroupCriterionLabels []AdGroupCriterionLabel `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return adGroupCriterionLabels, err
	}

	return mutateResp.AdGroupCriterionLabels, err
}

// Query is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupCriterionService#query
//
func (s *AdGroupCriterionService) Query(query string) (adGroupCriterions AdGroupCriterions, err error) {
	return adGroupCriterions, ERROR_NOT_YET_IMPLEMENTED
}
