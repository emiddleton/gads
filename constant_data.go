package gads

import (
	"encoding/xml"
)

var (
	CONSTANT_DATA_SERVICE_URL = ServiceUrl{"https://adwords.google.com/api/adwords/cm/v201309", "ConstantDataService"}
)

type constantDataService struct {
	Auth
}

func NewConstantDataService(auth Auth) *constantDataService {
	return &constantDataService{Auth: auth}
}

func (s *constantDataService) GetAgeRangeCriterion() (ageRanges []AgeRangeCriterion, err error) {
	respBody, err := s.Auth.Request(
		CONSTANT_DATA_SERVICE_URL,
		"getAgeRangeCriterion",
		struct {
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201309 getAgeRangeCriterion"`
		}{},
	)
	if err != nil {
		return ageRanges, err
	}
	getResp := struct {
		AgeRangeCriterions []AgeRangeCriterion `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return ageRanges, err
	}
	return getResp.AgeRangeCriterions, err
}

func (s *constantDataService) GetCarrierCriterion() (carriers []CarrierCriterion, err error) {
	respBody, err := s.Auth.Request(
		CONSTANT_DATA_SERVICE_URL,
		"getCarrierCriterion",
		struct {
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201309 getCarrierCriterion"`
		}{},
	)
	if err != nil {
		return carriers, err
	}
	getResp := struct {
		CarrierCriterions []CarrierCriterion `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return carriers, err
	}
	return getResp.CarrierCriterions, err
}

func (s *constantDataService) GetGenderCriterion() (genders []GenderCriterion, err error) {
	respBody, err := s.Auth.Request(
		CONSTANT_DATA_SERVICE_URL,
		"getGenderCriterion",
		struct {
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201309 getGenderCriterion"`
		}{},
	)
	if err != nil {
		return genders, err
	}
	getResp := struct {
		GenderCriterions []GenderCriterion `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return genders, err
	}
	return getResp.GenderCriterions, err
}

func (s *constantDataService) GetLanguageCriterion() (languages []LanguageCriterion, err error) {
	respBody, err := s.Auth.Request(
		CONSTANT_DATA_SERVICE_URL,
		"getLanguageCriterion",
		struct {
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201309 getLanguageCriterion"`
		}{},
	)
	if err != nil {
		return languages, err
	}
	getResp := struct {
		LanguageCriterions []LanguageCriterion `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return languages, err
	}
	return getResp.LanguageCriterions, err
}

func (s *constantDataService) GetMobileDeviceCriterion() (mobileDevices []MobileDeviceCriterion, err error) {
	respBody, err := s.Auth.Request(
		CONSTANT_DATA_SERVICE_URL,
		"getMobileDeviceCriterion",
		struct {
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201309 getMobileDeviceCriterion"`
		}{},
	)
	if err != nil {
		return mobileDevices, err
	}
	getResp := struct {
		MobileDeviceCriterions []MobileDeviceCriterion `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return mobileDevices, err
	}
	return getResp.MobileDeviceCriterions, err
}

func (s *constantDataService) GetOperatingSystemVersionCriterion() (operatingSystemVersions []OperatingSystemVersionCriterion, err error) {
	respBody, err := s.Auth.Request(
		CONSTANT_DATA_SERVICE_URL,
		"getOperatingSystemVersionCriterion",
		struct {
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201309 getOperatingSystemVersionCriterion"`
		}{},
	)
	if err != nil {
		return operatingSystemVersions, err
	}
	getResp := struct {
		OperatingSystemVersionCriterions []OperatingSystemVersionCriterion `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return operatingSystemVersions, err
	}
	return getResp.OperatingSystemVersionCriterions, err
}

func (s *constantDataService) GetUserInterestCriterion() (userInterests []UserInterestCriterion, err error) {
	respBody, err := s.Auth.Request(
		CONSTANT_DATA_SERVICE_URL,
		"getUserInterestCriterion",
		struct {
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201309 getUserInterestCriterion"`
		}{},
	)
	if err != nil {
		return userInterests, err
	}
	getResp := struct {
		UserInterestCriterions []UserInterestCriterion `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return userInterests, err
	}
	return getResp.UserInterestCriterions, err
}

func (s *constantDataService) GetVerticalCriterion() (verticals []VerticalCriterion, err error) {
	respBody, err := s.Auth.Request(
		CONSTANT_DATA_SERVICE_URL,
		"getVerticalCriterion",
		struct {
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201309 getVerticalCriterion"`
		}{},
	)
	if err != nil {
		return verticals, err
	}
	getResp := struct {
		VerticalCriterions []VerticalCriterion `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return verticals, err
	}
	return getResp.VerticalCriterions, err
}
