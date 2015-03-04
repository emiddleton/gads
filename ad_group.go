package gads

import (
	"encoding/xml"
	"fmt"
)

type AdGroupService struct {
	Auth
}

func NewAdGroupService(auth *Auth) *AdGroupService {
	return &AdGroupService{Auth: *auth}
}

type TargetSettingDetail struct {
	CriterionTypeGroup string `xml:"criterionTypeGroup"`
	TargetAll          bool   `xml:"targetAll"`
}

type AdSetting struct {
	XMLName xml.Name `xml:"settings"`
	Type    string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`

	OptIn   *bool                 `xml:"optIn"`
	Details []TargetSettingDetail `xml:"details"`
}

type AdGroup struct {
	Id                           int64                          `xml:"id,omitempty"`
	CampaignId                   int64                          `xml:"campaignId"`
	CampaignName                 string                         `xml:"campaignName,omitempty"`
	Name                         string                         `xml:"name"`
	Status                       string                         `xml:"status"`
	Settings                     []AdSetting                    `xml:"settings,omitempty"`
	BiddingStrategyConfiguration []BiddingStrategyConfiguration `xml:"biddingStrategyConfiguration"`
	ContentBidCriterionTypeGroup *string                        `xml:"contentBidCriterionTypeGroup"`
}

type AdGroupOperations map[string][]AdGroup

type AdGroupLabel struct {
	AdGroupId int64 `xml:"adGroupId"`
	LabelId   int64 `xml:"labelId"`
}

type AdGroupLabelOperations map[string][]AdGroupLabel

// Get returns an array of ad group's and the total number of ad group's matching
// the selector.
//
// Example
//
//   ads, totalCount, err := adGroupService.Get(
//     gads.Selector{
//       Fields: []string{
//         "Id",
//         "CampaignId",
//         "CampaignName",
//         "Name",
//         "Status",
//         "Settings",
//         "Labels",
//         "ContentBidCriterionTypeGroup",
//         "TrackingUrlTemplate",
//         "UrlCustomParameters",
//       },
//       Predicates: []gads.Predicate{
//         {"Id", "EQUALS", []string{adGroupId}},
//       },
//     },
//   )
//
// Selectable fields are
//   "Id", "CampaignId", "CampaignName", "Name", "Status", "Settings", "Labels"
//   "ContentBidCriterionTypeGroup", "TrackingUrlTemplate", "UrlCustomParameters"
//
// filterable fields are
//   "Id", "CampaignId", "CampaignName", "Name", "Status", "Labels"
//   "ContentBidCriterionTypeGroup", "TrackingUrlTemplate"
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupService#get
//
func (s *AdGroupService) Get(selector Selector) (adGroups []AdGroup, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		adGroupServiceUrl,
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
		return adGroups, totalCount, err
	}
	getResp := struct {
		Size     int64     `xml:"rval>totalNumEntries"`
		AdGroups []AdGroup `xml:"rval>entries"`
	}{}
	fmt.Printf("%s\n", respBody)
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroups, totalCount, err
	}
	return getResp.AdGroups, getResp.Size, err
}

// Mutate allows you to add, modify and remove ad group's, returning the
// modified ad group's.
//
// Example
//
//  ads, err := adGroupService.Mutate(
//    gads.AdGroupOperations{
//      "ADD": {
//        gads.AdGroup{
//          Name:       "my ad group ",
//          Status:     "PAUSED",
//          CampaignId:  campaignId,
//          BiddingStrategyConfiguration: []gads.BiddingStrategyConfiguration{
//            gads.BiddingStrategyConfiguration{
//              StrategyType: "MANUAL_CPC",
//              Bids: []gads.Bid{
//                gads.Bid{
//                  Type:   "CpcBid",
//                  Amount: 10000,
//                },
//              },
//            },
//          },
//        },
//      },
//      "SET": {
//        modifiedAdGroup,
//      },
//      "REMOVE": {
//        adGroupNeedingRemoval,
//      },
//    },
//  )
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupService#mutate
//
func (s *AdGroupService) Mutate(adGroupOperations AdGroupOperations) (adGroups []AdGroup, err error) {
	type adGroupOperation struct {
		Action  string  `xml:"operator"`
		AdGroup AdGroup `xml:"operand"`
	}
	operations := []adGroupOperation{}
	for action, adGroups := range adGroupOperations {
		for _, adGroup := range adGroups {
			operations = append(operations,
				adGroupOperation{
					Action:  action,
					AdGroup: adGroup,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []adGroupOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: operations,
	}
	respBody, err := s.Auth.request(adGroupServiceUrl, "mutate", mutation)
	if err != nil {
		return adGroups, err
	}
	mutateResp := struct {
		AdGroups []AdGroup `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return adGroups, err
	}

	return mutateResp.AdGroups, err
}

// MutateLabel allows you to add and removes labels from ad groups.
//
// Example
//
//  adGroups, err := adGroupService.MutateLabel(
//    gads.AdGroupLabelOperations{
//      "ADD": {
//        gads.AdGroupLabel{AdGroupId: 3200, LabelId: 5353},
//        gads.AdGroupLabel{AdGroupId: 4320, LabelId: 5643},
//      },
//      "REMOVE": {
//        gads.AdGroupLabel{AdGroupId: 3653, LabelId: 5653},
//      },
//    }
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupService#mutateLabel
//
func (s *AdGroupService) MutateLabel(adGroupLabelOperations AdGroupLabelOperations) (adGroupLabels []AdGroupLabel, err error) {
	type adGroupLabelOperation struct {
		Action       string       `xml:"operator"`
		AdGroupLabel AdGroupLabel `xml:"operand"`
	}
	operations := []adGroupLabelOperation{}
	for action, adGroupLabels := range adGroupLabelOperations {
		for _, adGroupLabel := range adGroupLabels {
			operations = append(operations,
				adGroupLabelOperation{
					Action:       action,
					AdGroupLabel: adGroupLabel,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []adGroupLabelOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutateLabel",
		},
		Ops: operations}
	respBody, err := s.Auth.request(adGroupServiceUrl, "mutateLabel", mutation)
	if err != nil {
		return adGroupLabels, err
	}
	mutateResp := struct {
		AdGroupLabels []AdGroupLabel `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return adGroupLabels, err
	}

	return mutateResp.AdGroupLabels, err
}

// Query is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupService#query
//
func (s *AdGroupService) Query(query string) (adGroups []AdGroup, err error) {
	return adGroups, ERROR_NOT_YET_IMPLEMENTED
}
