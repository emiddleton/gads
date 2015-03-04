package gads

import (
	//"fmt"
	"testing"
)

func testConstantDataService(t *testing.T) (service *ConstantDataService) {
	return &ConstantDataService{Auth: testAuthSetup(t)}
}

func TestConstantData(t *testing.T) {
	/*
			cds := testConstantDataService(t)

			ageRangeCriterions, err := cds.GetAgeRangeCriterion()
		  if err != nil {
		    t.Error(err)
		  }
		  fmt.Printf("Age Range\n")
		  for _, ageRange := range ageRangeCriterions {
		    //fmt.Printf("%d. %s\n",ageRange.Id,ageRange.AgeRangeType)
		    fmt.Printf("%d. %#v\n",ageRange.Id,ageRange)
		  }

			carrierCriterions, err := cds.GetCarrierCriterion()
		  if err != nil {
		    t.Error(err)
		  }
		  fmt.Printf("\nCarrier\n")
		  for _, carrier := range carrierCriterions {
		    fmt.Printf("%d. %s,%s\n",carrier.Id, carrier.Name,carrier.CountryCode)
		  }

			genderCriterions, err := cds.GetGenderCriterion()
		  if err != nil {
		    t.Error(err)
		  }
		  fmt.Printf("\nGender\n")
		  for _, gender := range genderCriterions {
		    fmt.Printf("%d. %s\n",gender.Id, gender.GenderType)
		  }

			languageCriterions, err := cds.GetLanguageCriterion()
		  if err != nil {
		    t.Error(err)
		  }
		  fmt.Printf("\nLanguage\n")
		  for _, language := range languageCriterions {
		    fmt.Printf("%d. %s,%s\n", language.Id, language.Code, language.Name)
		  }

			mobileDeviceCriterions, err := cds.GetMobileDeviceCriterion()
		  if err != nil {
		    t.Error(err)
		  }
		  fmt.Printf("\nMobile Device\n")
		  for _, mobile := range mobileDeviceCriterions {
		    fmt.Printf("%d. %s,%s\n", mobile.Id, mobile.DeviceName, mobile.DeviceType)
		  }

			operatingSystemVersionCriterions, err := cds.GetOperatingSystemVersionCriterion()
		  if err != nil {
		    t.Error(err)
		  }
		  fmt.Printf("\nOperating System Version\n")
		  for _, osv := range operatingSystemVersionCriterions {
		    fmt.Printf("%d. %s,%s\n", osv.Id, osv.Name, osv.OperatorType)
		  }

			userInterestCriterions, err := cds.GetUserInterestCriterion()
		  if err != nil {
		    t.Error(err)
		  }
		  fmt.Printf("\nUser Interest\n")
		  for _, userInterest := range userInterestCriterions {
		    fmt.Printf("%d. %s\n", userInterest.Id, userInterest.Name)
		  }

			verticalCriterions, err := cds.GetVerticalCriterion()
		  if err != nil {
		    t.Error(err)
		  }
		  fmt.Printf("\nVertical\n")
		  for _, vertical := range verticalCriterions {
		    fmt.Printf("%d. %#v\n",vertical.Id, vertical.Path)
		  }
	*/
}
