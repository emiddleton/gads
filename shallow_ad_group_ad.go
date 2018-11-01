package gads

import "encoding/xml"

type ShallowAdGroupAdService struct {
	Auth
}

func NewShallowAdGroupAdService(auth *Auth) *ShallowAdGroupAdService {
	return &ShallowAdGroupAdService{Auth: *auth}
}

type ShallowAdGroupAd struct {
	Id   int64  `xml:"id"`
	Name string `xml:"name"`
}

func (s *ShallowAdGroupAdService) Get(selector Selector) (adGroupAds []ShallowAdGroupAd, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		adGroupAdServiceUrl,
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
		nil,
	)
	if err != nil {
		return adGroupAds, totalCount, err
	}
	getResp := struct {
		Size       int64              `xml:"rval>totalNumEntries"`
		AdGroupAds []ShallowAdGroupAd `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroupAds, totalCount, err
	}
	return getResp.AdGroupAds, getResp.Size, err
}
