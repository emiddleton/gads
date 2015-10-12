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
	Type        string `xml:"xsi:type,attr,omitempty"`
	Id          int64  `xml:"id,omitempty"`
	DayOfWeek   string `xml:"dayOfWeek"`
	StartHour   string `xml:"startHour"`
	StartMinute string `xml:"startMinute"`
	EndHour     string `xml:"endHour"`
	EndMinute   string `xml:"endMinute"`
}

func (c AdScheduleCriterion) GetID() int64 {
	return c.Id
}

// AgeRangeType: AGE_RANGE_18_24, AGE_RANGE_25_34, AGE_RANGE_35_44, AGE_RANGE_45_54, AGE_RANGE_55_64, AGE_RANGE_65_UP, AGE_RANGE_UNDETERMINED, UNKNOWN
type AgeRangeCriterion struct {
	Type         string `xml:"xsi:type,attr,omitempty"`
	Id           int64  `xml:"id,omitempty"`
	AgeRangeType string `xml:"ageRangeType"`
}

func (c AgeRangeCriterion) GetID() int64 {
	return c.Id
}

type CarrierCriterion struct {
	Type        string `xml:"xsi:type,attr,omitempty"`
	Id          int64  `xml:"id,omitempty"`
	Name        string `xml:"name,emitempty"`
	CountryCode string `xml:"countryCode,emitempty"`
}

func (c CarrierCriterion) GetID() int64 {
	return c.Id
}

type ContentLabelCriterion struct {
	Type             string `xml:"xsi:type,attr,omitempty"`
	Id               int64  `xml:"id,omitempty"`
	ContentLabelType string `xml:"contentLabelType"` // ContentLabelType: "ADULTISH", "AFE", "BELOW_THE_FOLD", "CONFLICT", "DP", "EMBEDDED_VIDEO", "GAMES", "JACKASS", "PROFANITY", "UGC_FORUMS", "UGC_IMAGES", "UGC_SOCIAL", "UGC_VIDEOS", "SIRENS", "TRAGEDY", "VIDEO", "UNKNOWN"
}

func (c ContentLabelCriterion) GetID() int64 {
	return c.Id
}

type GenderCriterion struct {
	Type       string `xml:"xsi:type,attr,omitempty"`
	Id         int64  `xml:"id,omitempty"`
	GenderType string `xml:"genderType"` // GenderType:  "GENDER_MALE", "GENDER_FEMALE", "GENDER_UNDETERMINED"
}

func (c GenderCriterion) GetID() int64 {
	return c.Id
}

type KeywordCriterion struct {
	Type      string `xml:"xsi:type,attr,omitempty"`
	Id        int64  `xml:"id,omitempty"`
	Text      string `xml:"text,omitempty"`      // Text: up to 80 characters and ten words
	MatchType string `xml:"matchType,omitempty"` // MatchType:  "EXACT", "PHRASE", "BROAD"
}

func (c KeywordCriterion) GetID() int64 {
	return c.Id
}

type LanguageCriterion struct {
	Type string `xml:"xsi:type,attr,omitempty"`
	Id   int64  `xml:"id,omitempty"`
	Code string `xml:"code,omitempty"`
	Name string `xml:"name,omitempty"`
}

func (c LanguageCriterion) GetID() int64 {
	return c.Id
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

func (c Location) GetID() int64 {
	return c.Id
}

// MobileAppCategoryId:
//   https://developers.google.com/adwords/api/docs/appendix/mobileappcategories
// DisplayName:
type MobileAppCategoryCriterion struct {
	Type                string `xml:"xsi:type,attr,omitempty"`
	Id                  int64  `xml:"id,omitempty"`
	MobileAppCategoryId int64  `xml:"mobileAppCategoryId"`
	DisplayName         string `xml:"displayName,omitempty"`
}

func (c MobileAppCategoryCriterion) GetID() int64 {
	return c.Id
}

// AppId: "{platform}-{platform_native_id}"
// DisplayName:
type MobileApplicationCriterion struct {
	Type        string `xml:"xsi:type,attr,omitempty"`
	Id          int64  `xml:"id,omitempty"`
	AppId       string `xml:"appId"`
	DisplayName string `xml:"displayName,omitempty"`
}

func (c MobileApplicationCriterion) GetID() int64 {
	return c.Id
}

// DeviceName:
// ManufacturerName:
// DeviceType:  DEVICE_TYPE_MOBILE, DEVICE_TYPE_TABLET
// OperatingSystemName:
type MobileDeviceCriterion struct {
	Type                string `xml:"xsi:type,attr,omitempty"`
	Id                  int64  `xml:"id,omitempty"`
	DeviceName          string `xml:"deviceName,omitempty"`
	ManufacturerName    string `xml:"manufacturerName,omitempty"`
	DeviceType          string `xml:"deviceType,omitempty"`
	OperatingSystemName string `xml:"operatingSystemName,omitempty"`
}

func (c MobileDeviceCriterion) GetID() int64 {
	return c.Id
}

// Name:
// OsMajorVersion:
// OsMinorVersion:
// OperatorType: GREATER_THAN_EQUAL_TO, EQUAL_TO, UNKNOWN
type OperatingSystemVersionCriterion struct {
	Type           string `xml:"xsi:type,attr,omitempty"`
	Id             int64  `xml:"id,omitempty"`
	Name           string `xml:"name,omitempty"`
	OsMajorVersion int64  `xml:"osMajorVersion,omitempty"`
	OsMinorVersion int64  `xml:"osMinorVersion,omitempty"`
	OperatorType   string `xml:"operatorType,omitempty"`
}

func (c OperatingSystemVersionCriterion) GetID() int64 {
	return c.Id
}

// Url:
type PlacementCriterion struct {
	Type string `xml:"xsi:type,attr,omitempty"`
	Id   int64  `xml:"id,omitempty"`
	Url  string `xml:"url"`
}

func (c PlacementCriterion) GetID() int64 {
	return c.Id
}

// PlatformId:
//  Desktop	30000
//  HighEndMobile	30001
//  Tablet	30002
type PlatformCriterion struct {
	Type         string `xml:"xsi:type,attr,omitempty"`
	Id           int64  `xml:"id,omitempty"`
	PlatformName string `xml:"platformName,omitempty"`
}

func (c PlatformCriterion) GetID() int64 {
	return c.Id
}

// Argument:
// Operand: id, product_type, brand, adwords_grouping, condition, adwords_labels
type ProductCondition struct {
	Argument string `xml:"argument"`
	Operand  string `xml:"operand"`
}

func (c ProductCriterion) GetID() int64 {
	return c.Id
}

type ProductCriterion struct {
	Type       string             `xml:"xsi:type,attr,omitempty"`
	Id         int64              `xml:"id,omitempty"`
	Conditions []ProductCondition `xml:"conditions,omitempty"`
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
	Type                string   `xml:"xsi:type,attr,omitempty"`
	Id                  int64    `xml:"id,omitempty"`
	GeoPoint            GeoPoint `xml:"geoPoint"`
	RadiusDistanceUnits string   `xml:"radiusDistanceUnits"`
	RadiusInUnits       float64  `xml:"radiusInUnits"`
	Address             Address  `xml:"address"`
}

func (c ProximityCriterion) GetID() int64 {
	return c.Id
}

type UserInterestCriterion struct {
	Type string `xml:"xsi:type,attr,omitempty"`
	Id   int64  `xml:"userInterestId,omitempty"`
	Name string `xml:"userInterestName"`
}

func (c UserInterestCriterion) GetID() int64 {
	return c.Id
}

type UserListCriterion struct {
	Type                     string `xml:"xsi:type,attr,omitempty"`
	Id                       int64  `xml:"id,omitempty"`
	UserListId               int64  `xml:"userListId"`
	UserListName             string `xml:"userListName"`
	UserListMembershipStatus string `xml:"userListMembershipStatus"`
}

func (c UserListCriterion) GetID() int64 {
	return c.Id
}

type VerticalCriterion struct {
	Type     string   `xml:"xsi:type,attr,omitempty"`
	Id       int64    `xml:"verticalId,omitempty"`
	ParentId int64    `xml:"verticalParentId"`
	Path     []string `xml:"path"`
}

func (c VerticalCriterion) GetID() int64 {
	return c.Id
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
	Type             string           `xml:"xsi:type,attr,omitempty"`
	Id               int64            `xml:"id,omitempty"`
	Parameter        WebpageParameter `xml:"parameter"`
	CriteriaCoverage float64          `xml:"criteriaCoverage"`
	CriteriaSamples  []string         `xml:"criteriaSamples"`
}

func (c WebpageCriterion) GetID() int64 {
	return c.Id
}

type Criterion interface {
	GetID() int64
}

func criterionUnmarshalXML(dec *xml.Decoder, start xml.StartElement) (Criterion, error) {
	criterionType, err := findAttr(start.Attr, xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"})
	if err != nil {
		return nil, err
	}
	switch criterionType {
	case "AdSchedule":
		c := AdScheduleCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "AgeRange":
		c := AgeRangeCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Carrier":
		c := CarrierCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "ContentLabel":
		c := ContentLabelCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Gender":
		c := GenderCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Keyword":
		c := KeywordCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Language":
		c := LanguageCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "MobileAppCategory":
		c := MobileAppCategoryCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "MobileApplication":
		c := MobileApplicationCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "MobileDevice":
		c := MobileDeviceCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "OperatingSystemVersion":
		c := OperatingSystemVersionCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Placement":
		c := PlacementCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Platform":
		c := PlatformCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Product":
		c := ProductCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Proximity":
		c := ProximityCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "CriterionUserInterest":
		c := UserInterestCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "CriterionUserList":
		c := UserListCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Vertical":
		c := VerticalCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	case "Webpage":
		c := WebpageCriterion{}
		c.Type = criterionType
		err := dec.DecodeElement(&c, &start)
		return c, err
	default:
		if StrictMode {
			return nil, fmt.Errorf("unknown criterion type %#v", criterionType)
		} else {
			return nil, nil
		}
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
