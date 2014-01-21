package gads

import (
	"encoding/xml"
)

var (
	CAMPAIGN_SERVICE_URL = "https://adwords.google.com/api/adwords/cm/v201309/CampaignService"
)

type campaignService struct {
	Auth
}

func NewCampaignService(auth Auth) *campaignService {
	return &campaignService{Auth: auth}
}

type ConversionOptimizerEligibility struct {
	Eligible         bool     `xml:"eligible"`
	RejectionReasons []string `xml:"rejectionReasons"`
}

type FrequencyCap struct {
	Impressions int64  `xml:"impressions"`
	TimeUnit    string `xml:"timeUnit"`
	Level       string `xml:"level,omitempty"`
}

type Setting struct {
	XMLName xml.Name `xml:"settings"`
	Type    string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`

	// GeoTargetTypeSetting
	PositiveGeoTargetType *string `xml:"positiveGeoTargetType,omitempty"`
	NegativeGeoTargetType *string `xml:"negativeGeoTargetType,omitempty"`

	// KeywordMatchSetting
	OptIn *bool `xml:"optIn,omitempty"`

	// TrackingSetting
	TrackingUrl *string `xml:"trackingUrl,omitempty"`
}

func NewGeoTargetTypeSetting(positiveGeoTargetType, negativeGeoTargetType string) Setting {
	return Setting{
		Type: "GeoTargetTypeSetting",
		PositiveGeoTargetType: &positiveGeoTargetType,
		NegativeGeoTargetType: &negativeGeoTargetType,
	}
}

func NewKeywordMatchSetting(optIn bool) Setting {
	return Setting{
		Type:  "KeywordMatchSetting",
		OptIn: &optIn,
	}
}

func NewRealTimeBiddingSetting(optIn bool) Setting {
	return Setting{
		Type:  "RealTimeBiddingSetting",
		OptIn: &optIn,
	}
}

func NewTrackingSetting(trackingUrl string) Setting {
	return Setting{
		Type:        "TrackingSetting",
		TrackingUrl: &trackingUrl,
	}
}

type NetworkSetting struct {
	TargetGoogleSearch         bool `xml:"targetGoogleSearch"`
	TargetSearchNetwork        bool `xml:"targetSearchNetwork"`
	TargetContentNetwork       bool `xml:"targetContentNetwork"`
	TargetPartnerSearchNetwork bool `xml:"targetPartnerSearchNetwork"`
}

type BiddingScheme struct {
	Type               string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	EnhancedCpcEnabled bool   `xml:"enhancedCpcEnabled"`
}

type Bid struct {
	Type         string  `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	Amount       int64   `xml:"bid>microAmount"`
	CpcBidSource *string `xml:"cpcBidSource"`
	CpmBidSource *string `xml:"cpmBidSource"`
}

type BiddingStrategyConfiguration struct {
	StrategyType   string         `xml:"biddingStrategyType"`
	StrategySource string         `xml:"biddingStrategySource,omitempty"`
	Scheme         *BiddingScheme `xml:"biddingScheme,omitempty"`
	Bids           []Bid          `xml:"bids"`
}

type Campaign struct {
	Id                             int64                           `xml:"id,omitempty"`
	Name                           string                          `xml:"name"`
	Status                         string                          `xml:"status"`
	ServingStatus                  *string                         `xml:"servingStatus,omitempty"`
	StartDate                      string                          `xml:"startDate"`
	EndDate                        *string                         `xml:"endDate,omitempty"`
	BudgetId                       int64                           `xml:"budget>budgetId"`
	ConversionOptimizerEligibility *ConversionOptimizerEligibility `xml:"conversionOptimizerEligibility"`
	AdServingOptimizationStatus    string                          `xml:"adServingOptimizationStatus"`
	FrequencyCap                   *FrequencyCap                   `xml:"frequencyCap"`
	Settings                       []Setting                       `xml:"settings"`
	NetworkSetting                 *NetworkSetting                 `xml:"networkSetting"`
	BiddingStrategyConfiguration   *BiddingStrategyConfiguration   `xml:"biddingStrategyConfiguration"`
	//	ForwardCompatibilityMap        *map[string]string              `xml:"forwardCompatibilityMap,omitempty"`
}

type CampaignOperations map[string][]Campaign

func (s *campaignService) Get(selector Selector) (campaigns []Campaign, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.Request(
		CAMPAIGN_SERVICE_URL,
		"get",
		struct {
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201309 get"`
			Sel     Selector
		}{Sel: selector},
	)
	if err != nil {
		return campaigns, err
	}
	getResp := struct {
		Size      int64      `xml:"rval>totalNumEntries"`
		Campaigns []Campaign `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return campaigns, err
	}
	return getResp.Campaigns, err
}

func (s *campaignService) Mutate(campaignOperations CampaignOperations) (campaigns []Campaign, err error) {
	type campaignOperation struct {
		Action   string   `xml:"operator"`
		Campaign Campaign `xml:"operand"`
	}
	operations := []campaignOperation{}
	for action, campaigns := range campaignOperations {
		for _, campaign := range campaigns {
			operations = append(operations,
				campaignOperation{
					Action:   action,
					Campaign: campaign,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name            `xml:"https://adwords.google.com/api/adwords/cm/v201309 mutate"`
		Ops     []campaignOperation `xml:"operations"`
	}{Ops: operations}
	respBody, err := s.Auth.Request(CAMPAIGN_SERVICE_URL, "mutate", mutation)
	if err != nil {
		return campaigns, err
	}
	mutateResp := struct {
		Campaigns []Campaign `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return campaigns, err
	}

	return mutateResp.Campaigns, err
}
