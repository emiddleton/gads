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
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201708 getAgeRangeCriterion"`
		}{},
		nil,
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
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201708 getCarrierCriterion"`
		}{},
		nil,
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
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201708 getGenderCriterion"`
		}{},
		nil,
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
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201708 getLanguageCriterion"`
		}{},
		nil,
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
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201708 getMobileDeviceCriterion"`
		}{},
		nil,
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
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201708 getOperatingSystemVersionCriterion"`
		}{},
		nil,
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

func (s *ConstantDataService) GetProductBiddingCategoryCriterion(selector Selector) (categoryData []ProductBiddingCategoryData, err error) {
	respBody, err := s.Auth.request(
		constantDataServiceUrl,
		"getProductBiddingCategoryData",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: "https://adwords.google.com/api/adwords/cm/v201708",
				Local: "getProductBiddingCategoryData",
			},
			Sel: selector,
		},
		nil,
	)
	if err != nil {
		return categoryData, err
	}
	getResp := struct {
		ProductBiddingCategoryDatas []ProductBiddingCategoryData `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return categoryData, err
	}
	return getResp.ProductBiddingCategoryDatas, err
}

func (s *ConstantDataService) GetUserInterestCriterion() (userInterests []UserInterestCriterion, err error) {
	respBody, err := s.Auth.request(
		constantDataServiceUrl,
		"getUserInterestCriterion",
		struct {
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201708 getUserInterestCriterion"`
		}{},
		nil,
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
			XMLName xml.Name `xml:"https://adwords.google.com/api/adwords/cm/v201708 getVerticalCriterion"`
		}{},
		nil,
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
