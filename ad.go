package gads

import "encoding/xml"

// Ad type as defined
// https://developers.google.com/adwords/api/docs/reference/v201506/AdGroupAdService.Ad
type CommonAd struct {
	Type                string            `xml:"xsi:type,attr,omitempty"`
	ID                  int64             `xml:"id,omitempty"`
	URL                 string            `xml:"url,omitempty"`
	DisplayURL          string            `xml:"displayUrl,omitempty"`
	FinalURLs           []string          `xml:"finalUrls,omitempty"`
	FinalMobileURLs     []string          `xml:"finalMobileUrls,omitempty"`
	FinalAppURLs        []AppUrl          `xml:"finalAppUrls,omitempty"`
	TrackingURLTemplate *string           `xml:"trackingUrlTemplate,omitempty"`
	URLCustomParameters *CustomParameters `xml:"urlCustomParameters,omitempty"`
	DevicePreference    int64             `xml:"devicePreference,omitempty"`
}

type Ad interface {
	GetID() int64
	GetURL() string
	GetTrackingURLTemplate() *string
	GetFinalURLs() []string

	CloneForTemplate([]string, *string) Ad
}

func (c CommonAd) GetID() int64 {
	return c.ID
}

func (c CommonAd) UnsetID() {
}

func (c CommonAd) GetURL() string {
	return c.URL
}

func (c CommonAd) GetTrackingURLTemplate() *string {
	return c.TrackingURLTemplate
}

func (c CommonAd) GetFinalURLs() []string {
	return c.FinalURLs
}

func (c CommonAd) CloneForTemplate(finalURLs []string, trackingURLTemplate *string) Ad {
	c.ID = 0 // value used by go for omitempty
	c.FinalURLs = finalURLs
	c.TrackingURLTemplate = trackingURLTemplate
	return c
}

func (c TextAd) CloneForTemplate(finalURLs []string, trackingURLTemplate *string) Ad {
	c.ID = 0 // value used by go for omitempty
	c.FinalURLs = finalURLs
	c.TrackingURLTemplate = trackingURLTemplate

	return c
}

func (c ImageAd) CloneForTemplate(finalURLs []string, trackingURLTemplate *string) Ad {
	c.ID = 0 // value used by go for omitempty
	c.FinalURLs = finalURLs
	c.TrackingURLTemplate = trackingURLTemplate
	return c
}

func (c TemplateAd) CloneForTemplate(finalURLs []string, trackingURLTemplate *string) Ad {
	c.ID = 0 // value used by go for omitempty
	c.FinalURLs = finalURLs
	c.TrackingURLTemplate = trackingURLTemplate
	return c
}

func (c MobileAd) CloneForTemplate(finalURLs []string, trackingURLTemplate *string) Ad {
	c.ID = 0 // value used by go for omitempty
	c.FinalURLs = finalURLs
	c.TrackingURLTemplate = trackingURLTemplate
	return c
}

func (c DynamicSearchAd) CloneForTemplate(finalURLs []string, trackingURLTemplate *string) Ad {
	c.ID = 0 // value used by go for omitempty
	c.FinalURLs = finalURLs
	c.TrackingURLTemplate = trackingURLTemplate
	return c
}

type TextAd struct {
	CommonAd
	Headline     string `xml:"headline"`
	Description1 string `xml:"description1"`
	Description2 string `xml:"description2"`
}

type DynamicSearchAd struct {
	CommonAd
	Description1 string `xml:"description1"`
	Description2 string `xml:"description2"`
}

type ImageAd struct {
	CommonAd
	Image             int64  `xml:"imageId"` //TODO should actually be Image object, not just an int
	Name              string `xml:"name"`
	AdToCopyImageFrom int64  `xml:"adToCopyImageFrom"`
}

type MobileAd struct {
	CommonAd
	Headline        string   `xml:"headline"`
	Description     string   `xml:"description"`
	MarkupLanguages []string `xml:"markupLanguages"`
	MobileCarriers  []string `xml:"mobileCarriers"`
	BusinessName    string   `xml:"businessName"`
	CountryCode     string   `xml:"countryCode"`
	PhoneNumber     string   `xml:"phoneNumber"`
}

type TemplateAd struct {
	CommonAd
	TemplateId       int64             `xml:"templateId"`
	AdUnionId        int64             `xml:"adUnionId"`
	TemplateElements []TemplateElement `xml:"templateElements"`
	Dimensions       []Dimensions      `xml:"dimensions"`
	Name             string            `xml:"name"`
	Duration         int64             `xml:"duration"`
	originAdId       *int64            `xml:"originAdId"`
}

func (i ImageAd) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return ERROR_NOT_YET_IMPLEMENTED
}

func (t TemplateAd) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return ERROR_NOT_YET_IMPLEMENTED
}
