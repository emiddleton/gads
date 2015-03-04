package gads

import (
	"testing"
	//  "encoding/xml"
)

func testCampaignCriterionService(t *testing.T) (service *CampaignCriterionService) {
	return &CampaignCriterionService{Auth: testAuthSetup(t)}
}

func TestCampaignCriterion(t *testing.T) {
	campaign, cleanupCampaign := testCampaign(t)
	defer cleanupCampaign()

	ccs := testCampaignCriterionService(t)
	campaignCriterions, err := ccs.Mutate(
		CampaignCriterionOperations{
			"ADD": {
				CampaignCriterion{CampaignId: campaign.Id, Criterion: AdScheduleCriterion{DayOfWeek: "MONDAY", StartHour: "10", StartMinute: "ZERO", EndHour: "13", EndMinute: "ZERO"}},
				CampaignCriterion{CampaignId: campaign.Id, Criterion: Location{Id: 2392}},
				NegativeCampaignCriterion{CampaignId: campaign.Id, Criterion: KeywordCriterion{Text: rand_str(10), MatchType: "EXACT"}},
				NegativeCampaignCriterion{CampaignId: campaign.Id, Criterion: KeywordCriterion{Text: rand_str(10), MatchType: "EXACT"}},
				NegativeCampaignCriterion{CampaignId: campaign.Id, Criterion: KeywordCriterion{Text: rand_str(10), MatchType: "EXACT"}},
				NegativeCampaignCriterion{CampaignId: campaign.Id, Criterion: KeywordCriterion{Text: rand_str(10), MatchType: "EXACT"}},
				NegativeCampaignCriterion{CampaignId: campaign.Id, Criterion: KeywordCriterion{Text: rand_str(10), MatchType: "EXACT"}},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	//t.Fatalf("%#v\n",campaignCriterions)

	defer func() {
		_, err = ccs.Mutate(CampaignCriterionOperations{"REMOVE": campaignCriterions})
		if err != nil {
			t.Error(err)
		}
	}()
}
