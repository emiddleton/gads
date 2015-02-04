package gads

import (
	"encoding/xml"
	"fmt"
)

var (
	AD_GROUP_CRITERION_SERVICE_URL = ServiceUrl{
		baseUrl,
		"AdGroupCriterionService",
	}
)

type adGroupCriterionService struct {
	Auth
}

func NewAdGroupCriterionService(auth Auth) *adGroupCriterionService {
	return &adGroupCriterionService{Auth: auth}
}

/*
type Criterion struct {
	Id   int64  `xml:"id,omitempty"`
	Type string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`

	// AgeRange
	AgeRangeType *string `xml:"ageRangeType,omitempty"`

	// CriterionUserInterest
	UserInterestId   *int64  `xml:"userInterestId,omitempty"`
	UserInterestName *string `xml:"userInterestName,omitempty"`

	// CriterionUserList
	UserListId               *int64  `xml:"userListId,omitempty"`
	UserListName             *string `xml:"userListName,omitempty"`
	UserListMembershipStatus *string `xml:"userListMembershipStatus,omitempty"`

	// Gender
	GenderType *string `xml:"genderType,omitempty"`

	// Product
	Conditions []ProductCondition `xml:"conditions,omitempty"`

	Text *string `xml:"text,omitempty"`

	// Keyword
	MatchType *string `xml:"matchType,omitempty"`

	// MobileApplication
	AppId *string `xml:"appId,omitempty"`

	// MobileAppCategory
	MobileAppCategoryId *int64  `xml:"mobileAppCategoryId,omitempty"`
	DisplayName         *string `xml:"displayName,omitempty"`

	// Placement
	Url *string `xml:"url,omitempty"`

	// Vertical
	VerticalId       *int64    `xml:"verticalId,omitempty"`
	VerticalParentId *int64    `xml:"verticalParentId,omitempty"`
	Path             *[]string `xml:"path,omitempty"`

	// Webpage
	Parameter        *WebpageParameter `xml:"parameter,omitempty"`
	CriteriaCoverage *float64          `xml:"criteriaCoverage,omitempty"`
	CriteriaSamples  []string          `xml:"criteriaSamples,omitempty"`
}
*/
type QualityInfo struct {
	IsKeywordAdRelevanceAcceptable bool  `xml:"isKeywordAdRelevanceAcceptable,omitempty"`
	IsLandingPageQualityAcceptable bool  `xml:"isLandingPageQualityAcceptable,omitempty"`
	IsLandingPageLatencyAcceptable bool  `xml:"isLandingPageLatencyAcceptable,omitempty"`
	QualityScore                   int64 `xml:"QualityScore,omitempty"`
}

type Cpc struct {
	Amount int64 `xml:"amount"`
}

type NegativeAdGroupCriterion struct {
	AdGroupId int64     `xml:"adGroupId"`
	Criterion Criterion `xml:"criterion"`
}

func (nagc NegativeAdGroupCriterion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(
		start.Attr,
		xml.Attr{
			xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"},
			"NegativeAdGroupCriterion",
		},
	)
	e.EncodeToken(start)
	e.EncodeElement(&nagc.AdGroupId, xml.StartElement{Name: xml.Name{"", "adGroupId"}})
	criterionMarshalXML(nagc.Criterion, e)
	e.EncodeToken(start.End())
	return nil
}

func (nagc *NegativeAdGroupCriterion) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		if err != nil {
			return err
		}
		switch start := token.(type) {
		case xml.StartElement:
			tag := start.Name.Local
			switch tag {
			case "adGroupId":
				if err := dec.DecodeElement(&nagc.AdGroupId, &start); err != nil {
					return err
				}
			case "criterion":
				criterion, err := criterionUnmarshalXML(dec, start)
				if err != nil {
					return err
				}
				nagc.Criterion = criterion
			case "AdGroupCriterion.Type":
				break
			default:
				return fmt.Errorf("unknown NegativeAdGroupCriterion field %s", tag)
			}
		}
	}
	return nil
}

type BiddableAdGroupCriterion struct {
	AdGroupId int64     `xml:"adGroupId"`
	Criterion Criterion `xml:"criterion"`

	// BiddableAdGroupCriterion
	UserStatus          string   `xml:"userStatus,omitempty"`
	SystemServingStatus string   `xml:"systemServingStatus,omitempty"`
	ApprovalStatus      string   `xml:"approvalStatus,omitempty"`
	DisapprovalReasons  []string `xml:"disapprovalReasons,omitempty"`
	DestinationUrl      string   `xml:"destinationUrl,omitempty"`

	FirstPageCpc *Cpc `xml:"firstPageCpc>amount,omitempty"`
	TopOfPageCpc *Cpc `xml:"topOfPageCpc>amount,omitempty"`

	QualityInfo *QualityInfo `xml:"qualityInfo,omitempty"`

	BiddingStrategyConfiguration *BiddingStrategyConfiguration `xml:"biddingStrategyConfiguration,omitempty"`
	BidModifier                  int64                         `xml:"bidModifier,omitempty"`
}

type AdGroupCriterions []interface{}

func (bagc BiddableAdGroupCriterion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(
		start.Attr,
		xml.Attr{
			xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"},
			"BiddableAdGroupCriterion",
		},
	)
	e.EncodeToken(start)
	e.EncodeElement(&bagc.AdGroupId, xml.StartElement{Name: xml.Name{"", "adGroupId"}})
	criterionMarshalXML(bagc.Criterion, e)
	e.EncodeElement(&bagc.UserStatus, xml.StartElement{Name: xml.Name{"", "userStatus"}})
	if bagc.DestinationUrl != "" {
		e.EncodeElement(&bagc.DestinationUrl, xml.StartElement{Name: xml.Name{"", "destinationUrl"}})
	}
	e.EncodeElement(&bagc.BiddingStrategyConfiguration, xml.StartElement{Name: xml.Name{"", "biddingStrategyConfiguration"}})
	if bagc.BidModifier != 0 {
		e.EncodeElement(&bagc.BidModifier, xml.StartElement{Name: xml.Name{"", "bidModifier"}})
	}
	e.EncodeToken(start.End())
	return nil
}

func (bagc *BiddableAdGroupCriterion) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		if err != nil {
			return err
		}
		switch start := token.(type) {
		case xml.StartElement:
			tag := start.Name.Local
			switch tag {
			case "adGroupId":
				if err := dec.DecodeElement(&bagc.AdGroupId, &start); err != nil {
					return err
				}
			case "criterion":
				criterion, err := criterionUnmarshalXML(dec, start)
				if err != nil {
					return err
				}
				bagc.Criterion = criterion
			case "userStatus":
				if err := dec.DecodeElement(&bagc.UserStatus, &start); err != nil {
					return err
				}
			case "systemServingStatus":
				if err := dec.DecodeElement(&bagc.SystemServingStatus, &start); err != nil {
					return err
				}
			case "approvalStatus":
				if err := dec.DecodeElement(&bagc.ApprovalStatus, &start); err != nil {
					return err
				}
			case "disapprovalReasons":
				if err := dec.DecodeElement(&bagc.DisapprovalReasons, &start); err != nil {
					return err
				}
			case "destinationUrl":
				if err := dec.DecodeElement(&bagc.DestinationUrl, &start); err != nil {
					return err
				}
			case "firstPageCpc":
				if err := dec.DecodeElement(&bagc.FirstPageCpc, &start); err != nil {
					return err
				}
			case "topOfPageCpc":
				if err := dec.DecodeElement(&bagc.TopOfPageCpc, &start); err != nil {
					return err
				}
			case "qualityInfo":
				if err := dec.DecodeElement(&bagc.QualityInfo, &start); err != nil {
					return err
				}
			case "biddingStrategyConfiguration":
				if err := dec.DecodeElement(&bagc.BiddingStrategyConfiguration, &start); err != nil {
					return err
				}
			case "bidModifier":
				if err := dec.DecodeElement(&bagc.BidModifier, &start); err != nil {
					return err
				}
			case "AdGroupCriterion.Type":
				break
			default:
				return fmt.Errorf("unknown BiddableAdGroupCriterion field %s", tag)
			}
		}
	}
	return nil
}

func (agcs *AdGroupCriterions) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	adGroupCriterionType, err := findAttr(start.Attr, xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"})
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

func (s adGroupCriterionService) Get(selector Selector) (adGroupCriterions AdGroupCriterions, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.Request(
		AD_GROUP_CRITERION_SERVICE_URL,
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
		return adGroupCriterions, err
	}
	getResp := struct {
		Size              int64             `xml:"rval>totalNumEntries"`
		AdGroupCriterions AdGroupCriterions `xml:"rval>entries"`
	}{}
	fmt.Printf("%s\n", respBody)
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroupCriterions, err
	}
	return getResp.AdGroupCriterions, err
}

func (s *adGroupCriterionService) Mutate(adGroupCriterionOperations AdGroupCriterionOperations) (adGroupCriterions AdGroupCriterions, err error) {
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
	respBody, err := s.Auth.Request(AD_GROUP_CRITERION_SERVICE_URL, "mutate", mutation)
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
