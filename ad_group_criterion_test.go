package gads

import (
	"testing"
	//	"encoding/xml"
)

func testAdGroupCriterionService(t *testing.T) (service *adGroupCriterionService) {
	return &adGroupCriterionService{Auth: testAuthSetup(t)}
}

func TestAdGroupCriterion(t *testing.T) {
	adGroup, _ := testAdGroup(t)
	//defer cleanupAdGroup()

	agcs := testAdGroupCriterionService(t)
	adGroupCriterions, err := agcs.Mutate(
		AdGroupCriterionOperations{
			"ADD": {
				// NewBiddableAdGroupCriterion(adGroup.Id, NewAgeRangeCriterion("AGE_RANGE_25_34")),
				// NewBiddableAdGroupCriterion(adGroup.Id, NewGenderCriterion()),
				// NewBiddableAdGroupCriterion(adGroup.Id, NewMobileAppCategoryCriterion(60000,"My Google Play Android Apps")),
				NewBiddableAdGroupCriterion(adGroup.Id, NewKeywordCriterion("test1", "EXACT")),
				NewBiddableAdGroupCriterion(adGroup.Id, NewKeywordCriterion("test2", "PHRASE")),
				NewBiddableAdGroupCriterion(adGroup.Id, NewKeywordCriterion("test3", "BROAD")),
				NewNegativeAdGroupCriterion(adGroup.Id, NewKeywordCriterion("test4", "BROAD")),
				NewBiddableAdGroupCriterion(adGroup.Id, NewPlacementCriterion("https://classdo.com")),
				// NewBiddableAdGroupCriterion(adGroup.Id, NewUserInterestCriterion()),
				// NewBiddableAdGroupCriterion(adGroup.Id, NewUserListCriterion()),
				// NewBiddableAdGroupCriterion(adGroup.Id, NewVerticalCriterion(0, 0, []string{"Pets & Anamals","Pets","Dogs"})),
				NewBiddableAdGroupCriterion(adGroup.Id, NewWebpageCriterion("test criterion", []WebpageCondition{WebpageCondition{Operand: "URL", Argument: "example.com"}})),
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_, err = agcs.Mutate(AdGroupCriterionOperations{"REMOVE": adGroupCriterions})
		if err != nil {
			t.Error(err)
		}
	}()
	/*
	   reqBody, err := xml.MarshalIndent(adGroupCriterions,"  ", "  ")
	   t.Fatalf("%s\n",reqBody)
	*/
}
