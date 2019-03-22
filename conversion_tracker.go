package gads

import (
	"encoding/xml"
)

type ConversionTrackerService struct {
	Auth
}

func NewConversionTrackerService(auth *Auth) *ConversionTrackerService {
	return &ConversionTrackerService{Auth: *auth}
}

type ConversionTracker struct {
	Id                       int64  `xml:"id"`
	OriginalConversionTypeId int64  `xml:"originalConversionTypeId"`
	Name                     string `xml:"name"`
	Status                   string `xml:"status"`
	Category                 string `xml:"category"`
	GoogleEventSnippet       string `xml:"googleEventSnippet"`
	GoogleGlobalSiteTag      string `xml:"googleGlobalSiteTag"`
}

func (c *ConversionTrackerService) Get(selector Selector) (conversions []ConversionTracker, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := c.Auth.request(
		conversionTrackerServiceUrl,
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
		return conversions, totalCount, err
	}
	getResp := struct {
		Size        int64               `xml:"rval>totalNumEntries"`
		Conversions []ConversionTracker `xml:"rval>entries"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return conversions, totalCount, err
	}
	return getResp.Conversions, getResp.Size, err
}
