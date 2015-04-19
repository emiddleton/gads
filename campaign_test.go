package gads

import (
	"fmt"
	"testing"
	"time"
)

func testCampaignService(t *testing.T) (service *CampaignService) {
	return &CampaignService{Auth: testAuthSetup(t)}
}

func testCampaign(t *testing.T) (Campaign, func()) {
	budget, cleanupBudget := testBudget(t)
	cs := testCampaignService(t)
	campaigns, err := cs.Mutate(
		CampaignOperations{
			"ADD": {
				Campaign{
					Name:                        "test campaign " + rand_str(10),
					Status:                      "PAUSED",
					StartDate:                   time.Now().Format("20060102"),
					BudgetId:                    budget.Id,
					AdServingOptimizationStatus: "ROTATE_INDEFINITELY",
					Settings: []CampaignSetting{
						NewRealTimeBiddingSetting(true),
					},
					AdvertisingChannelType: "SEARCH",
					BiddingStrategyConfiguration: &BiddingStrategyConfiguration{
						StrategyType: "MANUAL_CPC",
					},
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	cleanupCampaign := func() {
		campaigns[0].Status = "REMOVED"
		_, err = cs.Mutate(CampaignOperations{"SET": campaigns})
		if err != nil {
			t.Error(err)
		}
		cleanupBudget()
	}
	return campaigns[0], cleanupCampaign
}

func TestCampaign(t *testing.T) {
	budget, cleanupBudget := testBudget(t)
	defer cleanupBudget()

	cs := testCampaignService(t)
	campaigns, err := cs.Mutate(
		CampaignOperations{
			"ADD": {
				Campaign{
					Name:                        "test campaign " + rand_str(10),
					Status:                      "PAUSED",
					StartDate:                   time.Now().Format("20060102"),
					BudgetId:                    budget.Id,
					AdServingOptimizationStatus: "ROTATE_INDEFINITELY",
					Settings: []CampaignSetting{
						NewRealTimeBiddingSetting(true),
					},
					AdvertisingChannelType: "SEARCH",
					NetworkSetting: &NetworkSetting{
						TargetGoogleSearch:         true,
						TargetSearchNetwork:        true,
						TargetContentNetwork:       false,
						TargetPartnerSearchNetwork: false,
					},
					BiddingStrategyConfiguration: &BiddingStrategyConfiguration{
						StrategyType: "MANUAL_CPC",
					},
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	defer func(campaigns []Campaign) {
		campaigns[0].Status = "REMOVED"
		_, err = cs.Mutate(CampaignOperations{"SET": campaigns})
		if err != nil {
			t.Error(err)
		}
	}(campaigns)

	label, labelCleanup := testLabel(t)
	defer labelCleanup()

	campaignLabels, err := cs.MutateLabel(
		CampaignLabelOperations{
			"ADD": {
				CampaignLabel{CampaignId: campaigns[0].Id, LabelId: label.Id},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		campaignLabels, err = cs.MutateLabel(CampaignLabelOperations{"REMOVE": campaignLabels})
		if err != nil {
			t.Fatal(err)
		}
	}()

	foundCampaigns, _, err := cs.Get(
		Selector{
			Fields: []string{
				"Id",
				"Name",
				"Status",
				"ServingStatus",
				"StartDate",
				"EndDate",
				"AdServingOptimizationStatus",
				"Settings",
				"Labels",
			},
			Predicates: []Predicate{
				{"Status", "EQUALS", []string{"PAUSED"}},
			},
			Ordering: []OrderBy{
				{"Id", "ASCENDING"},
			},
			Paging: &Paging{
				Offset: 0,
				Limit:  100,
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("found %d campaigns\n", len(foundCampaigns))
	for _, c := range campaigns {
		func(campaign Campaign) {
			for _, foundCampaign := range foundCampaigns {
				if foundCampaign.Id == campaign.Id {
					fmt.Printf("%#v", foundCampaign)
					return
				}
			}
			t.Errorf("campaign %d not found in \n%#v\n", campaign.Id, foundCampaigns)
		}(c)
	}

}
