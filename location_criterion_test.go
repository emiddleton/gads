package gads

import (
	"fmt"
	"testing"
	//  "encoding/xml"
)

func testLocationCriterionService(t *testing.T) (service *LocationCriterionService) {
	return &LocationCriterionService{Auth: testAuthSetup(t)}
}

func TestLocationCriterion(t *testing.T) {
	lcs := testLocationCriterionService(t)
	locationCriterions, err := lcs.Get(
		Selector{
			Fields: []string{
				"Id",
				"LocationName",
				"CanonicalName",
				"DisplayType",
				"ParentLocations",
				"Reach",
				"TargetingStatus",
			},
			Predicates: []Predicate{
				{"LocationName", "IN", []string{"Cameroon", "India", "Iraq", "Nigeria", "Pakistan", "Philippines"}},
				{"Locale", "EQUALS", []string{"en"}},
			},
		},
	)
	if err != nil {
		t.Error(err)
	}
	for _, locationCriterion := range locationCriterions {
		location := locationCriterion.Location
		fmt.Printf("%d. %s, %s\n", location.Id, location.LocationName, location.DisplayType)
	}
}
