package gads

import (
	"encoding/xml"
	"fmt"
)

type BiddableAdGroupCriterion struct {
	AdGroupId    int64     `xml:"adGroupId"`
	CriterionUse string    `xml:"criterionUse"`
	Criterion    Criterion `xml:"criterion"`

	// BiddableAdGroupCriterion
	UserStatus          string   `xml:"userStatus,omitempty"`
	SystemServingStatus string   `xml:"systemServingStatus,omitempty"`
	ApprovalStatus      string   `xml:"approvalStatus,omitempty"`
	DisapprovalReasons  []string `xml:"disapprovalReasons,omitempty"`
	DestinationUrl      string   `xml:"destinationUrl,omitempty"`

	// TODO add ExperimentData
	FirstPageCpc *Cpc `xml:"firstPageCpc,omitempty"`
	TopOfPageCpc *Cpc `xml:"topOfPageCpc,omitempty"`

	QualityInfo *QualityInfo `xml:"qualityInfo,omitempty"`

	BiddingStrategyConfiguration *BiddingStrategyConfiguration `xml:"biddingStrategyConfiguration,omitempty"`
	BidModifier                  int64                         `xml:"bidModifier,omitempty"`

	FinalUrls           []string         `xml:"finalUrls,omitempty"`
	FinalMobileUrls     []string         `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls        []string         `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate string           `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters CustomParameters `xml:"urlCustomParameters,omitempty"`
	Labels              []Label          `xml:"labels,omitempty"`
}

func (bagc BiddableAdGroupCriterion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(
		start.Attr,
		xml.Attr{
			xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"},
			"BiddableAdGroupCriterion",
		},
	)
	e.EncodeToken(start)
	e.EncodeElement(&bagc.AdGroupId, xml.StartElement{Name: xml.Name{baseUrl, "adGroupId"}})
	criterionMarshalXML(bagc.Criterion, e)
	if bagc.UserStatus != "" {
		e.EncodeElement(&bagc.UserStatus, xml.StartElement{Name: xml.Name{baseUrl, "userStatus"}})
	}
	if bagc.DestinationUrl != "" {
		e.EncodeElement(&bagc.DestinationUrl, xml.StartElement{Name: xml.Name{baseUrl, "destinationUrl"}})
	}
	e.EncodeElement(&bagc.BiddingStrategyConfiguration, xml.StartElement{Name: xml.Name{baseUrl, "biddingStrategyConfiguration"}})
	if bagc.BidModifier != 0 {
		e.EncodeElement(&bagc.BidModifier, xml.StartElement{Name: xml.Name{baseUrl, "bidModifier"}})
	}
	e.EncodeElement(&bagc.UrlCustomParameters, xml.StartElement{Name: xml.Name{baseUrl, "urlCustomParameters"}})
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
			case "criterionUse":
				if err := dec.DecodeElement(&bagc.CriterionUse, &start); err != nil {
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
			case "finalUrls":
				if err := dec.DecodeElement(&bagc.FinalUrls, &start); err != nil {
					return err
				}
			case "finalMobileUrls":
				if err := dec.DecodeElement(&bagc.FinalMobileUrls, &start); err != nil {
					return err
				}
			case "finalAppUrls":
				if err := dec.DecodeElement(&bagc.FinalAppUrls, &start); err != nil {
					return err
				}
			case "trackingUrlTemplate":
				if err := dec.DecodeElement(&bagc.TrackingUrlTemplate, &start); err != nil {
					return err
				}
			case "urlCustomParameters":
				if err := dec.DecodeElement(&bagc.UrlCustomParameters, &start); err != nil {
					return err
				}
			case "labels":
				if err := dec.DecodeElement(&bagc.Labels, &start); err != nil {
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
