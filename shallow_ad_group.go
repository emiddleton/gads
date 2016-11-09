package gads

import "encoding/xml"

type ShallowAdGroupService struct {
	Auth
}

func NewShallowAdGroupService(auth *Auth) *ShallowAdGroupService {
	return &ShallowAdGroupService{Auth: *auth}
}

type ShallowAdGroup struct {
	Id                             int64                           `xml:"id"`
	Name                           string                          `xml:"name"`
}

func (s *ShallowAdGroupService) Get(selector Selector) (adGroups []ShallowAdGroup, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		adGroupServiceUrl,
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
		return adGroups, totalCount, err
	}
	getResp := struct {
		Size     int64     `xml:"rval>totalNumEntries"`
		AdGroups []ShallowAdGroup `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroups, totalCount, err
	}
	return getResp.AdGroups, getResp.Size, err
}
