package gads

import (
	"encoding/xml"
	"fmt"
)

// DayOfWeek: MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY
// StartHour: 0~23 inclusive
// StartMinute: ZERO, FIFTEEN, THIRTY, FORTY_FIVE
// EndHour: 0~24 inclusive
// EndMinute: ZERO, FIFTEEN, THIRTY, FORTY_FIVE
type AdScheduleCriterion struct {
	Id          int64  `xml:"id,omitempty"`
	DayOfWeek   string `xml:"dayOfWeek"`
	StartHour   string `xml:"startHour"`
	StartMinute string `xml:"startMinute"`
	EndHour     string `xml:"endHour"`
	EndMinute   string `xml:"endMinute"`
}

// AgeRangeType: AGE_RANGE_18_24, AGE_RANGE_25_34, AGE_RANGE_35_44, AGE_RANGE_45_54, AGE_RANGE_55_64, AGE_RANGE_65_UP, AGE_RANGE_UNDETERMINED, UNKNOWN
type AgeRangeCriterion struct {
	Id           int64  `xml:"id,omitempty"`
	AgeRangeType string `xml:"ageRangeType"`
}

type CarrierCriterion struct {
	Id          int64  `xml:"id,omitempty"`
	Name        string `xml:"name,emitempty"`
	CountryCode string `xml:"countryCode,emitempty"`
}

type ContentLabelCriterion struct {
	Id               int64  `xml:"id,omitempty"`
	ContentLabelType string `xml:"contentLabelType"` // ContentLabelType: "ADULTISH", "AFE", "BELOW_THE_FOLD", "CONFLICT", "DP", "EMBEDDED_VIDEO", "GAMES", "JACKASS", "PROFANITY", "UGC_FORUMS", "UGC_IMAGES", "UGC_SOCIAL", "UGC_VIDEOS", "SIRENS", "TRAGEDY", "VIDEO", "UNKNOWN"
}

type GenderCriterion struct {
	Id         int64  `xml:"id,omitempty"`
	GenderType string `xml:"genderType"` // GenderType:  "GENDER_MALE", "GENDER_FEMALE", "GENDER_UNDETERMINED"
}

type KeywordCriterion struct {
	Id        int64  `xml:"id,omitempty"`
	Text      string `xml:"text,omitempty"`      // Text: up to 80 characters and ten words
	MatchType string `xml:"matchType,omitempty"` // MatchType:  "EXACT", "PHRASE", "BROAD"
}

type LanguageCriterion struct {
	Id   int64  `xml:"id,omitempty"`
	Code string `xml:"code,omitempty"`
	Name string `xml:"name,omitempty"`
}

// LocationName:
// DisplayType:
// TargetingStatus: ACTIVE, OBSOLETE, PHASING_OUT
// ParentLocations:
type Location struct {
	Id              int64      `xml:"id,omitempty"`
	LocationName    string     `xml:"locationName,omitempty"`
	DisplayType     string     `xml:"displayType,omitempty"`
	TargetingStatus string     `xml:"targetingStatus,omitempty"`
	ParentLocations []Location `xml:"parentLocations,omitempty"`
}

// MobileAppCategoryId:
//   https://developers.google.com/adwords/api/docs/appendix/mobileappcategories
// DisplayName:
type MobileAppCategoryCriterion struct {
	Id                  int64  `xml:"id,omitempty"`
	MobileAppCategoryId int64  `xml:"mobileAppCategoryId"`
	DisplayName         string `xml:"displayName,omitempty"`
}

// AppId: "{platform}-{platform_native_id}"
// DisplayName:
type MobileApplicationCriterion struct {
	Id          int64  `xml:"id,omitempty"`
	AppId       string `xml:"appId"`
	DisplayName string `xml:"displayName,omitempty"`
}

// DeviceName:
// ManufacturerName:
// DeviceType:  DEVICE_TYPE_MOBILE, DEVICE_TYPE_TABLET
// OperatingSystemName:
type MobileDeviceCriterion struct {
	Id                  int64  `xml:"id,omitempty"`
	DeviceName          string `xml:"deviceName,omitempty"`
	ManufacturerName    string `xml:"manufacturerName,omitempty"`
	DeviceType          string `xml:"deviceType,omitempty"`
	OperatingSystemName string `xml:"operatingSystemName,omitempty"`
}

// Name:
// OsMajorVersion:
// OsMinorVersion:
// OperatorType: GREATER_THAN_EQUAL_TO, EQUAL_TO, UNKNOWN
type OperatingSystemVersionCriterion struct {
	Id             int64  `xml:"id,omitempty"`
	Name           string `xml:"name,omitempty"`
	OsMajorVersion int64  `xml:"osMajorVersion,omitempty"`
	OsMinorVersion int64  `xml:"osMinorVersion,omitempty"`
	OperatorType   string `xml:"operatorType,omitempty"`
}

// Url:
type PlacementCriterion struct {
	Id  int64  `xml:"id,omitempty"`
	Url string `xml:"url"`
}

// PlatformId:
//  Desktop	30000
//  HighEndMobile	30001
//  Tablet	30002
type PlatformCriterion struct {
	Id           int64  `xml:"id,omitempty"`
	PlatformName string `xml:"platformName,omitempty"`
}

// Argument:
// Operand: id, product_type, brand, adwords_grouping, condition, adwords_labels
type ProductCondition struct {
	Argument string `xml:"argument"`
	Operand  string `xml:"operand"`
}

type ProductCriterion struct {
	Id         int64              `xml:"id,omitempty"`
	Conditions []ProductCondition `xml:"conditions"`
	Text       string             `xml:"text,omitempty"`
}

type GeoPoint struct {
	Latitude  int64 `xml:"latitudeInMicroDegrees"`
	Longitude int64 `xml:"longitudeInMicroDegrees"`
}

type Address struct {
	StreetAddress  string `xml:"streetAddress"`
	StreetAddress2 string `xml:"streetAddress2"`
	CityName       string `xml:"cityName"`
	ProvinceCode   string `xml:"provinceCode"`
	ProvinceName   string `xml:"provinceName"`
	PostalCode     string `xml:"postalCode"`
	CountryCode    string `xml:"countryCode"`
}

// RadiusDistanceUnits: KILOMETERS, MILES
// RadiusUnits:
type ProximityCriterion struct {
	Id                  int64    `xml:"id,omitempty"`
	GeoPoint            GeoPoint `xml:"geoPoint"`
	RadiusDistanceUnits string   `xml:"radiusDistanceUnits"`
	RadiusInUnits       float64  `xml:"radiusInUnits"`
	Address             Address  `xml:"address"`
}

type UserInterestCriterion struct {
	Id   int64  `xml:"userInterestId,omitempty"`
	Name string `xml:"userInterestName"`
}

type UserListCriterion struct {
	Id                       int64  `xml:"id,omitempty"`
	UserListId               int64  `xml:"userListId"`
	UserListName             string `xml:"userListName"`
	UserListMembershipStatus string `xml:"userListMembershipStatus"`
}

type VerticalCriterion struct {
	Id       int64    `xml:"verticalId,omitempty"`
	ParentId int64    `xml:"verticalParentId"`
	Path     []string `xml:"path"`
}

type WebpageCondition struct {
	Operand  string `xml:"operand"`
	Argument string `xml:"argument"`
}

type WebpageParameter struct {
	CriterionName string             `xml:"criterionName"`
	Conditions    []WebpageCondition `xml:"conditions"`
}

type WebpageCriterion struct {
	Id               int64            `xml:"id,omitempty"`
	Parameter        WebpageParameter `xml:"parameter"`
	CriteriaCoverage float64          `xml:"criteriaCoverage"`
	CriteriaSamples  []string         `xml:"criteriaSamples"`
}

type Criterion interface{}

func criterionUnmarshalXML(dec *xml.Decoder, start xml.StartElement) (Criterion, error) {
	criterionType, err := findAttr(start.Attr, xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"})
	if err != nil {
		return nil, err
	}
	switch criterionType {
	case "AdSchedule":
		c := AdScheduleCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "AgeRange":
		c := AgeRangeCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Carrier":
		c := CarrierCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "ContentLabel":
		c := ContentLabelCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Gender":
		c := GenderCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Keyword":
		c := KeywordCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Language":
		c := LanguageCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Location":
		c := Location{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "MobileAppCategory":
		c := MobileAppCategoryCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "MobileApplication":
		c := MobileApplicationCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "MobileDevice":
		c := MobileDeviceCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "OperatingSystemVersion":
		c := OperatingSystemVersionCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Placement":
		c := PlacementCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Platform":
		c := PlatformCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Product":
		c := ProductCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Proximity":
		c := ProximityCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "CriterionUserInterest":
		c := UserInterestCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "CriterionUserList":
		c := UserListCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Vertical":
		c := VerticalCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Webpage":
		c := WebpageCriterion{}
		err := dec.DecodeElement(&c, &start)
		return c, err
	default:
		return nil, fmt.Errorf("unknown criterion type %#v", criterionType)
	}
}

func criterionMarshalXML(c Criterion, e *xml.Encoder) error {
	criterionType := ""
	switch t := c.(type) {
	case AdScheduleCriterion:
		criterionType = "AdSchedule"
	case AgeRangeCriterion:
		criterionType = "AgeRange"
	case CarrierCriterion:
		criterionType = "Carrier"
	case ContentLabelCriterion:
		criterionType = "ContentLabel"
	case GenderCriterion:
		criterionType = "Gender"
	case KeywordCriterion:
		criterionType = "Keyword"
	case LanguageCriterion:
		criterionType = "Language"
	case Location:
		criterionType = "Location"
	case MobileAppCategoryCriterion:
		criterionType = "MobileAppCategory"
	case MobileApplicationCriterion:
		criterionType = "MobileApplication"
	case MobileDeviceCriterion:
		criterionType = "MobileDevice"
	case OperatingSystemVersionCriterion:
		criterionType = "OperatingSystemVersion"
	case PlacementCriterion:
		criterionType = "Placement"
	case PlatformCriterion:
		criterionType = "Platform"
	case ProductCriterion:
		criterionType = "Product"
	case ProximityCriterion:
		criterionType = "Proximity"
	case UserInterestCriterion:
		criterionType = "CriterionUserInterest"
	case UserListCriterion:
		criterionType = "CriterionUserList"
	case VerticalCriterion:
		criterionType = "Vertical"
	case WebpageCriterion:
		criterionType = "Webpage"
	default:
		return fmt.Errorf("unknown criterion type %#v\n", t)
	}
	e.EncodeElement(&c, xml.StartElement{
		xml.Name{"", "criterion"},
		[]xml.Attr{
			xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, criterionType},
		},
	})
	return nil
}
