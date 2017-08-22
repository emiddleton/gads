package gads

import (
	"encoding/xml"
	"fmt"
)

// https://developers.google.com/adwords/api/docs/reference/v201708/AdGroupExtensionSettingService.ExtensionSetting
// A setting specifying when and which extensions should serve at a given level (customer, campaign, or ad group).
type ExtensionSetting struct {
	PlatformRestrictions ExtensionSettingPlatform `xml:"platformRestrictions,omitempty"`

	Extensions Extension `xml:"https://adwords.google.com/api/adwords/cm/v201708 extensions,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201708/AdGroupExtensionSettingService.ExtensionSetting.Platform
// Different levels of platform restrictions
// DESKTOP, MOBILE, NONE
type ExtensionSettingPlatform string

type Extension interface{}

// https://developers.google.com/adwords/api/docs/reference/v201708/AdGroupExtensionSettingService.ExtensionFeedItem
// Contains base extension feed item data for an extension in an extension feed managed by AdWords.
type ExtensionFeedItem struct {
	XMLName xml.Name `json:"-" xml:"extensions"`

	FeedId                  int64                      `xml:"https://adwords.google.com/api/adwords/cm/v201708 feedId,omitempty"`
	FeedItemId              int64                      `xml:"https://adwords.google.com/api/adwords/cm/v201708 feedItemId,omitempty"`
	Status                  *FeedItemStatus            `xml:"https://adwords.google.com/api/adwords/cm/v201708 status,omitempty"`
	FeedType                *FeedType                  `xml:"https://adwords.google.com/api/adwords/cm/v201708 feedType,omitempty"`
	StartTime               string                     `xml:"https://adwords.google.com/api/adwords/cm/v201708 startTime,omitempty"` //  special value "00000101 000000" may be used to clear an existing start time.
	EndTime                 string                     `xml:"https://adwords.google.com/api/adwords/cm/v201708 endTime,omitempty"`   //  special value "00000101 000000" may be used to clear an existing end time.
	DevicePreference        *FeedItemDevicePreference  `xml:"https://adwords.google.com/api/adwords/cm/v201708 devicePreference,omitempty"`
	Scheduling              *FeedItemScheduling        `xml:"https://adwords.google.com/api/adwords/cm/v201708 scheduling,omitempty"`
	CampaignTargeting       *FeedItemCampaignTargeting `xml:"https://adwords.google.com/api/adwords/cm/v201708 campaignTargeting,omitempty"`
	AdGroupTargeting        *FeedItemAdGroupTargeting  `xml:"https://adwords.google.com/api/adwords/cm/v201708 adGroupTargeting,omitempty"`
	KeywordTargeting        *Keyword                   `xml:"https://adwords.google.com/api/adwords/cm/v201708 keywordTargeting,omitempty"`
	GeoTargeting            *Location                  `xml:"https://adwords.google.com/api/adwords/cm/v201708 geoTargeting,omitempty"`
	GeoTargetingRestriction *FeedItemGeoRestriction    `xml:"https://adwords.google.com/api/adwords/cm/v201708 geoTargetingRestriction,omitempty"`
	PolicyData              *[]FeedItemPolicyData      `xml:"https://adwords.google.com/api/adwords/cm/v201708 policyData,omitempty"`

	ExtensionFeedItemType string `xml:"https://adwords.google.com/api/adwords/cm/v201708 ExtensionFeedItem.Type,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201708/AdGroupExtensionSettingService.CallFeedItem
// Represents a Call extension.
type CallFeedItem struct {
	ExtensionFeedItem

	CallPhoneNumber               string             `xml:"https://adwords.google.com/api/adwords/cm/v201708 callPhoneNumber,omitempty"`
	CallCountryCode               string             `xml:"https://adwords.google.com/api/adwords/cm/v201708 callCountryCode,omitempty"`
	CallTracking                  bool               `xml:"https://adwords.google.com/api/adwords/cm/v201708 callTracking,omitempty"`
	CallConversionType            CallConversionType `xml:"https://adwords.google.com/api/adwords/cm/v201708 callConversionType,omitempty"`
	DisableCallConversionTracking bool               `xml:"https://adwords.google.com/api/adwords/cm/v201708 disableCallConversionTracking,omitempty"`
}

func extensionsUnmarshalXML(dec *xml.Decoder, start xml.StartElement) (ext interface{}, err error) {
	extensionsType, err := findAttr(start.Attr, xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"})
	if err != nil {
		return
	}
	switch extensionsType {
	case "CallFeedItem":
		c := CallFeedItem{}
		err = dec.DecodeElement(&c, &start)
		ext = c
	default:
		err = fmt.Errorf("unknown Extensions type %#v", extensionsType)
	}
	return
}

func (s ExtensionSetting) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if s.PlatformRestrictions != "NONE" {
		e.EncodeElement(&s.PlatformRestrictions, xml.StartElement{Name: xml.Name{
			"https://adwords.google.com/api/adwords/cm/v201708",
			"platformRestrictions"}})
	}
	switch extType := s.Extensions.(type) {
	case []CallFeedItem:
		e.EncodeElement(s.Extensions.([]CallFeedItem), xml.StartElement{
			xml.Name{baseUrl, "extensions"},
			[]xml.Attr{
				xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, "CallFeedItem"},
			},
		})
	default:
		return fmt.Errorf("unknown extension type %#v\n", extType)

	}

	e.EncodeToken(start.End())
	return nil
}

func (s *ExtensionSetting) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) (err error) {
	s.Extensions = []interface{}{}

	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		if err != nil {
			return err
		}
		switch start := token.(type) {
		case xml.StartElement:
			switch start.Name.Local {
			case "platformRestrictions":
				if err := dec.DecodeElement(&s.PlatformRestrictions, &start); err != nil {
					return err
				}
			case "extensions":
				extension, err := extensionsUnmarshalXML(dec, start)
				if err != nil {
					return err
				}
				s.Extensions = append(s.Extensions.([]interface{}), extension)
			}
		}
	}
	return nil
}
