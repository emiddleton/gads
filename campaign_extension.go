package gads

import (
	"encoding/xml"
)

type CampaignExtensionSettingService struct {
	Auth
}

func NewCampaignExtensionService(auth *Auth) *CampaignExtensionSettingService {
	return &CampaignExtensionSettingService{Auth: *auth}
}

// https://developers.google.com/adwords/api/docs/reference/v201708/CampaignExtensionSettingService.CampaignExtensionSetting
// A CampaignExtensionSetting is used to add or modify extensions being served for the specified campaign.
type CampaignExtensionSetting struct {
	CampaignId       int64            `xml:"https://adwords.google.com/api/adwords/cm/v201708 campaignId,omitempty"`
	ExtensionType    FeedType         `xml:"https://adwords.google.com/api/adwords/cm/v201708 extensionType,omitempty"`
	ExtensionSetting ExtensionSetting `xml:"https://adwords.google.com/api/adwords/cm/v201708 extensionSetting,omitempty"`
}

type CampaignExtensionSettingOperations map[string][]CampaignExtensionSetting

// https://developers.google.com/adwords/api/docs/reference/v201708/CampaignExtensionSettingService#query
func (s *CampaignExtensionSettingService) Query(query string) (settings []CampaignExtensionSetting, totalCount int64, err error) {
	respBody, err := s.Auth.request(
		campaignExtensionSettingUrl,
		"query",
		AWQLQuery{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "query",
			},
			Query: query,
		},
		nil,
	)
	if err != nil {
		return
	}

	getResp := struct {
		Size     int64                      `xml:"rval>totalNumEntries"`
		Settings []CampaignExtensionSetting `xml:"rval>entries"`
	}{}

	err = xml.Unmarshal(respBody, &getResp)
	if err != nil {
		return
	}
	return getResp.Settings, getResp.Size, err
}

// https://developers.google.com/adwords/api/docs/reference/v201708/CampaignExtensionSettingService#mutate
func (s *CampaignExtensionSettingService) Mutate(settingsOperations CampaignExtensionSettingOperations) (settings []CampaignExtensionSetting, err error) {
	type settingOperations struct {
		Action  string                   `xml:"operator"`
		Setting CampaignExtensionSetting `xml:"operand"`
	}
	operations := []settingOperations{}
	for action, settings := range settingsOperations {
		for _, setting := range settings {
			operations = append(operations,
				settingOperations{
					Action:  action,
					Setting: setting,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []settingOperations `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: operations,
	}

	respBody, err := s.Auth.request(campaignExtensionSettingUrl, "mutate", mutation, nil)
	if err != nil {
		return settings, err
	}
	mutateResp := struct {
		Settings []CampaignExtensionSetting `xml:"rval>value"`
	}{}
	err = xml.Unmarshal(respBody, &mutateResp)
	if err != nil {
		return settings, err
	}

	return mutateResp.Settings, err
}
