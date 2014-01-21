package gads

import (
	"testing"
	//  "encoding/xml"
)

func testAdGroupService(t *testing.T) (service *adGroupService) {
	return &adGroupService{Auth: testAuthSetup(t)}
}

func testAdGroup(t *testing.T) (AdGroup, func()) {
	campaign, cleanupCampaign := testCampaign(t)
	ags := testAdGroupService(t)
	adGroups, err := ags.Mutate(
		AdGroupOperations{
			"ADD": {
				AdGroup{
					Name:                         "test ad group " + rand_str(10),
					Status:                       "PAUSED",
					CampaignId:                   campaign.Id,
					ContentBidCriterionTypeGroup: "KEYWORD",
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	cleanupAdGroup := func() {
		adGroups[0].Status = "DELETED"
		_, err = ags.Mutate(AdGroupOperations{"SET": adGroups})
		if err != nil {
			t.Error(err)
		}
		cleanupCampaign()
	}
	return adGroups[0], cleanupAdGroup
}

func TestAdGroup(t *testing.T) {
	campaign, cleanupCampaign := testCampaign(t)
	defer cleanupCampaign()

	ags := testAdGroupService(t)
	adGroups, err := ags.Mutate(
		AdGroupOperations{
			"ADD": {
				AdGroup{
					Name:                         "test ad group " + rand_str(10),
					Status:                       "PAUSED",
					CampaignId:                   campaign.Id,
					ContentBidCriterionTypeGroup: "KEYWORD",
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		adGroups[0].Status = "DELETED"
		_, err = ags.Mutate(AdGroupOperations{"SET": adGroups})
		if err != nil {
			t.Error(err)
		}
	}()

}
