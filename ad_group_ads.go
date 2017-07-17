package gads

import (
	"encoding/xml"
	"fmt"
)

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
		e.EncodeElement(ad.Labels, xml.StartElement{Name: xml.Name{"", "labels"}})
	case ExpandedTextAd:
		ad := a.(ExpandedTextAd)
		e.EncodeElement(ad.AdGroupId, xml.StartElement{Name: xml.Name{"", "adGroupId"}})
		e.EncodeElement(ad, xml.StartElement{
			xml.Name{"", "ad"},
			[]xml.Attr{
				xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, "ExpandedTextAd"},
			},
		})
		e.EncodeElement(ad.Status, xml.StartElement{Name: xml.Name{"", "status"}})
		e.EncodeElement(ad.Labels, xml.StartElement{Name: xml.Name{"", "labels"}})
	case ImageAd:
		return ERROR_NOT_YET_IMPLEMENTED
	case TemplateAd:
		return ERROR_NOT_YET_IMPLEMENTED
	default:
		return fmt.Errorf("unknown Ad type -> %#v", start)
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
	var labels []Label
	var experimentData *AdGroupExperimentData
	var trademarks []string
	var baseCampaignId *int64
	var baseAdGroupId *int64
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
				case "ExpandedTextAd":
					a := ExpandedTextAd{AdGroupId: adGroupId}
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
				case "DynamicSearchAd":
					a := DynamicSearchAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
				case "ProductAd":
					a := ProductAd{AdGroupId: adGroupId}
					err := dec.DecodeElement(&a, &start)
					if err != nil {
						return err
					}
				default:
					return fmt.Errorf("unknown AdGroupCriterion -> %#v", start)
				}
			case "experimentData":
				err := dec.DecodeElement(&experimentData, &start)
				if err != nil {
					return err
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
			case "trademarks":
				err := dec.DecodeElement(&trademarks, &start)
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
			case "labels":
				err := dec.DecodeElement(&labels, &start)
				if err != nil {
					return err
				}
			case "baseCampaignId":
				err := dec.DecodeElement(&baseCampaignId, &start)
				if err != nil {
					return err
				}
			case "baseAdGroupId":
				err := dec.DecodeElement(&baseAdGroupId, &start)
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
	case ExpandedTextAd:
		a.ExperimentData = experimentData
		a.Status = status
		a.ApprovalStatus = approvalStatus
		a.Trademarks = trademarks
		a.DisapprovalReasons = disapprovalReasons
		a.TrademarkDisapproved = trademarkDisapproved
		a.Labels = labels
		a.BaseCampaignId = baseCampaignId
		a.BaseAdGroupId = baseAdGroupId
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
	case DynamicSearchAd:
		a.Status = status
		a.ApprovalStatus = approvalStatus
		a.DisapprovalReasons = disapprovalReasons
		a.TrademarkDisapproved = trademarkDisapproved
		*aga = append(*aga, a)
	case ProductAd:
		a.Status = status
		a.ApprovalStatus = approvalStatus
		a.DisapprovalReasons = disapprovalReasons
		a.TrademarkDisapproved = trademarkDisapproved
		*aga = append(*aga, a)
	}
	return nil
}
