package gads

import (
	"encoding/xml"
)

var (
	LOCATION_CRITERION_SERVICE_URL = ServiceUrl{
		baseUrl,
		"LocationCriterionService",
	}
)

type locationCriterionService struct {
	Auth
}

func NewLocationCriterionService(auth Auth) *locationCriterionService {
	return &locationCriterionService{Auth: auth}
}

type LocationCriterion struct {
	Location      Location `xml:"location"`
	CanonicalName string   `xml:"canonicalName,omitempty"`
	Reach         string   `xml:"reach,omitempty"`
	Locale        string   `xml:"locale,omitempty"`
	SearchTerm    string   `xml:"searchTerm"`
}

type LocationCriterions []LocationCriterion

func (s *locationCriterionService) Get(selector Selector) (locationCriterions LocationCriterions, err error) {
	selector.XMLName = xml.Name{"", "selector"}
	respBody, err := s.Auth.Request(
		LOCATION_CRITERION_SERVICE_URL,
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
		return locationCriterions, err
	}
	getResp := struct {
		LocationCriterions LocationCriterions `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return locationCriterions, err
	}
	return getResp.LocationCriterions, err
}
