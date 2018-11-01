package gads

import "encoding/xml"

type ShallowCampaignService struct {
	Auth
}

func NewShallowCampaignService(auth *Auth) *ShallowCampaignService {
	return &ShallowCampaignService{Auth: *auth}
}

type ShallowCampaign struct {
	Id   int64  `xml:"id"`
	Name string `xml:"name"`
}

func (s *ShallowCampaignService) Get(selector Selector) (campaigns []ShallowCampaign, totalCount int64, err error) {
	selector.XMLName = xml.Name{"", "serviceSelector"}
	respBody, err := s.Auth.request(
		campaignServiceUrl,
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
		return campaigns, totalCount, err
	}
	getResp := struct {
		Size      int64             `xml:"rval>totalNumEntries"`
		Campaigns []ShallowCampaign `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return campaigns, totalCount, err
	}
	return getResp.Campaigns, getResp.Size, err
}
