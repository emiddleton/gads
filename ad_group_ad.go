package gads

import "encoding/xml"

type AdGroupAdService struct {
	Auth
}

type AppUrl struct {
	Url    string `xml:"url"`
	OsType string `xml:"osType"` // "OS_TYPE_IOS", "OS_TYPE_ANDROID", "UNKNOWN"
}

type AdGroupAd struct {
	AdGroupId            int64    `xml:"adGroupId"`
	Ad                   Ad       `xml:"ad"`
	Status               string   `xml:"status,omitempty"`
	ApprovalStatus       string   `xml:"approvalStatus,omitempty"`
	DisapprovalReasons   []string `xml:"disapprovalReasons,omitempty"`
	TrademarkDisapproved bool     `xml:"trademarkDisapproved,omitempty"`
	Labels               []Label  `xml:"labels,omitempty"`
}

type TemplateElementField struct {
	Name       string `xml:"name"`
	Type       string `xml:"type"`
	FieldText  string `xml:"fieldText"`
	FieldMedia string `xml:"fieldMedia"`
}

type TemplateElement struct {
	UniqueName string                 `xml:"uniqueName"`
	Fields     []TemplateElementField `xml:"fields"`
}

type AdGroupAdOperations map[string]AdGroupAds

func NewAdGroupAdService(auth *Auth) *AdGroupAdService {
	return &AdGroupAdService{Auth: *auth}
}

func NewAdGroupTextAd(
	adGroupId int64,
	URL string,
	displayURL string,
	headline string,
	description1 string,
	description2 string,
	status string,
) AdGroupAd {
	return AdGroupAd{
		AdGroupId: adGroupId,
		Ad: TextAd{
			CommonAd: CommonAd{
				URL:        URL,
				DisplayURL: displayURL,
			},
			Headline:     headline,
			Description1: description1,
			Description2: description2,
		},
		Status: status,
	}
}

type AdGroupAdLabel struct {
	AdGroupAdId int64 `xml:"adGroupAdId"`
	LabelId     int64 `xml:"labelId"`
}

type AdGroupAdLabelOperations map[string][]AdGroupAdLabel

type AdURLUpgrade struct {
	AdId                int64  `xml:"adId"`
	FinalURL            string `xml:"finalUrl,omitempty"`
	FinalMobileUrl      string `xml:"finalMobileUrl,omitempty"`
	TrackingURLTemplate string `xml:"trackingUrlTemplate"`
}

// Get returns an array of ad's and the total number of ad's matching
// the selector.
//
// Example
//
//   ads, totalCount, err := adGroupAdService.Get(
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
//   "AdGroupId", "Id", "Url", "DisplayUrl", "CreativeFinalUrls", "CreativeFinalMobileUrls",
//   "CreativeFinalAppUrls", "CreativeTrackingUrlTemplate", "CreativeUrlCustomParameters",
//   "DevicePreference", "Status", "AdGroupCreativeApprovalStatus", "AdGroupAdDisapprovalReasons"
//   "AdGroupAdTrademarkDisapproved", "Labels"
//
//   TextAd
//     "Headline", "Description1", "Description2"
//
//   ImageAd
//     "ImageCreativeName"
//
// filterable fields are
//   "AdGroupId", "Id", "Url", "DisplayUrl", "CreativeFinalUrls", "CreativeFinalMobileUrls",
//   "CreativeFinalAppUrls", "CreativeTrackingUrlTemplate", "CreativeUrlCustomParameters",
//   "DevicePreference", "Status", "AdGroupCreativeApprovalStatus", "AdGroupAdDisapprovalReasons"
//   "Labels"
//
//   TextAd specific fields
//     "Headline", "Description1", "Description2"
//
//   ImageAd specific fields
//     "ImageCreativeName"
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#get
//
func (s AdGroupAdService) Get(selector Selector) (adGroupAds AdGroupAds, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		adGroupAdServiceUrl,
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
		return adGroupAds, totalCount, err
	}
	getResp := struct {
		Size       int64      `xml:"rval>totalNumEntries"`
		AdGroupAds AdGroupAds `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroupAds, totalCount, err
	}
	return getResp.AdGroupAds, getResp.Size, err
}

// Mutate allows you to add, modify and remove ads, returning the
// modified ads.
//
// Example
//
//  ads, err := adGroupAdService.Mutate(
//    gads.AdGroupAdOperations{
//      "ADD": {
//        gads.NewAdGroupTextAd(
//          adGroup.Id,
//          "https://classdo.com/en",
//          "classdo.com",
//          "test headline",
//          "test line one",
//          "test line two",
//          "PAUSED",
//        ),
//      },
//      "SET": {
//        modifiedAd,
//      },
//      "REMOVE": {
//        adNeedingRemoval,
//      },
//    },
//  )
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#mutate
//
func (s *AdGroupAdService) Mutate(adGroupAdOperations AdGroupAdOperations) (adGroupAds AdGroupAds, err error) {
	type adGroupAdOperation struct {
		Action    string    `xml:"operator"`
		AdGroupAd AdGroupAd `xml:"operand"`
	}
	operations := []adGroupAdOperation{}
	for action, adGroupAds := range adGroupAdOperations {
		for _, adGroupAd := range adGroupAds {
			operations = append(operations,
				adGroupAdOperation{
					Action:    action,
					AdGroupAd: adGroupAd,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []adGroupAdOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: operations,
	}
	respBody, err := s.Auth.request(adGroupAdServiceUrl, "mutate", mutation)
	if err != nil {
		return adGroupAds, err
	}
	mutateResp := struct {
		AdGroupAds AdGroupAds `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return adGroupAds, err
	}
	return mutateResp.AdGroupAds, err
}

// MutateLabel allows you to add and removes labels from ads.
//
// Example
//
//  ads, err := adGroupAdService.MutateLabel(
//    gads.AdGroupAdLabelOperations{
//      "ADD": {
//        gads.AdGroupAdLabel{AdGroupAdId: 3200, LabelId: 5353},
//        gads.AdGroupAdLabel{AdGroupAdId: 4320, LabelId: 5643},
//      },
//      "REMOVE": {
//        gads.AdGroupAdLabel{AdGroupAdId: 3653, LabelId: 5653},
//      },
//    }
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#mutateLabel
//
func (s *AdGroupAdService) MutateLabel(adGroupAdLabelOperations AdGroupAdLabelOperations) (adGroupAdLabels []AdGroupAdLabel, err error) {
	type adGroupAdLabelOperation struct {
		Action         string         `xml:"operator"`
		AdGroupAdLabel AdGroupAdLabel `xml:"operand"`
	}
	operations := []adGroupAdLabelOperation{}
	for action, adGroupAdLabels := range adGroupAdLabelOperations {
		for _, adGroupAdLabel := range adGroupAdLabels {
			operations = append(operations,
				adGroupAdLabelOperation{
					Action:         action,
					AdGroupAdLabel: adGroupAdLabel,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []adGroupAdLabelOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutateLabel",
		},
		Ops: operations}
	respBody, err := s.Auth.request(adGroupAdServiceUrl, "mutateLabel", mutation)
	if err != nil {
		return adGroupAdLabels, err
	}
	mutateResp := struct {
		AdGroupAdLabels []AdGroupAdLabel `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return adGroupAdLabels, err
	}

	return mutateResp.AdGroupAdLabels, err
}

// Query is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#query
//
func (s *AdGroupAdService) Query(query string) (adGroupAds AdGroupAds, totalCount int64, err error) {
	return adGroupAds, totalCount, ERROR_NOT_YET_IMPLEMENTED
}

//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#upgradeUrl
//
func (s *AdGroupAdService) UpgradeUrl(adURLUpgrades []AdURLUpgrade) (ads []Ad, err error) {
	upgrade := struct {
		XMLName xml.Name
		Ops     []AdURLUpgrade `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "upgradeUrl",
		},
		Ops: adURLUpgrades,
	}

	respBody, err := s.Auth.request(adGroupAdServiceUrl, "upgradeUrl", upgrade)

	if err != nil {
		return ads, err
	}
	upgradeResp := struct {
		Ads []Ad `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &upgradeResp)
	if err != nil {
		return ads, err
	}
	return upgradeResp.Ads, nil
}
