package gads

import (
	"encoding/xml"
	"fmt"
)

var (
	AD_GROUP_AD_SERVICE_URL = ServiceUrl{
		baseUrl,
		"AdGroupAdService",
	}
)

type adGroupAdService struct {
	Auth
}

type TextAd struct {
	AdGroupId            int64    `xml:"-"`
	Id                   int64    `xml:"id,omitempty"`
	Url                  string   `xml:"url"`
	DisplayUrl           string   `xml:"displayUrl"`
	DevicePreference     int64    `xml:"devicePreference,omitempty"`
	Headline             string   `xml:"headline"`
	Description1         string   `xml:"description1"`
	Description2         string   `xml:"description2"`
	Status               string   `xml:"-"`
	ApprovalStatus       string   `xml:"-"`
	DisapprovalReasons   []string `xml:"-"`
	TrademarkDisapproved bool     `xml:"-"`
}

type ImageAd struct {
	AdGroupId            int64    `xml:"-"`
	Id                   int64    `xml:"id,omitempty"`
	Url                  string   `xml:"url"`
	DisplayUrl           string   `xml:"displayUrl"`
	DevicePreference     int64    `xml:"devicePreference,omitempty"`
	ImageId              int64    `xml:"imageId"`
	Name                 string   `xml:"name"`
	AdToCopyImageFrom    int64    `xml:"adToCopyImageFrom"`
	Status               string   `xml:"-"`
	ApprovalStatus       string   `xml:"-"`
	DisapprovalReasons   []string `xml:"-"`
	TrademarkDisapproved bool     `xml:"-"`
}

type MobileAd struct {
	AdGroupId            int64    `xml:"-"`
	Id                   int64    `xml:"id,omitempty"`
	Url                  string   `xml:"url"`
	DisplayUrl           string   `xml:"displayUrl"`
	DevicePreference     int64    `xml:"devicePreference,omitempty"`
	Headline             string   `xml:"headline"`
	Description          string   `xml:"description"`
	MarkupLanguages      []string `xml:"markupLanguages"`
	MobileCarriers       []string `xml:"mobileCarriers"`
	BusinessName         string   `xml:"businessName"`
	CountryCode          string   `xml:"countryCode"`
	PhoneNumber          string   `xml:"phoneNumber"`
	Status               string   `xml:"-"`
	ApprovalStatus       string   `xml:"-"`
	DisapprovalReasons   []string `xml:"-"`
	TrademarkDisapproved bool     `xml:"-"`
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

type Dimensions struct {
	Width  int64 `xml:"width"`
	Height int64 `xml:"height"`
}

type TemplateAd struct {
	AdGroupId            int64             `xml:"-"`
	Id                   int64             `xml:"id,omitempty"`
	Url                  string            `xml:"url"`
	DisplayUrl           string            `xml:"displayUrl"`
	DevicePreference     int64             `xml:"devicePreference,omitempty"`
	TemplateId           int64             `xml:"templateId"`
	AdUnionId            int64             `xml:"adUnionId"`
	TemplateElements     []TemplateElement `xml:"templateElements"`
	Dimensions           Dimensions        `xml:"dimensions"`
	Name                 string            `xml:"name"`
	Duration             int64             `xml:"duration"`
	originAdId           *int64            `xml:"originAdId"`
	Status               string            `xml:"-"`
	ApprovalStatus       string            `xml:"-"`
	DisapprovalReasons   []string          `xml:"-"`
	TrademarkDisapproved bool              `xml:"-"`
}

type AdGroupAds []interface{}

func (a1 AdGroupAds) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	a := a1[0]
	e.EncodeToken(start)
	switch a.(type) {
	case TextAd:
		ad := a.(TextAd)
		e.EncodeElement(ad.AdGroupId, xml.StartElement{Name: xml.Name{"", "adGroupId"}})
		e.EncodeElement(ad, xml.StartElement{
			xml.Name{"", "ad"},
			[]xml.Attr{
				xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, "TextAd"},
			},
		})
		e.EncodeElement(ad.Status, xml.StartElement{Name: xml.Name{"", "status"}})
	}
	e.EncodeToken(start.End())
	return nil
}

func (aga *AdGroupAds) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	typeName := xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"}
	var adGroupId int64
	var status, approvalStatus string
	var disapprovalReasons []string
	var trademarkDisapproved bool
	var ad interface{}
	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		if err != nil {
			return err
		}
		switch start := token.(type) {
		case xml.StartElement:
			tag := start.Name.Local
			switch tag {
			case "adGroupId":
				err := dec.DecodeElement(&adGroupId, &start)
				if err != nil {
					return err
				}
			case "ad":
				typeName, err := findAttr(start.Attr, typeName)
				if err != nil {
					return err
				}
				switch typeName {
				case "TextAd":
					a := TextAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "ImageAd":
					a := ImageAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				case "TemplateAd":
					a := TemplateAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
					ad = a
				default:
					return fmt.Errorf("unknown AdGroupCriterion -> %#v", start)
				}
			case "status":
				err := dec.DecodeElement(&status, &start)
				if err != nil {
					return err
				}
			case "approvalStatus":
				err := dec.DecodeElement(&approvalStatus, &start)
				if err != nil {
					return err
				}
			case "disapprovalReasons":
				err := dec.DecodeElement(&disapprovalReasons, &start)
				if err != nil {
					return err
				}
			case "trademarkDisapproved":
				err := dec.DecodeElement(&trademarkDisapproved, &start)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("unknown AdGroupAd field -> %#v", tag)
			}

		}
	}
	switch a := ad.(type) {
	case TextAd:
		a.Status = status
		a.ApprovalStatus = approvalStatus
		a.DisapprovalReasons = disapprovalReasons
		a.TrademarkDisapproved = trademarkDisapproved
		*aga = append(*aga, a)
	case ImageAd:
		a.Status = status
		a.ApprovalStatus = approvalStatus
		a.DisapprovalReasons = disapprovalReasons
		a.TrademarkDisapproved = trademarkDisapproved
		*aga = append(*aga, a)
	case TemplateAd:
		a.Status = status
		a.ApprovalStatus = approvalStatus
		a.DisapprovalReasons = disapprovalReasons
		a.TrademarkDisapproved = trademarkDisapproved
		*aga = append(*aga, a)
	}
	return nil
}

type AdGroupAdOperations map[string]AdGroupAds

func NewAdGroupAdService(auth Auth) *adGroupAdService {
	return &adGroupAdService{Auth: auth}
}

func NewTextAd(adGroupId int64, url, displayUrl, headline, description1, description2, status string) TextAd {
	return TextAd{
		AdGroupId:    adGroupId,
		Url:          url,
		DisplayUrl:   displayUrl,
		Headline:     headline,
		Description1: description1,
		Description2: description2,
		Status:       status,
	}
}

func (s adGroupAdService) Get(selector Selector) (adGroupAds AdGroupAds, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.Request(
		AD_GROUP_AD_SERVICE_URL,
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
		return adGroupAds, err
	}
	getResp := struct {
		Size       int64      `xml:"rval>totalNumEntries"`
		AdGroupAds AdGroupAds `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroupAds, err
	}
	return getResp.AdGroupAds, err
}

func (s *adGroupAdService) Mutate(adGroupAdOperations AdGroupAdOperations) (adGroupAds AdGroupAds, err error) {
	type adGroupAdOperation struct {
		Action    string     `xml:"operator"`
		AdGroupAd AdGroupAds `xml:"operand"`
	}
	operations := []adGroupAdOperation{}
	for action, adGroupAds := range adGroupAdOperations {
		for _, adGroupAd := range adGroupAds {
			operations = append(operations,
				adGroupAdOperation{
					Action:    action,
					AdGroupAd: []interface{}{adGroupAd},
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
	respBody, err := s.Auth.Request(AD_GROUP_AD_SERVICE_URL, "mutate", mutation)
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
