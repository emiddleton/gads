package gads

import (
	"encoding/xml"
	"fmt"
)

var (
	AD_GROUP_CRITERION_SERVICE_URL = "https://adwords.google.com/api/adwords/cm/v201309/AdGroupCriterionService"
)

type adGroupCriterionService struct {
	Auth
}

func NewAdGroupCriterionService(auth Auth) *adGroupCriterionService {
	return &adGroupCriterionService{Auth: auth}
}

type WebpageCondition struct {
	Operand  string `xml:"operand"`
	Argument string `xml:"argument"`
}

type WebpageParameter struct {
	CriterionName string             `xml:"criterionName"`
	Conditions    []WebpageCondition `xml:"conditions"`
}

type ProductCondition struct {
	Argument string `xml:"argument"`
	Operand  string `xml:"operand"`
}

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
	Type      string    `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
	AdGroupId int64     `xml:"adGroupId"`
	Criterion Criterion `xml:"criterion"`
}

type BiddableAdGroupCriterion struct {
	Type      string    `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
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

func NewBiddableAdGroupCriterion(adGroupId int64, criterion Criterion) BiddableAdGroupCriterion {
	return BiddableAdGroupCriterion{
		AdGroupId: adGroupId,
		Type:      "BiddableAdGroupCriterion",
		Criterion: criterion,
	}
}

func NewNegativeAdGroupCriterion(adGroupId int64, criterion Criterion) NegativeAdGroupCriterion {
	return NegativeAdGroupCriterion{
		AdGroupId: adGroupId,
		Type:      "NegativeAdGroupCriterion",
		Criterion: criterion,
	}
}

func NewAgeRangeCriterion(ageRangeId int64) Criterion {
	return Criterion{
		Id:   ageRangeId,
		Type: "AgeRange",
	}
}

func NewGenderCriterion(genderId int64) Criterion {
	return Criterion{
		Id:   genderId,
		Type: "Gender",
	}
}

func NewKeywordCriterion(text, matchType string) Criterion {
	return Criterion{
		Type:      "Keyword",
		Text:      &text,
		MatchType: &matchType,
	}
}

func NewMobileAppCategoryCriterion(mobileAppCategoryId int64, displayName string) Criterion {
	return Criterion{
		Type:                "MobileAppCategory",
		MobileAppCategoryId: &mobileAppCategoryId,
		DisplayName:         &displayName,
	}
}

func NewMobileApplicationCriterion(appId string) Criterion {
	return Criterion{
		Type:  "MobileApplication",
		AppId: &appId,
	}
}

func NewPlacementCriterion(url string) Criterion {
	return Criterion{
		Type: "Placement",
		Url:  &url,
	}
}

func NewUserInterestCriterion(userInterestId int64, userInterestName string) Criterion {
	return Criterion{
		Type:             "CriterionUserInterest",
		UserInterestId:   &userInterestId,
		UserInterestName: &userInterestName,
	}
}

func NewUserListCriterion(userListId int64, userListName, userListMembershipStatus string) Criterion {
	return Criterion{
		Type:                     "CriterionUserList",
		UserListId:               &userListId,
		UserListName:             &userListName,
		UserListMembershipStatus: &userListMembershipStatus,
	}
}

func NewVerticalCriterion(verticalId, verticalParentId int64, path []string) Criterion {
	return Criterion{
		Type:             "Vertical",
		VerticalId:       &verticalId,
		VerticalParentId: &verticalParentId,
		Path:             &path,
	}
}

func NewWebpageCriterion(criterionName string, conditions []WebpageCondition) Criterion {
	return Criterion{
		Type: "Webpage",
		Parameter: &WebpageParameter{
			CriterionName: criterionName,
			Conditions:    conditions,
		},
	}
}

func NewProductCriterion(text string, conditions []ProductCondition) Criterion {
	return Criterion{
		Type:       "Product",
		Text:       &text,
		Conditions: conditions,
	}
}

func (s adGroupCriterionService) Get(selector Selector) (adGroupCriterions AdGroupCriterions, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.Request(
		AD_GROUP_CRITERION_SERVICE_URL,
		"get",
		struct {
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201309 get"`
			Sel     Selector
		}{Sel: selector},
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
		XMLName xml.Name                    `xml:"https://adwords.google.com/api/adwords/cm/v201309 mutate"`
		Ops     []adGroupCriterionOperation `xml:"operations"`
	}{Ops: operations}
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
