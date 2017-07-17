package gads

import (
	"encoding/xml"
)

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService#query
type AdGroupExtensionSettingService struct {
	Auth
}

func NewAdGroupExtensionSettingService(auth *Auth) *AdGroupExtensionSettingService {
	return &AdGroupExtensionSettingService{Auth: *auth}
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService.AdGroupExtensionSetting
// An AdGroupExtensionSetting is used to add or modify extensions being served for the specified ad group.
type AdGroupExtensionSetting struct {
	AdGroupId        int64            `xml:"https://adwords.google.com/api/adwords/cm/v201609 adGroupId,omitempty"`
	ExtensionType    FeedType         `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensionType,omitempty"`
	ExtensionSetting ExtensionSetting `xml:"https://adwords.google.com/api/adwords/cm/v201609 extensionSetting,omitempty"`
}

type AdGroupExtensionSettingOperations map[string][]AdGroupExtensionSetting

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService#query
func (s *AdGroupExtensionSettingService) Query(query string) (settings []AdGroupExtensionSetting, totalCount int64, err error) {
	respBody, err := s.Auth.request(
		adGroupExtensionSettingServiceUrl,
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
		Size     int64                     `xml:"rval>totalNumEntries"`
		Settings []AdGroupExtensionSetting `xml:"rval>entries"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return
	}
	return getResp.Settings, getResp.Size, err
}

// https://developers.google.com/adwords/api/docs/reference/v201609/AdGroupExtensionSettingService#mutate
func (s *AdGroupExtensionSettingService) Mutate(settingsOperations AdGroupExtensionSettingOperations) (settings []AdGroupExtensionSetting, err error) {
	type settingOperations struct {
		Action  string                  `xml:"operator"`
		Setting AdGroupExtensionSetting `xml:"operand"`
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

	respBody, err := s.Auth.request(adGroupExtensionSettingServiceUrl, "mutate", mutation, nil)
	if err != nil {
		return settings, err
	}
	mutateResp := struct {
		Settings []AdGroupExtensionSetting `xml:"rval>value"`
	}{}
	err = xml.Unmarshal(respBody, &mutateResp)
	if err != nil {
		return settings, err
	}

	return mutateResp.Settings, err
}
