package gads

import (
	"encoding/xml"
)

var (
	CAMPAIGN_SERVICE_URL = ServiceUrl{
		Url:  baseUrl,
		Name: "CampaignService",
	}
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

type CampaignSetting struct {
	XMLName xml.Name `xml:"settings"`
	Type    string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`

	// GeoTargetTypeSetting
	PositiveGeoTargetType *string `xml:"positiveGeoTargetType,omitempty"`
	NegativeGeoTargetType *string `xml:"negativeGeoTargetType,omitempty"`

	// RealTimeBiddingSetting
	OptIn *bool `xml:"optIn,omitempty"`

	// DynamicSearchAdsSetting
	DomainName   *string `xml:"domainName,omitempty"`
	LanguageCode *string `xml:"langaugeCode,omitempty"`

	// TrackingSetting
	TrackingUrl *string `xml:"trackingUrl,omitempty"`
}

func NewDynamicSearchAdsSetting(domainName, languageCode string) CampaignSetting {
	return CampaignSetting{
		Type:         "DynamicSearchAdsSetting",
		DomainName:   &domainName,
		LanguageCode: &languageCode,
	}
}

func NewGeoTargetTypeSetting(positiveGeoTargetType, negativeGeoTargetType string) CampaignSetting {
	return CampaignSetting{
		Type: "GeoTargetTypeSetting",
		PositiveGeoTargetType: &positiveGeoTargetType,
		NegativeGeoTargetType: &negativeGeoTargetType,
	}
}

func NewRealTimeBiddingSetting(optIn bool) CampaignSetting {
	return CampaignSetting{
		Type:  "RealTimeBiddingSetting",
		OptIn: &optIn,
	}
}

func NewTrackingSetting(trackingUrl string) CampaignSetting {
	return CampaignSetting{
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
	StrategyId     int64          `xml:"biddingStrategyId,omitempty"`
	StrategyName   string         `xml:"biddingStrategyName"`
	StrategyType   string         `xml:"biddingStrategyType"`
	StrategySource string         `xml:"biddingStrategySource,omitempty"`
	Scheme         *BiddingScheme `xml:"biddingScheme,omitempty"`
	Bids           []Bid          `xml:"bids"`
}

// Status: ENABLED, PAUSED, REMOVED
// ServingStatus: SERVING, NONE, ENDED, PENDING, SUSPENDED
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
	Settings                       []CampaignSetting               `xml:"settings"`
	AdvertisingChannelType         string                          `xml:"advertisingChannelType,omitempty"`
	NetworkSetting                 *NetworkSetting                 `xml:"networkSetting"`
	BiddingStrategyConfiguration   *BiddingStrategyConfiguration   `xml:"biddingStrategyConfiguration"`
	//	ForwardCompatibilityMap        *map[string]string              `xml:"forwardCompatibilityMap,omitempty"`
	Errors []error `xml:"-"`
}

type CampaignOperations map[string][]Campaign

func (s *campaignService) Get(selector Selector) (campaigns []Campaign, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.Request(
		CAMPAIGN_SERVICE_URL,
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
		XMLName xml.Name
		Ops     []campaignOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: operations}
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
