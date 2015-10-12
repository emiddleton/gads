package gads

import (
	"encoding/xml"
	"fmt"
)

type AdGroupAds []AdGroupAd

func (agas *AdGroupAds) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	typeName := xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"}
	var adGroupId int64
	var status, approvalStatus string
	var disapprovalReasons []string
	var trademarkDisapproved bool
	var labels []Label
	var a Ad

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
					atyped := TextAd{}
					atyped.Type = typeName
					err = dec.DecodeElement(&atyped, &start)
					if err != nil {
						return err
					}
					a = atyped
				case "ImageAd":
					atyped := ImageAd{}
					atyped.Type = typeName
					err = dec.DecodeElement(&atyped, &start)
					if err != nil {
						return err
					}
					a = atyped
				case "TemplateAd":
					atyped := TemplateAd{}
					atyped.Type = typeName
					err = dec.DecodeElement(&atyped, &start)
					if err != nil {
						return err
					}
					a = atyped
				case "MobileAd":
					atyped := MobileAd{}
					atyped.Type = typeName
					err = dec.DecodeElement(&atyped, &start)
					if err != nil {
						return err
					}
					a = atyped
				case "DynamicSearchAd":
					atyped := DynamicSearchAd{}
					atyped.Type = typeName
					err = dec.DecodeElement(&atyped, &start)
					if err != nil {
						return err
					}
					a = atyped
				default:
					return fmt.Errorf("unknown Ad -> %#v", typeName)
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
			case "labels":
				err := dec.DecodeElement(&labels, &start)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("unknown AdGroupAd field -> %#v", tag)
			}

		}
	}

	aga := AdGroupAd{
		AdGroupId:            adGroupId,
		Ad:                   a,
		Status:               status,
		DisapprovalReasons:   disapprovalReasons,
		TrademarkDisapproved: trademarkDisapproved,
		Labels:               labels,
	}
	*agas = append(*agas, aga)
	return nil
}

func (agads AdGroupAds) GetAds() (ads []Ad) {
	for _, aga := range agads {
		ads = append(ads, aga.Ad)
	}
	return ads
}
