package gads

import (
	"encoding/xml"
	"fmt"
)

var (
	AD_GROUP_SERVICE_URL = ServiceUrl{
		baseUrl,
		"AdGroupService",
	}
)

type adGroupService struct {
	Auth
}

func NewAdGroupService(auth Auth) *adGroupService {
	return &adGroupService{Auth: auth}
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

func (s *adGroupService) Get(selector Selector) (adGroups []AdGroup, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.Request(
		AD_GROUP_SERVICE_URL,
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
		return adGroups, err
	}
	getResp := struct {
		Size     int64     `xml:"rval>totalNumEntries"`
		AdGroups []AdGroup `xml:"rval>entries"`
	}{}
	fmt.Printf("%s\n", respBody)
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroups, err
	}
	return getResp.AdGroups, err
}

func (s *adGroupService) Mutate(adGroupOperations AdGroupOperations) (adGroups []AdGroup, err error) {
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
	respBody, err := s.Auth.Request(AD_GROUP_SERVICE_URL, "mutate", mutation)
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

func (s *adGroupService) Query(query string) (adGroups []AdGroup, err error) {
	return adGroups, err
}
