package gads

import (
	"encoding/xml"
	"fmt"
)

// TargetingIdeaService Use this service to generate new keyword ideas based on the parameters specified in the selector. See the TargetingIdeaSelector documentation for more details.
// You can also use this service to retrieve statistics for existing keyword ideas by setting the selector's requestType to RequestType.STATS and passing in the appropriate search parameters.
type TargetingIdeaService struct {
	Auth
}

// NewTargetingIdeaService returns a new TargetingIdeaService
func NewTargetingIdeaService(auth *Auth) *TargetingIdeaService {
	return &TargetingIdeaService{Auth: *auth}
}

// TargetingIdeaSelector A descriptor for finding TargetingIdeas that match the specified criteria.
// https://developers.google.com/adwords/api/docs/reference/v201708/TargetingIdeaService.TargetingIdeaSelector
type TargetingIdeaSelector struct {
	SearchParameters        []SearchParameter `xml:"searchParameters"`
	IdeaType                string            `xml:"ideaType"`
	RequestType             string            `xml:"requestType"`
	RequestedAttributeTypes []string          `xml:"requestedAttributeTypes"`
	Paging                  Paging            `xml:"paging"`
	LocaleCode              string            `xml:"localeCode,omitempty"`
	CurrencyCode            string            `xml:"currencyCode,omitempty"`
}

type SearchParameter interface{}

type CategoryProductsAndServicesSearchParameter struct {
	CategoryID int `xml:"categoryId"`
}

type CompetitionSearchParameter struct {
	Levels []string `xml:"levels"`
}

type IdeaTextFilterSearchParameter struct {
	Included []string `xml:"included"`
	Excluded []string `xml:"excluded"`
}

type IncludeAdultContentSearchParameter struct{}

type LanguageSearchParameter struct {
	Languages []LanguageCriterion `xml:"languages"`
}

type LocationSearchParameter struct {
	Locations []Location `xml:"locations"`
}

type NetworkSearchParameter struct {
	NetworkSetting NetworkSetting `xml:"networkSetting"`
}

type RelatedToQuerySearchParameter struct {
	Queries []string `xml:"queries"`
}

type RelatedToUrlSearchParameter struct {
	Urls           []string `xml:"urls"`
	IncludeSubUrls bool     `xml:"includeSubUrls"`
}

type SearchVolumeSearchParameter struct {
	Minimum int `xml:"operation>minimum"`
	Maximum int `xml:"operation>maximum"`
}

type SeedAdGroupIdSearchParameter struct {
	AdGroupID int64 `xml:"adGroupId"`
}

type TargetingIdeas struct {
	TargetingIdea []TargetingIdea `xml:"data"`
}

type TargetingIdea struct {
	Key   string    `xml:"key"`
	Value Attribute `xml:"value"`
}

type BooleanAttribute struct {
	Value bool `xml:"value"`
}

type DoubleAttribute struct {
	Value float64 `xml:"value"`
}

type IdeaTypeAttribute struct {
	Value string `xml:"value"`
}

type IntegerSetAttribute struct {
	Value []int `xml:"value"`
}

type LongAttribute struct {
	Value int64 `xml:"value"`
}

type MoneyAttribute struct {
	Value int64 `xml:"value>microAmount"`
}

type MonthlySearchVolumeAttribute struct {
	Value []MonthlySearchVolume `xml:"value"`
}

type StringAttribute struct {
	Value string `xml:"value"`
}

type WebpageDescriptorAttribute struct {
	Value WebpageDescriptor `xml:"value"`
}

type WebpageDescriptor struct {
	Url   string `xml:"url"`
	Title string `xml:"title"`
}

type MonthlySearchVolume struct {
	Year  int   `xml:"year"`
	Month int   `xml:"month"`
	Count int64 `xml:"count"`
}

type Attribute interface{}

func (ti *TargetingIdea) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) (err error) {
	for token, err := dec.Token(); err == nil; token, err = dec.Token() {
		if err != nil {
			return err
		}
		switch start := token.(type) {
		case xml.StartElement:
			tag := start.Name.Local
			switch tag {
			case "key":
				var key string
				err := dec.DecodeElement(&key, &start)
				if err != nil {
					return err
				}
				ti.Key = key
			case "value":
				value, err := attributeUnmarshalXML(dec, start)
				// because there are multiple "value" elements we only want the ones that have an xsi attr
				if err != nil {
					break
				}
				ti.Value = value
			}
		}
	}
	return nil
}

func attributeUnmarshalXML(dec *xml.Decoder, start xml.StartElement) (Attribute, error) {
	attributeType, err := findAttr(start.Attr, xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"})
	if err != nil {
		return nil, err
	}
	switch attributeType {
	case "BooleanAttribute":
		ba := BooleanAttribute{}
		err := dec.DecodeElement(&ba, &start)
		return ba.Value, err
	case "DoubleAttribute":
		da := DoubleAttribute{}
		err := dec.DecodeElement(&da, &start)
		return da.Value, err
	case "IdeaTypeAttribute":
		ita := IdeaTypeAttribute{}
		err := dec.DecodeElement(&ita, &start)
		return ita.Value, err
	case "IntegerSetAttribute":
		isa := IntegerSetAttribute{}
		err := dec.DecodeElement(&isa, &start)
		return isa.Value, err
	case "LongAttribute":
		la := LongAttribute{}
		err := dec.DecodeElement(&la, &start)
		return la.Value, err
	case "MoneyAttribute":
		ma := MoneyAttribute{}
		err := dec.DecodeElement(&ma, &start)
		return ma.Value, err
	case "MonthlySearchVolumeAttribute":
		msva := MonthlySearchVolumeAttribute{}
		err := dec.DecodeElement(&msva, &start)
		return msva.Value, err
	case "StringAttribute":
		sa := StringAttribute{}
		err := dec.DecodeElement(&sa, &start)
		return sa.Value, err
	case "WebpageDescriptorAttribute":
		wda := WebpageDescriptorAttribute{}
		err := dec.DecodeElement(&wda, &start)
		return wda.Value, err
	default:
		return nil, fmt.Errorf("unknown attribute type %#v", attributeType)
	}
}

func (tis TargetingIdeaSelector) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	e.EncodeToken(start)

	for _, searchParam := range tis.SearchParameters {
		err := searchParameterMarshalXML(searchParam, e)
		if err != nil {
			return err
		}
	}

	e.EncodeElement(&tis.IdeaType, xml.StartElement{Name: xml.Name{"", "ideaType"}})
	e.EncodeElement(&tis.RequestType, xml.StartElement{Name: xml.Name{"", "requestType"}})
	e.EncodeElement(&tis.RequestedAttributeTypes, xml.StartElement{Name: xml.Name{"", "requestedAttributeTypes"}})
	e.EncodeElement(&tis.Paging, xml.StartElement{Name: xml.Name{"", "paging"}})

	if tis.LocaleCode != "" {
		e.EncodeElement(&tis.LocaleCode, xml.StartElement{Name: xml.Name{"", "localCode"}})
	}

	if tis.CurrencyCode != "" {
		e.EncodeElement(&tis.CurrencyCode, xml.StartElement{Name: xml.Name{"", "currencyCode"}})
	}

	e.EncodeToken(start.End())
	return nil
}

func searchParameterMarshalXML(sp SearchParameter, e *xml.Encoder) error {
	searchType := ""

	switch t := sp.(type) {
	case CategoryProductsAndServicesSearchParameter:
		searchType = "CategoryProductsAndServicesSearchParameter"
	case CompetitionSearchParameter:
		searchType = "CompetitionSearchParameter"
	case IdeaTextFilterSearchParameter:
		searchType = "IdeaTextFilterSearchParameter"
	case IncludeAdultContentSearchParameter:
		searchType = "IncludeAdultContentSearchParameter"
	case LanguageSearchParameter:
		searchType = "LanguageSearchParameter"
	case LocationSearchParameter:
		searchType = "LocationSearchParameter"
	case NetworkSearchParameter:
		searchType = "NetworkSearchParameter"
	case RelatedToQuerySearchParameter:
		searchType = "RelatedToQuerySearchParameter"
	case RelatedToUrlSearchParameter:
		searchType = "RelatedToUrlSearchParameter"
	case SearchVolumeSearchParameter:
		searchType = "SearchVolumeSearchParameter"
	case SeedAdGroupIdSearchParameter:
		searchType = "SeedAdGroupIdSearchParameter"
	default:
		return fmt.Errorf("unknown search parameter type %#v\n", t)
	}

	// encode the inner element
	e.EncodeElement(&sp, xml.StartElement{
		xml.Name{"", "searchParameters"},
		[]xml.Attr{
			xml.Attr{xml.Name{"http://www.w3.org/2001/XMLSchema-instance", "type"}, searchType},
		},
	})

	return nil
}

// Get Returns a page of ideas that match the query described by the specified TargetingIdeaSelector.
// https://developers.google.com/adwords/api/docs/reference/v201708/TargetingIdeaService
func (s *TargetingIdeaService) Get(selector TargetingIdeaSelector) (targetingIdeas []TargetingIdeas, totalCount int64, err error) {

	respBody, err := s.Auth.request(
		targetingIdeaServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     TargetingIdeaSelector `xml:"selector"`
		}{
			XMLName: xml.Name{
				Space: baseTrafficUrl,
				Local: "get",
			},
			Sel: selector,
		},
		nil,
	)
	if err != nil {
		return targetingIdeas, totalCount, err
	}
	getResp := struct {
		Size           int64            `xml:"rval>totalNumEntries"`
		TargetingIdeas []TargetingIdeas `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return targetingIdeas, totalCount, err
	}
	return getResp.TargetingIdeas, getResp.Size, err
}
