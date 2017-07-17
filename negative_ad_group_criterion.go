package gads

import (
	"encoding/xml"
	"fmt"
)

type NegativeAdGroupCriterion struct {
	AdGroupId    int64     `xml:"adGroupId"`
	CriterionUse string    `xml:"criterionUse"`
	Criterion    Criterion `xml:"criterion"`
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
			case "criterionUse":
				if err := dec.DecodeElement(&nagc.CriterionUse, &start); err != nil {
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
