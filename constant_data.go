package gads

import (
	"encoding/xml"
)

type ConstantDataService struct {
	Auth
}

func NewConstantDataService(auth *Auth) *ConstantDataService {
	return &ConstantDataService{Auth: *auth}
}

func (s *ConstantDataService) GetAgeRangeCriterion() (ageRanges []AgeRangeCriterion, err error) {
	respBody, err := s.Auth.request(
		constantDataServiceUrl,
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

func (s *ConstantDataService) GetCarrierCriterion() (carriers []CarrierCriterion, err error) {
	respBody, err := s.Auth.request(
		constantDataServiceUrl,
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

func (s *ConstantDataService) GetGenderCriterion() (genders []GenderCriterion, err error) {
	respBody, err := s.Auth.request(
		constantDataServiceUrl,
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

func (s *ConstantDataService) GetLanguageCriterion() (languages []LanguageCriterion, err error) {
	respBody, err := s.Auth.request(
		constantDataServiceUrl,
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

func (s *ConstantDataService) GetMobileDeviceCriterion() (mobileDevices []MobileDeviceCriterion, err error) {
	respBody, err := s.Auth.request(
		constantDataServiceUrl,
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

func (s *ConstantDataService) GetOperatingSystemVersionCriterion() (operatingSystemVersions []OperatingSystemVersionCriterion, err error) {
	respBody, err := s.Auth.request(
		constantDataServiceUrl,
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

func (s *ConstantDataService) GetUserInterestCriterion() (userInterests []UserInterestCriterion, err error) {
	respBody, err := s.Auth.request(
		constantDataServiceUrl,
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

func (s *ConstantDataService) GetVerticalCriterion() (verticals []VerticalCriterion, err error) {
	respBody, err := s.Auth.request(
		constantDataServiceUrl,
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
