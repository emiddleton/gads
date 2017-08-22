package gads

import (
	"encoding/xml"
)

// A campaignService holds the connection information for the
// campaign service.
type CampaignService struct {
	Auth
}

// NewCampaignService creates a new campaignService
func NewCampaignService(auth *Auth) *CampaignService {
	return &CampaignService{Auth: *auth}
}

// ConversionOptimizerEligibility
//
// RejectionReasons can be any of
//   "CAMPAIGN_IS_NOT_ACTIVE", "NOT_CPC_CAMPAIGN","CONVERSION_TRACKING_NOT_ENABLED",
//   "NOT_ENOUGH_CONVERSIONS", "UNKNOWN"
//
type conversionOptimizerEligibility struct {
	Eligible         bool     `xml:"eligible"`         // is eligible for optimization
	RejectionReasons []string `xml:"rejectionReasons"` // reason for why campaign is
	// not eligible for conversion optimization.
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
	TargetGoogleSearch         bool `xml:"https://adwords.google.com/api/adwords/cm/v201708 targetGoogleSearch"`
	TargetSearchNetwork        bool `xml:"https://adwords.google.com/api/adwords/cm/v201708 targetSearchNetwork"`
	TargetContentNetwork       bool `xml:"https://adwords.google.com/api/adwords/cm/v201708 targetContentNetwork"`
	TargetPartnerSearchNetwork bool `xml:"https://adwords.google.com/api/adwords/cm/v201708 targetPartnerSearchNetwork"`
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
	StrategyName   string         `xml:"biddingStrategyName,omitempty"`
	StrategyType   string         `xml:"biddingStrategyType,omitempty"`
	StrategySource string         `xml:"biddingStrategySource,omitempty"`
	Scheme         *BiddingScheme `xml:"biddingScheme,omitempty"`
	Bids           []Bid          `xml:"bids"`
}

type CustomParameter struct {
	Key      string `xml:"key"`
	Value    string `xml:"value"`
	IsRemove bool   `xml:"isRemove"`
}

type CustomParameters struct {
	CustomParameters []CustomParameter `xml:"parameters"`
	DoReplace        bool              `xml:"doReplace"`
}

type Campaign struct {
	Id                             int64                           `xml:"id,omitempty"`
	Name                           string                          `xml:"name"`
	Status                         string                          `xml:"status"`                  // Status: "ENABLED", "PAUSED", "REMOVED"
	ServingStatus                  *string                         `xml:"servingStatus,omitempty"` // ServingStatus: "SERVING", "NONE", "ENDED", "PENDING", "SUSPENDED"
	StartDate                      string                          `xml:"startDate"`
	EndDate                        *string                         `xml:"endDate,omitempty"`
	BudgetId                       int64                           `xml:"budget>budgetId"`
	ConversionOptimizerEligibility *conversionOptimizerEligibility `xml:"conversionOptimizerEligibility"`
	AdServingOptimizationStatus    string                          `xml:"adServingOptimizationStatus"`
	FrequencyCap                   *FrequencyCap                   `xml:"frequencyCap"`
	Settings                       []CampaignSetting               `xml:"settings"`
	AdvertisingChannelType         string                          `xml:"advertisingChannelType,omitempty"`    // "UNKNOWN", "SEARCH", "DISPLAY", "SHOPPING"
	AdvertisingChannelSubType      *string                         `xml:"advertisingChannelSubType,omitempty"` // "UNKNOWN", "SEARCH_MOBILE_APP", "DISPLAY_MOBILE_APP", "SEARCH_EXPRESS", "DISPLAY_EXPRESS"
	NetworkSetting                 *NetworkSetting                 `xml:"networkSetting"`
	Labels                         []Label                         `xml:"labels"`
	BiddingStrategyConfiguration   *BiddingStrategyConfiguration   `xml:"biddingStrategyConfiguration"`
	ForwardCompatibilityMap        *map[string]string              `xml:"forwardCompatibilityMap,omitempty"`
	TrackingUrlTemplate            *string                         `xml:"trackingUrlTemplate"`
	UrlCustomParameters            *CustomParameters               `xml:"urlCustomParameters"`
	Errors                         []error                         `xml:"-"`
}

type CampaignOperations map[string][]Campaign

type CampaignLabel struct {
	CampaignId int64 `xml:"campaignId"`
	LabelId    int64 `xml:"labelId"`
}

type CampaignLabelOperations map[string][]CampaignLabel

// Get returns an array of Campaign's and the total number of campaign's matching
// the selector.
//
// Example
//
//   campaigns, totalCount, err := campaignService.Get(
//     gads.Selector{
//       Fields: []string{
//         "AdGroupId",
//         "Status",
//         "AdGroupCreativeApprovalStatus",
//         "AdGroupAdDisapprovalReasons",
//         "AdGroupAdTrademarkDisapproved",
//       },
//       Predicates: []gads.Predicate{
//         {"AdGroupId", "EQUALS", []string{adGroupId}},
//       },
//     },
//   )
//
// Selectable fields are
//   "Id", "Name", "Status", "ServingStatus", "StartDate", "EndDate", "AdServingOptimizationStatus",
//   "Settings", "AdvertisingChannelType", "AdvertisingChannelSubType", "Labels", "TrackingUrlTemplate",
//   "UrlCustomParameters"
//
// filterable fields are
//   "Id", "Name", "Status", "ServingStatus", "StartDate", "EndDate", "AdvertisingChannelType",
//   "AdvertisingChannelSubType", "Labels", "TrackingUrlTemplate"
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201708/CampaignService#get
//
func (s *CampaignService) Get(selector Selector) (campaigns []Campaign, totalCount int64, err error) {
	// The default namespace, "", will break in 1.5 with the addition of
	// custom namespace support.  Hence, we have to ensure that the baseUrl is
	// set again as the proper namespace for the service/serviceSelector element
	selector.XMLName = xml.Name{baseUrl, "serviceSelector"}
	respBody, err := s.Auth.request(
		campaignServiceUrl,
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
		nil,
	)
	if err != nil {
		return campaigns, totalCount, err
	}
	getResp := struct {
		Size      int64      `xml:"rval>totalNumEntries"`
		Campaigns []Campaign `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return campaigns, totalCount, err
	}
	return getResp.Campaigns, getResp.Size, err
}

// Mutate allows you to add and modify campaigns, returning the
// campaigns.  Note that the "REMOVE" operator is not supported.
// To remove a campaign set its Status to "REMOVED".
//
// Example
//
//  campaignNeedingRemoval.Status = "REMOVED"
//  ads, err := campaignService.Mutate(
//    gads.CampaignOperations{
//      "ADD": {
//        gads.Campaign{
//          Name: "my campaign name",
//          Status: "PAUSED",
//          StartDate: time.Now().Format("20060102"),
//          BudgetId: 321543214,
//          AdServingOptimizationStatus: "ROTATE_INDEFINITELY",
//          Settings: []gads.CampaignSetting{
//            gads.NewRealTimeBiddingSetting(true),
//          },
//          AdvertisingChannelType: "SEARCH",
//          BiddingStrategyConfiguration: &gads.BiddingStrategyConfiguration{
//            StrategyType: "MANUAL_CPC",
//          },
//        },
//        campaignNeedingRemoval,
//      },
//      "SET": {
//        modifiedCampaign,
//      },
//    }
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201708/CampaignService#mutate
//
func (s *CampaignService) Mutate(campaignOperations CampaignOperations) (campaigns []Campaign, err error) {
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
	respBody, err := s.Auth.request(campaignServiceUrl, "mutate", mutation, nil)
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

// Mutate allows you to add and removes labels from campaigns.
//
// Example
//
//  cls, err := campaignService.MutateLabel(
//    gads.CampaignOperations{
//      "ADD": {
//        gads.CampaignLabel{CampaignId: 3200, LabelId: 5353},
//        gads.CampaignLabel{CampaignId: 4320, LabelId: 5643},
//      },
//      "REMOVE": {
//        gads.CampaignLabel{CampaignId: 3653, LabelId: 5653},
//      },
//    }
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201708/CampaignService#mutateLabel
//
func (s *CampaignService) MutateLabel(campaignLabelOperations CampaignLabelOperations) (campaignLabels []CampaignLabel, err error) {
	type campaignLabelOperation struct {
		Action        string        `xml:"operator"`
		CampaignLabel CampaignLabel `xml:"operand"`
	}
	operations := []campaignLabelOperation{}
	for action, campaignLabels := range campaignLabelOperations {
		for _, campaignLabel := range campaignLabels {
			operations = append(operations,
				campaignLabelOperation{
					Action:        action,
					CampaignLabel: campaignLabel,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []campaignLabelOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutateLabel",
		},
		Ops: operations}
	respBody, err := s.Auth.request(campaignServiceUrl, "mutateLabel", mutation, nil)
	if err != nil {
		return campaignLabels, err
	}
	mutateResp := struct {
		CampaignLabels []CampaignLabel `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return campaignLabels, err
	}

	return mutateResp.CampaignLabels, err
}

// Query documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201708/CampaignService#query
//
func (s *CampaignService) Query(query string) (campaigns []Campaign, totalCount int64, err error) {

	respBody, err := s.Auth.request(
		campaignServiceUrl,
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
		return campaigns, totalCount, err
	}

	getResp := struct {
		Size      int64      `xml:"rval>totalNumEntries"`
		Campaigns []Campaign `xml:"rval>entries"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return campaigns, totalCount, err
	}
	return getResp.Campaigns, getResp.Size, err
}
