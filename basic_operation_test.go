package gads

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func ExampleCampaignService_Get() {
	// load credentials from
	authConf, _ := NewCredentials(context.TODO())
	cs := NewCampaignService(&authConf.Auth)

	// This example illustrates how to retrieve all the campaigns for an account.
	var pageSize int64 = 500
	var offset int64 = 0
	paging := Paging{
		Offset: offset,
		Limit:  pageSize,
	}
	totalCount := 0
	for {
		campaigns, totalCount, err := cs.Get(
			Selector{
				Fields: []string{
					"Id",
					"Name",
					"Status",
				},
				Ordering: []OrderBy{
					{"Name", "ASCENDING"},
				},
				Paging: &paging,
			},
		)
		if err != nil {
			fmt.Printf("Error occured finding campaigns")
		}
		for _, c := range campaigns {
			fmt.Printf("Campaign ID %d, name '%s' and status '%s'", c.Id, c.Name, c.Status)
		}
		// Increment values to request the next page.
		offset += pageSize
		paging.Offset = offset
		if totalCount < offset {
			break
		}
	}
	fmt.Printf("\tTotal number of campaigns found: %d.", totalCount)
}

func ExampleCampaignService_Mutate() {
	// load credentials from
	authConf, err := NewCredentials(context.TODO())
	cs := NewCampaignService(&authConf.Auth)

	var budgetId int64 = 1

	// This example illustrates how to create campaigns.
	campaigns, err := cs.Mutate(
		CampaignOperations{
			"ADD": {
				Campaign{
					Name:   fmt.Sprintf("Interplanetary Cruise #%d", time.Now().Unix()),
					Status: "ACTIVE",
					BiddingStrategyConfiguration: &BiddingStrategyConfiguration{
						StrategyType: "MANUAL_CPC",
					},
					// Budget (required) - note only the budget ID is required
					BudgetId:               budgetId,
					AdvertisingChannelType: "SEARCH",
					// Optional Fields:
					StartDate:                   time.Now().Format("20060102"),
					AdServingOptimizationStatus: "ROTATE",
					NetworkSetting: &NetworkSetting{
						TargetGoogleSearch:         true,
						TargetSearchNetwork:        true,
						TargetContentNetwork:       false,
						TargetPartnerSearchNetwork: false,
					},
					Settings: []CampaignSetting{
						NewGeoTargetTypeSetting(
							"DONT_CARE",
							"DONT_CARE",
						),
					},
					FrequencyCap: &FrequencyCap{
						Impressions: 5,
						TimeUnit:    "DAY",
						Level:       "ADGROUP",
					},
				},
				Campaign{
					Name:   fmt.Sprintf("Interplanetary Cruise banner #%d", time.Now().Unix()),
					Status: "PAUSED",
					BiddingStrategyConfiguration: &BiddingStrategyConfiguration{
						StrategyType: "MANUAL_CPC",
					},
					// Budget (required) - note only the budget ID is required
					BudgetId:               budgetId,
					AdvertisingChannelType: "DISPLAY",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("Error occured creating campaign.")
	}
	for _, c := range campaigns {
		fmt.Printf("Campaign with name '%s' and ID %d was added.", c.Name, c.Id)
	}

	// This example illustrates how to update a campaign, setting its status to 'PAUSED'
	campaigns, err = cs.Mutate(
		CampaignOperations{
			"SET": {
				Campaign{
					Id:     campaigns[0].Id,
					Status: "PAUSED",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("No campaigns were updated.")
	} else {
		fmt.Printf("Campaign ID %d was successfully updated, status was set to '%s'.", campaigns[0].Id, campaigns[0].Status)
	}

	// This example removes a campaign by setting the status to 'REMOVED'.
	campaigns, err = cs.Mutate(
		CampaignOperations{
			"SET": {
				Campaign{
					Id:     campaigns[0].Id,
					Status: "REMOVED",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("No campaigns were updated.")
	} else {
		fmt.Printf("Campaign ID %d was removed.", campaigns[0].Id)
	}
}

func ExampleAdGroupService_Get() {
	authConf, _ := NewCredentials(context.TODO())
	ags := NewAdGroupService(&authConf.Auth)

	// This example illustrates how to retrieve all the ad groups for a campaign.
	campaignId := "3"
	var pageSize int64 = 500
	var offset int64 = 0
	paging := Paging{
		Offset: offset,
		Limit:  pageSize,
	}
	totalCount := 0
	for {
		adGroups, totalCount, err := ags.Get(
			Selector{
				Fields: []string{
					"Id",
					"Name",
				},
				Ordering: []OrderBy{
					{"Name", "ASCENDING"},
				},
				Predicates: []Predicate{
					{"CampaignId", "IN", []string{campaignId}},
				},
				Paging: &paging,
			},
		)
		if err != nil {
			fmt.Printf("Error occured finding ad group")
		}
		for _, ag := range adGroups {
			fmt.Printf("Ad group name is '%s' and ID is %d", ag.Id, ag.Name)
		}
		// Increment values to request the next page.
		offset += pageSize
		paging.Offset = offset
		if totalCount < offset {
			break
		}
	}
	fmt.Printf("\tCampaign ID %d has %d ad group(s).", campaignId, totalCount)
}

func ExampleAdGroupService_Mutate() {
	authConf, err := NewCredentials(context.TODO())
	ags := NewAdGroupService(&authConf.Auth)

	var campaignId int64 = 1

	// This example illustrates how to create ad groups.
	adGroups, err := ags.Mutate(
		AdGroupOperations{
			"ADD": {
				AdGroup{
					Name:       fmt.Sprintf("Earth to Mars Cruises #%d", time.Now().Unix()),
					Status:     "ENABLED",
					CampaignId: campaignId,
					BiddingStrategyConfiguration: []BiddingStrategyConfiguration{
						{
							Bids: []Bid{
								Bid{
									Type:   "CpcBid",
									Amount: 10000000,
								},
							},
						},
					},
					Settings: []AdSetting{
						AdSetting{
							Details: []TargetSettingDetail{
								TargetSettingDetail{
									CriterionTypeGroup: "PLACEMENT",
									TargetAll:          true,
								},
								TargetSettingDetail{
									CriterionTypeGroup: "VERTICAL",
									TargetAll:          false,
								},
							},
						},
					},
				},
				AdGroup{
					Name:       fmt.Sprintf("Earth to Pluto Cruises #%d", time.Now().Unix()),
					Status:     "ENABLED",
					CampaignId: campaignId,
					BiddingStrategyConfiguration: []BiddingStrategyConfiguration{
						{
							Bids: []Bid{
								Bid{
									Type:   "CpcBid",
									Amount: 10000000,
								},
							},
						},
					},
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("")
	} else {
		for _, ag := range adGroups {
			fmt.Printf("Ad group ID %d was successfully added.", ag.Id)
		}
	}

	// This example illustrates how to update an ad group
	adGroups, err = ags.Mutate(
		AdGroupOperations{
			"SET": {
				AdGroup{
					Id:     adGroups[0].Id,
					Status: "PAUSE",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("No ad groups were updated.")
	} else {
		fmt.Printf("Ad group id %d was successfully updated.", adGroups[0].Id)
	}

	// This example removes an ad group by setting the status to 'REMOVED'.
	adGroups, err = ags.Mutate(
		AdGroupOperations{
			"SET": {
				AdGroup{
					Id:     adGroups[0].Id,
					Status: "REMOVE",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("No ad groups were updated.")
	} else {
		fmt.Printf("Ad group id %d was successfully removed.", adGroups[0].Id)
	}
}

func ExampleAdGroupCriterionService_Get() {
	authConf, _ := NewCredentials(context.TODO())
	agcs := NewAdGroupCriterionService(&authConf.Auth)

	// This example illustrates how to retrieve all keywords for an ad group.
	adGroupId := "1"
	var pageSize int64 = 500
	var offset int64 = 0
	paging := Paging{
		Offset: offset,
		Limit:  pageSize,
	}
	for {
		adGroupCriterions, totalCount, err := agcs.Get(
			Selector{
				Fields: []string{
					"Id",
					"CriteriaType",
					"KeywordText",
				},
				Ordering: []OrderBy{
					{"Id", "ASCENDING"},
				},
				Predicates: []Predicate{
					{"AdGroupId", "EQUALS", []string{adGroupId}},
					{"CriteriaType", "EQUALS", []string{"KEYWORD"}},
				},
				Paging: &paging,
			},
		)
		if err != nil {
			fmt.Printf("Error occured finding ad group criterion")
		}
		for _, agc := range adGroupCriterions {
			kc := agc.(BiddableAdGroupCriterion).Criterion.(KeywordCriterion)
			fmt.Printf("Keyword ID %d, type '%s' and text '%s'", kc.Id, kc.MatchType, kc.Text)
		}
		// Increment values to request the next page.
		offset += pageSize
		paging.Offset = offset
		if totalCount < offset {
			fmt.Printf("\tAd group ID %d has %d keyword(s).", totalCount)
			break
		}
	}
}

func ExampleAdGroupCriterionService_Mutate() {
	authConf, err := NewCredentials(context.TODO())
	agcs := NewAdGroupCriterionService(&authConf.Auth)

	var adGroupId int64 = 1

	// This example illustrates how to add multiple keywords to a given ad group.
	adGroupCriterions, err := agcs.Mutate(
		AdGroupCriterionOperations{
			"ADD": {
				BiddableAdGroupCriterion{
					AdGroupId: adGroupId,
					Criterion: KeywordCriterion{
						Text:      "mars cruise",
						MatchType: "BROAD",
					},
					UserStatus:     "PAUSED",
					DestinationUrl: "http://example.com/mars",
				},
				BiddableAdGroupCriterion{
					AdGroupId: adGroupId,
					Criterion: KeywordCriterion{
						Text:      "space hotel",
						MatchType: "BROAD",
					},
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("No keywords were added.")
	} else {
		fmt.Printf("Added %d keywords to ad group ID %d:", len(adGroupCriterions), adGroupId)
		for _, agc := range adGroupCriterions {
			k := agc.(BiddableAdGroupCriterion).Criterion.(KeywordCriterion)
			fmt.Printf("\tKeyword ID is %d and type is '%s'", k.Id, k.MatchType)
		}
	}

	// This example updates the bid of a keyword.
	keywordCriterion := adGroupCriterions[0].(BiddableAdGroupCriterion).Criterion.(KeywordCriterion)
	biddingStrategyConfigurations := BiddingStrategyConfiguration{
		Bids: []Bid{
			Bid{
				Type:   "CpcBid",
				Amount: 10000000,
			},
		},
	}
	adGroupCriterions, err = agcs.Mutate(
		AdGroupCriterionOperations{
			"SET": {
				BiddableAdGroupCriterion{
					AdGroupId:                    adGroupId,
					Criterion:                    keywordCriterion,
					BiddingStrategyConfiguration: &biddingStrategyConfigurations,
				},
			},
		},
	)
	biddableAdGroupCriterion := adGroupCriterions[0].(BiddableAdGroupCriterion)
	keywordCriterion = biddableAdGroupCriterion.Criterion.(KeywordCriterion)
	if err != nil {
		fmt.Printf("No keywords were updated.")
	} else {
		fmt.Printf("Keyword ID %d was successfully updated, current bids are:", keywordCriterion.Id)
		for _, bid := range biddableAdGroupCriterion.BiddingStrategyConfiguration.Bids {
			fmt.Printf("\tType: '%s', value: %d", bid.Type, bid.Amount)
		}
	}

	// This example removes a keyword using the 'REMOVE' operator.
	adGroupCriterions, err = agcs.Mutate(
		AdGroupCriterionOperations{
			"REMOVE": {
				BiddableAdGroupCriterion{
					AdGroupId: adGroupId,
					Criterion: keywordCriterion,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("No keywords were removed.")
	} else {
		biddableAdGroupCriterion := adGroupCriterions[0].(BiddableAdGroupCriterion)
		keywordCriterion = biddableAdGroupCriterion.Criterion.(KeywordCriterion)
		fmt.Printf("Keyword ID %d was successfully removed.", keywordCriterion.Id)
	}
}

func ExampleAdGroupAdService_Get() {
	authConf, _ := NewCredentials(context.TODO())
	agas := NewAdGroupAdService(&authConf.Auth)

	// This example illustrates how to retrieve all text ads for an ad group.
	adGroupId := "1"
	var pageSize int64 = 500
	var offset int64 = 0
	paging := Paging{
		Offset: offset,
		Limit:  pageSize,
	}
	var totalCount int64 = 0
	for {
		adGroupAds, totalCount, err := agas.Get(
			Selector{
				Fields: []string{
					"Id",
					"Status",
					"AdType",
				},
				Ordering: []OrderBy{
					{"Id", "ASCENDING"},
				},
				Predicates: []Predicate{
					{"AdGroupId", "IN", []string{adGroupId}},
					{"Status", "IN", []string{"ENABLED", "PAUSED", "DISABLED"}},
					{"AdType", "EQUALS", []string{"TEXT_AD"}},
				},
				Paging: &paging,
			},
		)
		if err != nil {
			fmt.Printf("Error occured finding ad group ad")
		}
		for _, aga := range adGroupAds {
			ta := aga.(TextAd)
			fmt.Printf("Ad ID is %d, type is 'TextAd' and status is '%s'", ta.Id, ta.Status)
		}
		// Increment values to request the next page.
		offset += pageSize
		paging.Offset = offset
		if totalCount < offset {
			break
		}
	}
	fmt.Printf("\tAd group ID %d has %d ad(s).", totalCount)
}

func ExampleAdGroupAdService_Mutate() {
	authConf, err := NewCredentials(context.TODO())
	agas := NewAdGroupAdService(&authConf.Auth)

	// This example illustrates how to add text ads to a given ad group.
	var adGroupId int64 = 1

	adGroupAds, err := agas.Mutate(
		AdGroupAdOperations{
			"ADD": {
				NewTextAd(
					adGroupId,
					"http://www.example.com",
					"example.com",
					"Luxury Cruise to Mars",
					"Visit the Red Planet in style.",
					"Low-gravity fun for everyone!",
					"ACTIVE",
				),
				NewTextAd(
					adGroupId,
					"http://www.example.com",
					"www.example.com",
					"Luxury Cruise to Mars",
					"Enjoy your stay at Red Planet.",
					"Buy your tickets now!",
					"ACTIVE",
				),
			},
		},
	)
	if err != nil {
		fmt.Printf("No ads were added.")
	} else {
		fmt.Printf("Added %d ad(s) to ad group ID %d:", len(adGroupAds), adGroupId)
		for _, ada := range adGroupAds {
			ta := ada.(TextAd)
			fmt.Printf("\tAd ID %d, type 'TextAd' and status '%s'", ta.Id, ta.Status)
		}
	}

	// This example illustrates how to update an ad, setting its status to 'PAUSED'.
	textAdId := adGroupAds[0].(TextAd).Id
	adGroupAds, err = agas.Mutate(
		AdGroupAdOperations{
			"SET": {
				TextAd{
					AdGroupId: adGroupId,
					Id:        textAdId,
					Status:    "PAUSED",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("No ads were updated.")
	} else {
		textAd := adGroupAds[0].(TextAd)
		fmt.Printf("Ad ID %d was successfully updated, status set to '%s'.", textAd.Id, textAd.Status)
	}

	// This example removes an ad using the 'REMOVE' operator.
	adGroupAds, err = agas.Mutate(
		AdGroupAdOperations{
			"SET": {
				TextAd{
					AdGroupId: adGroupId,
					Id:        textAdId,
					Status:    "REMOVE",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("No ads were removed.")
	} else {
		textAd := adGroupAds[0].(TextAd)
		fmt.Printf("Ad ID %d was successfully removed.", textAd.Id)
	}
}
